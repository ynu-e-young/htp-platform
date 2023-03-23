package service

import (
	"context"

	interfaceV1 "github.com/ynu-e-young/apis-go/htpp/htpp/interface/v1"

	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/biz"
)

func (s *InterfaceService) Login(ctx context.Context, in *interfaceV1.LoginRequest) (*interfaceV1.UserReply, error) {
	u := in.GetUser()
	rv, err := s.uu.Login(ctx, u.GetEmail(), u.GetPassword())
	if err != nil {
		return nil, err
	}

	return &interfaceV1.UserReply{User: &interfaceV1.UserReply_User{
		Id:       rv.Id,
		Email:    rv.Email,
		Token:    rv.Token,
		Username: rv.Username,
	}}, nil
}

func (s *InterfaceService) Register(ctx context.Context, in *interfaceV1.RegisterRequest) (*interfaceV1.UserReply, error) {
	u := in.GetUser()
	rv, err := s.uu.Register(ctx, u.GetUsername(), u.GetEmail(), u.GetPassword())
	if err != nil {
		return nil, err
	}

	return &interfaceV1.UserReply{
		User: &interfaceV1.UserReply_User{
			Id:       rv.Id,
			Email:    rv.Email,
			Token:    rv.Token,
			Username: rv.Username,
		}}, nil
}

func (s *InterfaceService) GetCurrentUser(ctx context.Context, in *interfaceV1.GetCurrentUserRequest) (*interfaceV1.UserReply, error) {
	rv, err := s.uu.Get(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	return &interfaceV1.UserReply{User: &interfaceV1.UserReply_User{
		Id:       rv.Id,
		Email:    rv.Email,
		Token:    "",
		Username: rv.Username,
	}}, nil
}

func (s *InterfaceService) UpdateUser(ctx context.Context, in *interfaceV1.UpdateUserRequest) (*interfaceV1.UserReply, error) {
	u := in.GetUser()

	rv, err := s.uu.Update(ctx, &biz.User{
		Id:           u.GetId(),
		Email:        u.GetEmail(),
		Username:     u.GetUsername(),
		PasswordHash: u.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &interfaceV1.UserReply{User: &interfaceV1.UserReply_User{
		Id:       rv.Id,
		Email:    rv.Email,
		Token:    rv.Token,
		Username: rv.Username,
	}}, nil
}
