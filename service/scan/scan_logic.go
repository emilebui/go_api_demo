package scan

import (
	"bufio"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/google/uuid"
	"go_api/helpers"
	"go_api/load_config"
	"go_api/models"
	"go_api/proto_gen"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func InitScan(in *proto_gen.ScanTriggerReq) error {
	repoName := in.Name
	var repo models.Repo
	res := models.DB.First(&repo, "name = ?", repoName)
	if res.Error != nil {
		return res.Error
	}
	go InitScanLogic(repo)
	return nil
}

func InitScanLogic(repo models.Repo) {

	result := models.Result{
		Id:             uuid.NewString(),
		Status:         "Queued",
		RepositoryName: repo.Name,
		RepositoryUrl:  repo.Url,
		QueuedAt:       time.Now().Unix(),
	}

	models.DB.Create(&result)
	startScanning(&result)

}

func startScanning(result *models.Result) {

	result.Status = "In Progress"
	result.ScanningAt = time.Now().Unix()
	models.DB.Save(result)

	path, err := DownloadRepo(result)
	if err != nil {
		UpdateResultIfError(result, err)
		return
	}

	err = ScanRepo(path, result)
	if err != nil {
		UpdateResultIfError(result, err)
		return
	}

	// After scanning, updating Result info
	result.Status = "Success"
	result.FinishedAt = time.Now().Unix()
	models.DB.Save(result)
}

func UpdateResultIfError(result *models.Result, err error) {
	result.Status = "Failure"
	result.ErrorLog = err.Error()
	models.DB.Save(&result)
}

func DownloadRepo(result *models.Result) (string, error) {
	path := helpers.GetDirPath(fmt.Sprintf("temp/%s", result.Id))
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}
	_, err = git.PlainClone(path, false, &git.CloneOptions{
		URL: result.RepositoryUrl,
	})
	return path, err
}

func NotInIgnoreSet(path string) bool {
	for _, s := range load_config.Conf.SkipList {
		if strings.HasSuffix(path, s) {
			return false
		}
	}

	return true
}

func ScanRepo(path string, result *models.Result) error {
	var wg sync.WaitGroup
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && !NotInIgnoreSet(info.Name()) {
			return filepath.SkipDir
		}
		if !info.IsDir() && NotInIgnoreSet(path) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := ScanLogic(path, result)
				if err != nil {
					log.Fatal(err)
				}
			}()
		}
		return nil
	})
	wg.Wait()
	return err
}

func ScanLogic(path string, result *models.Result) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		lineText := scanner.Text()

		var rules []models.Rule
		res := models.DB.Find(&rules)
		if res.Error != nil {
			return res.Error
		}

		for _, val := range rules {
			if strings.Contains(lineText, val.StringCompare) {
				vul := models.Vulnerability{
					ID:       uuid.NewString(),
					ResultID: result.Id,
					Type:     "sast",
					RuleID:   val.ID,
					Path:     helpers.GetExtractedPath(result.Id, path),
					Line:     lineNum,
				}
				res = models.DB.Create(&vul)
				if res.Error != nil {
					return err
				}
				fmt.Printf("%s - Line %d\n", helpers.GetExtractedPath(result.Id, path), lineNum)
			}
		}
		lineNum++

	}

	return nil
}

func GetResults(in *proto_gen.GetScanResultReq) (*proto_gen.GetScanResultResp, error) {

	repoName := in.Name

	var repo models.Repo
	res := models.DB.First(&repo, "name = ?", repoName)
	if res.Error != nil {
		return nil, res.Error
	}

	response := proto_gen.GetScanResultResp{}
	var resultsList []*proto_gen.Result

	var results []models.Result
	models.DB.Where("repository_name = ?", repo.Name).Find(&results)

	for _, result := range results {
		var vuls []models.Vulnerability
		models.DB.Where("result_id = ?", result.Id).Find(&vuls)
		var findings []*proto_gen.Vulnerability
		for _, vul := range vuls {
			var rule models.Rule
			models.DB.First(&rule, "id = ?", vul.RuleID)
			findings = append(findings, &proto_gen.Vulnerability{
				Type:   "sast",
				RuleId: vul.RuleID,
				Location: &proto_gen.Vulnerability_Location{
					Line: int32(vul.Line),
					Path: vul.Path,
				},
				Metadata: &proto_gen.Vulnerability_MetaData{
					Description: rule.Description,
					Severity:    rule.Severity,
				},
			})
		}
		resultsList = append(resultsList, &proto_gen.Result{
			Id:         result.Id,
			Status:     result.Status,
			RepoName:   result.RepositoryName,
			RepoUrl:    result.RepositoryUrl,
			Findings:   findings,
			QueueAt:    result.QueuedAt,
			ScanningAt: result.ScanningAt,
			FinishedAt: result.FinishedAt,
		})
	}
	response.Data = resultsList

	return &response, nil
}
