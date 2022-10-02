package service

import (
	"context"
	"go_api/helpers"
	"go_api/proto_gen"
	"go_api/service/scan"
)

func (s *Server) TriggerScan(_ context.Context, in *proto_gen.ScanTriggerReq) (*proto_gen.ResponseMessage, error) {
	err := scan.InitScan(in)
	if err != nil {
		return nil, helpers.Errorf(4090000, err.Error())
	}
	return &proto_gen.ResponseMessage{
		Msg: "Scan triggered successfully!!",
	}, nil
}

func (s *Server) GetScanResults(_ context.Context, in *proto_gen.GetScanResultReq) (*proto_gen.GetScanResultResp, error) {
	res, err := scan.GetResults(in)

	if err != nil {
		return nil, helpers.Errorf(4090000, err.Error())
	}
	return res, err
}
