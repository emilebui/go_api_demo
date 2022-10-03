package repo

import (
	"go_api/models"
	"go_api/proto_gen"
	"gorm.io/gorm/utils/tests"
	"testing"
)

func TestHelloWorldLogic(t *testing.T) {
	msg := HelloWorldLogic()
	if msg != "Hello World" {
		t.Fatalf("Fail in logic")
	}
}

func TestCreateRepoLogic(t *testing.T) {

	models.DB = models.InitTestDB()

	err := CreateRepoLogic(&proto_gen.CreateRepoReq{
		Name: "abc",
		Url:  "blah",
	})

	if err != nil {
		t.Fatalf("Fail in logic")
	}

	err = CreateRepoLogic(&proto_gen.CreateRepoReq{
		Name: "test",
		Url:  "test",
	})

	if err == nil {
		t.Fatalf("Fail in logic")
	}
}

func TestDeleteRepoLogic(t *testing.T) {
	models.DB = models.InitTestDB()

	err := DeleteRepoLogic(&proto_gen.DeleteRepoReq{
		Name: "test",
	})

	if err != nil {
		t.Fatalf("Fail in logic")
	}

	_, err = GetRepoLogic(&proto_gen.GetRepoReq{
		Name: "test",
	})

	if err == nil {
		t.Fatalf("fail in logic")
	}
}

func TestGetRepoLogic(t *testing.T) {
	models.DB = models.InitTestDB()

	repo, err := GetRepoLogic(&proto_gen.GetRepoReq{
		Name: "test",
	})

	if err != nil {
		t.Fatalf("fail in logic")
	}
	tests.AssertEqual(t, repo.Name, "test")
	tests.AssertEqual(t, repo.Url, "https://github.com/emilebui/Minigames-on-Processing")

	repo, err = GetRepoLogic(&proto_gen.GetRepoReq{
		Name: "abc",
	})

	if err == nil {
		t.Fatalf("fail in logic")
	}
}

func TestUpdateRepoLogic(t *testing.T) {
	models.DB = models.InitTestDB()

	err := UpdateRepoLogic(&proto_gen.UpdateRepoReq{
		Name: "test",
		Url:  "blah",
	})

	if err != nil {
		t.Fatalf("fail in logic")
	}
}

func TestGetAllRepoLogic(t *testing.T) {
	models.DB = models.InitTestDB()

	res, err := GetAllRepoLogic()
	if err != nil {
		t.Fatalf("fail in logic")
	}
	tests.AssertEqual(t, len(res.Repositories), 1)
}
