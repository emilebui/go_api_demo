package demo

import (
	"context"
	"fmt"
	"go_api/helpers"
	"go_api/proto_gen"
)

type Server struct {
	proto_gen.UnimplementedDemoServer
}

func (s *Server) HelloWorld(ctx context.Context, in *proto_gen.EmptyMessage) (*proto_gen.ResponseMessage, error) {
	msg := HelloWorldLogic()
	return &proto_gen.ResponseMessage{Msg: msg}, nil
}

func (s *Server) CreateRepo(ctx context.Context, in *proto_gen.CreateRepoReq) (*proto_gen.ResponseMessage, error) {

	err := in.Validate()
	if err != nil {
		return nil, helpers.Errorf(4090000, "Repo name can't contain special characters")
	}

	fmt.Printf("%s - %s\n", in.Name, in.Url)

	err = CreateRepoLogic(in)
	if err != nil {
		return nil, helpers.Errorf(4090000, err.Error())
	}
	return &proto_gen.ResponseMessage{
		Msg: "Created repo successfully!",
	}, nil
}

func (s *Server) DeleteRepo(ctx context.Context, in *proto_gen.DeleteRepoReq) (*proto_gen.ResponseMessage, error) {
	err := DeleteRepoLogic(in)
	if err != nil {
		return nil, helpers.Errorf(4090000, err.Error())
	}
	return &proto_gen.ResponseMessage{
		Msg: "Deleted repo successfully!",
	}, nil
}

func (s *Server) UpdateRepo(ctx context.Context, in *proto_gen.UpdateRepoReq) (*proto_gen.ResponseMessage, error) {
	err := UpdateRepoLogic(in)
	if err != nil {
		return nil, helpers.Errorf(4090000, err.Error())
	}
	return &proto_gen.ResponseMessage{
		Msg: "Updated repo successfully!",
	}, nil
}

func (s *Server) GetRepo(ctx context.Context, in *proto_gen.GetRepoReq) (*proto_gen.GetRepoResp, error) {
	result, err := GetRepoLogic(in)
	if err != nil {
		return nil, helpers.Errorf(4090000, err.Error())
	}
	return &result, nil
}
