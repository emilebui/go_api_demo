package service

import (
	"context"
	"go_api/models"
	"go_api/proto_gen"
	"testing"
)

func TestServer_TriggerScan(t *testing.T) {
	models.DB = models.InitTestDB()

	server := &Server{}
	_, err := server.TriggerScan(context.Background(), &proto_gen.ScanTriggerReq{Name: "test"})
	if err != nil {
		t.Fatalf("Fail in logic")
	}
}

func TestServer_GetScanResults(t *testing.T) {
	models.DB = models.InitTestDB()

	server := &Server{}
	_, err := server.GetScanResults(context.Background(), &proto_gen.GetScanResultReq{Name: "test"})
	if err != nil {
		t.Fatalf("fail in logic")
	}
}
