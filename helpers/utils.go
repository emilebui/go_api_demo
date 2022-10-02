package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func FindLoadConfigPath(path string) string {

	maxTry := 10
	curTry := 0
	curPath := path

	for curTry < maxTry {
		if checkIfLoadConfigExist(curPath) {
			return curPath
		}
		curPath = filepath.Clean(filepath.Join(curPath, ".."))
		curTry++
	}

	return curPath
}

func checkIfLoadConfigExist(path string) bool {
	temp := filepath.Join(path, "load_config")
	return ExistsPath(temp)
}

func ExistsPath(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GetDirPath(subPath string) string {
	path, _ := os.Getwd()
	truePath := filepath.Join(FindLoadConfigPath(path), subPath)
	return truePath
}

func GetExtractedPath(key string, path string) string {
	paths := strings.Split(path, fmt.Sprintf("\\%s\\", key))
	return paths[1]
}
