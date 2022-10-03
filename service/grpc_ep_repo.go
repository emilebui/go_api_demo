package service

import (
	"context"
	"go_api/helpers"
	"go_api/proto_gen"
	"go_api/service/repo"
)

type Server struct {
	proto_gen.UnimplementedDemoServer
}

func (s *Server) HelloWorld(_ context.Context, _ *proto_gen.EmptyMessage) (*proto_gen.ResponseMessage, error) {
	msg := repo.HelloWorldLogic()
	return &proto_gen.ResponseMessage{Msg: msg}, nil
}

func (s *Server) CreateRepo(_ context.Context, in *proto_gen.CreateRepoReq) (*proto_gen.ResponseMessage, error) {

	err := in.Validate()
	if err != nil {
		return nil, helpers.Errorf(4090000, err.Error())
	}

	err = repo.CreateRepoLogic(in)
	if err != nil {
		return nil, helpers.Errorf(4090000, err.Error())
	}
	return &proto_gen.ResponseMessage{
		Msg: "Created repo successfully!",
	}, nil
}

func (s *Server) DeleteRepo(_ context.Context, in *proto_gen.DeleteRepoReq) (*proto_gen.ResponseMessage, error) {
	err := repo.DeleteRepoLogic(in)
	if err != nil {
		return nil, helpers.Errorf(4090000, err.Error())
	}
	return &proto_gen.ResponseMessage{
		Msg: "Deleted repo successfully!",
	}, nil
}

func (s *Server) UpdateRepo(_ context.Context, in *proto_gen.UpdateRepoReq) (*proto_gen.ResponseMessage, error) {
	err := repo.UpdateRepoLogic(in)
	if err != nil {
		return nil, helpers.Errorf(4090000, err.Error())
	}
	return &proto_gen.ResponseMessage{
		Msg: "Updated repo successfully!",
	}, nil
}

func (s *Server) GetRepo(_ context.Context, in *proto_gen.GetRepoReq) (*proto_gen.GetRepoResp, error) {
	result, err := repo.GetRepoLogic(in)
	if err != nil {
		return nil, helpers.Errorf(4090000, err.Error())
	}
	return &result, nil
}

func (s *Server) GetAllRepo(_ context.Context, _ *proto_gen.EmptyMessage) (*proto_gen.GetAllRepoRes, error) {
	res, err := repo.GetAllRepoLogic()
	if err != nil {
		return nil, helpers.Errorf(4090000, err.Error())
	}
	return res, nil
}
