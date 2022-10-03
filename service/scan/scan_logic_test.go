package scan

import (
	"go_api/models"
	"go_api/proto_gen"
	"gorm.io/gorm/utils/tests"
	"os"
	"strings"
	"testing"
)

func TestDownloadRepo(t *testing.T) {

	models.DB = models.InitTestDB()

	var result models.Result
	models.DB.First(&result, "id = ?", "test")

	path, err := DownloadRepo(&result)

	if err != nil {
		t.Fatalf("Fail in logic %v\n", err)
	}
	tests.AssertEqual(t, strings.HasSuffix(path, "test"), true)
	os.RemoveAll(path)
}

func TestGetResults(t *testing.T) {
	models.DB = models.InitTestDB()

	res, err := GetResults(&proto_gen.GetScanResultReq{
		Name: "test",
	})

	if err != nil {
		t.Fatalf("Fail in logic")
	}
	tests.AssertEqual(t, len(res.Data), 1)
}

func TestInitScanLogic(t *testing.T) {

	models.DB = models.InitTestDB()

	var repo models.Repo
	models.DB.First(&repo, "name = ?", "test")

	InitScanLogic(repo)
}

func TestInitScan(t *testing.T) {

	models.DB = models.InitTestDB()

	err := InitScan(&proto_gen.ScanTriggerReq{
		Name: "test",
	})
	if err != nil {
		t.Fatalf("Fail in logic")
	}
}
