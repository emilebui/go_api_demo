package service

import (
	"context"
	"go_api/helpers"
	"go_api/proto_gen"
	"go_api/service/rule"
)

func (s *Server) AddRule(_ context.Context, in *proto_gen.AddRule) (*proto_gen.ResponseMessage, error) {
	err := rule.AddRuleLogic(in)
	if err != nil {
		return nil, helpers.Errorf(4090000, err.Error())
	}

	return &proto_gen.ResponseMessage{
		Msg: "Created rule successfully!",
	}, nil
}

func (s *Server) EditRule(_ context.Context, in *proto_gen.Rule) (*proto_gen.ResponseMessage, error) {

	err := rule.EditRuleLogic(in)
	if err != nil {
		return nil, helpers.Errorf(4090000, err.Error())
	}

	return &proto_gen.ResponseMessage{
		Msg: "Edited rule successfully!",
	}, nil

}

func (s *Server) DeleteRule(_ context.Context, in *proto_gen.DeleteRuleReq) (*proto_gen.ResponseMessage, error) {

	err := rule.DeleteRuleLogic(in)

	if err != nil {
		return nil, helpers.Errorf(4090000, err.Error())
	}

	return &proto_gen.ResponseMessage{
		Msg: "Deleted rule successfully!",
	}, nil
}

func (s *Server) GetAllRules(_ context.Context, _ *proto_gen.EmptyMessage) (*proto_gen.GetAllRulesRes, error) {

	res, err := rule.GetAllRulesLogic()

	if err != nil {
		return nil, helpers.Errorf(4090000, err.Error())
	}

	return res, nil
}
