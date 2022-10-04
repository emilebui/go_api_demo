package logger

import (
	"errors"
	"go_api/helpers"
	"go_api/load_config"
	"gorm.io/gorm/utils/tests"
	"testing"
)

func TestInitLogger(t *testing.T) {
	InitLogger()

	load_config.Conf.Log2File = true
	InitLogger()
	logPath := helpers.GetDirPath("log.txt")
	tests.AssertEqual(t, helpers.ExistsPath(logPath), true)
	load_config.Conf.Log2File = false
}

func TestLogInfo(t *testing.T) {
	LogInfo("blah", "blah")
}

func TestLogError(t *testing.T) {
	LogError(errors.New("blah"), "blah")
}
