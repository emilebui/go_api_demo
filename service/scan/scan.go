package scan

import (
	"bufio"
	"fmt"
	"github.com/go-git/go-git/v5"
	"go_api/helpers"
	"go_api/load_config"
	"go_api/models"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func DownloadRepo(url string, key string) (string, error) {
	path := helpers.GetDirPath(fmt.Sprintf("temp/%s", key))
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}
	_, err = git.PlainClone(path, false, &git.CloneOptions{
		URL: url,
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

func ScanRepo(path string, key string) error {
	var wg sync.WaitGroup
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && !NotInIgnoreSet(info.Name()) {
			return filepath.SkipDir
		}
		if !info.IsDir() && NotInIgnoreSet(path) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := ScanLogic(path, key)
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

func ScanLogic(path string, key string) error {

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
				//vul := models.Vulnerability{
				//	ID:       uuid.NewString(),
				//	ResultID: key,
				//	Type:     "sast",
				//	Rule:     val,
				//	Path:     helpers.GetExtractedPath(key, path),
				//}
				fmt.Printf("%s - Line %d\n", helpers.GetExtractedPath(key, path), lineNum)
			}
		}
		lineNum++

	}

	return nil
}
