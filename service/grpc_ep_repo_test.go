package service

import (
	"context"
	"go_api/models"
	"go_api/proto_gen"
	"testing"
)

func TestServer_HelloWorld(t *testing.T) {
	server := &Server{}
	r, err := server.HelloWorld(context.Background(), &proto_gen.EmptyMessage{})
	if err != nil {
		t.Fatalf("something is wrong, %v", err)
	}
	if r.Msg != "Hello World" {
		t.Fatalf("Output is different")
	}
}

func TestServer_CreateRepo(t *testing.T) {
	models.DB = models.InitTestDB()

	server := &Server{}
	_, err := server.CreateRepo(context.Background(), &proto_gen.CreateRepoReq{Name: "test2"})
	if err != nil {
		t.Fatalf("something is wrong, %v", err)
	}
	_, err = server.CreateRepo(context.Background(), &proto_gen.CreateRepoReq{Name: "test2"})
	if err == nil {
		t.Fatalf("something is wrong, %v", err)
	}
}

func TestServer_DeleteRepo(t *testing.T) {
	models.DB = models.InitTestDB()

	server := &Server{}
	_, err := server.DeleteRepo(context.Background(), &proto_gen.DeleteRepoReq{Name: "test"})
	if err != nil {
		t.Fatalf("something is wrong, %v", err)
	}
}

func TestServer_UpdateRepo(t *testing.T) {
	models.DB = models.InitTestDB()

	server := &Server{}
	_, err := server.UpdateRepo(context.Background(), &proto_gen.UpdateRepoReq{Name: "test", Url: "blah"})
	if err != nil {
		t.Fatalf("something is wrong, %v", err)
	}
	_, err = server.UpdateRepo(context.Background(), &proto_gen.UpdateRepoReq{Name: "test2", Url: "blah"})
	if err == nil {
		t.Fatalf("something is wrong, %v", err)
	}
}

func TestServer_GetRepo(t *testing.T) {
	models.DB = models.InitTestDB()

	server := &Server{}
	_, err := server.GetRepo(context.Background(), &proto_gen.GetRepoReq{Name: "test"})
	if err != nil {
		t.Fatalf("something is wrong, %v", err)
	}
	_, err = server.GetRepo(context.Background(), &proto_gen.GetRepoReq{Name: "test2"})
	if err == nil {
		t.Fatalf("something is wrong, %v", err)
	}
}

func TestServer_GetAllRepo(t *testing.T) {

	models.DB = models.InitTestDB()
	server := &Server{}
	_, err := server.GetAllRepo(context.Background(), &proto_gen.EmptyMessage{})
	if err != nil {
		t.Fatalf("something is wrong, %v", err)
	}
}
