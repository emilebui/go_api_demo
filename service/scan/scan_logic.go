package scan

import (
	"bufio"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/google/uuid"
	"go_api/helpers"
	"go_api/load_config"
	"go_api/logger"
	"go_api/models"
	"go_api/proto_gen"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

/*
-------------------------------------SCAN-LOGIC-DOCUMENT-----------------------------------
The entire scanning process works as follows:
	- Step 1: Create a Result obj, init an `ResultID` using uuid
	- Step 2: Download the repo into a newly-created folder named after `ResultID`
	- Step 3: Loop through every file in the download folder and do as follows
		- Skip file if it is in the Whitelist
		- Open a go-routine to scan the file
	- Scan file logic:
		- Read file line by line
		- Check if line contains keyword from Rule table
___________________________________________________________________________________________
*/

// InitScan Init Scan functions
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
	logger.LogInfo(fmt.Sprintf("Init scanning - Id %s", result.Id), "scan")
	startScanning(&result)

}

// startScanning Begin to Scan
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
	os.RemoveAll(path)
	logger.LogInfo(fmt.Sprintf("Finished Scanning - Id %s", result.Id), "scan")
}

// UpdateResultIfError Error Handler
func UpdateResultIfError(result *models.Result, err error) {
	result.Status = "Failure"
	result.ErrorLog = err.Error()
	models.DB.Save(&result)
}

// DownloadRepo Download Repo to Scan
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

// NotInIgnoreSet Checks if file/dir is in the ignore list
func NotInIgnoreSet(path string) bool {
	for _, s := range load_config.Conf.SkipList {
		if strings.HasSuffix(path, s) {
			return false
		}
	}

	return true
}

// ScanRepo Scan a repo
func ScanRepo(path string, result *models.Result) error {
	var wg sync.WaitGroup

	// Walk through every file in the repo path to scan
	// If file is valid (not in ignore-list), create a go-routine to scan (optimization)
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
					logger.LogError(err, "scan")
				}
			}()
		}
		return nil
	})
	wg.Wait()
	return err
}

// ScanLogic Scan a file by scanning every lines to see if lines contain dangerous words
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

				filePath := helpers.GetExtractedPath(result.Id, path)
				vul := models.Vulnerability{
					ID:       uuid.NewString(),
					ResultID: result.Id,
					Type:     "sast",
					RuleID:   val.ID,
					Path:     filePath,
					Line:     lineNum,
				}
				res = models.DB.Create(&vul)
				if res.Error != nil {
					return err
				}
				logger.LogInfo(fmt.Sprintf(
					"Found Vulnerability in Repo %s - Scan ID %s - File %s - Line %d - Rule %s (%s)",
					result.RepositoryName, result.Id, filePath, lineNum, val.ID, val.StringCompare), "scan")
			}
		}
		lineNum++

	}

	return nil
}

// GetResults Return scan results for a repo
func GetResults(in *proto_gen.GetScanResultReq) (*proto_gen.GetScanResultResp, error) {

	repoName := in.Name

	var repo models.Repo
	res := models.DB.First(&repo, "name = ?", repoName)
	if res.Error != nil {
		logger.LogError(res.Error, "scan")
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
	logger.LogInfo(fmt.Sprintf("Get all results from repo - %s", in.Name), "scan")
	return &response, nil
}
