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

// FindLoadConfigPath Get the workdir folder
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

// checkIfLoadConfigExist Check if config folder exists
func checkIfLoadConfigExist(path string) bool {
	temp := filepath.Join(path, "load_config")
	return ExistsPath(temp)
}

// ExistsPath Check if path exists
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

// GetDirPath Get the absolute path from a relative path
func GetDirPath(subPath string) string {
	path, _ := os.Getwd()
	truePath := filepath.Join(FindLoadConfigPath(path), subPath)
	return truePath
}

// GetExtractedPath Get the relative path from key folder
func GetExtractedPath(key string, path string) string {

	// Get path on Window
	paths := strings.Split(path, fmt.Sprintf("\\%s\\", key))

	// Get path on linux
	if len(paths) < 2 {
		paths = strings.Split(path, fmt.Sprintf("/%s/", key))
	}

	return paths[1]
}
