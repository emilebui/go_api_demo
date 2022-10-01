package helpers

import (
	"os"
	"path/filepath"
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
	return exists(temp)
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
