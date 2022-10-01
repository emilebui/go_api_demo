package load_config

import (
	"path/filepath"
	"testing"
)

func TestLoadConfigLogic(t *testing.T) {
	testConfigPath, err := filepath.Abs("test_config.yaml")
	if err != nil {
		t.Fatalf("unable to load test config")
	}

	conf := LoadConfigLogic(testConfigPath)

	if conf.GrpcPort != "FU" {
		t.Fatalf("value is different from config")
	}
}

func TestLoadConfigInit(t *testing.T) {
	conf := LoadConfigInit()
	if conf.GrpcPort != ":50051" {
		t.Fatalf("value is different from config")
	}
}
