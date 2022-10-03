package service

import (
	"context"
	"go_api/models"
	"go_api/proto_gen"
	"testing"
)

func TestServer_AddRule(t *testing.T) {

	models.DB = models.InitTestDB()
	server := &Server{}
	_, err := server.AddRule(context.Background(), &proto_gen.AddRule{
		Severity:      "blah",
		Description:   "blah",
		StringCompare: "blah",
	})
	if err != nil {
		t.Fatalf("something is wrong, %v", err)
	}
}

func TestServer_EditRule(t *testing.T) {

	models.DB = models.InitTestDB()
	server := &Server{}
	_, err := server.EditRule(context.Background(), &proto_gen.Rule{
		Description:   "blah",
		Id:            "test",
		Severity:      "blah",
		StringCompare: "blah",
	})
	if err != nil {
		t.Fatalf("something is wrong, %v", err)
	}

	_, err = server.EditRule(context.Background(), &proto_gen.Rule{
		Description:   "blah",
		Id:            "test2",
		Severity:      "blah",
		StringCompare: "blah",
	})
	if err == nil {
		t.Fatalf("something is wrong, %v", err)
	}
}

func TestServer_DeleteRule(t *testing.T) {

	models.DB = models.InitTestDB()
	server := &Server{}
	_, err := server.DeleteRule(context.Background(), &proto_gen.DeleteRuleReq{
		Id: "test",
	})
	if err != nil {
		t.Fatalf("something is wrong, %v", err)
	}
}

func TestServer_GetAllRules(t *testing.T) {

	models.DB = models.InitTestDB()
	server := &Server{}
	_, err := server.GetAllRules(context.Background(), &proto_gen.EmptyMessage{})
	if err != nil {
		t.Fatalf("something is wrong, %v", err)
	}
}
