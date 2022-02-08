package service

import (
	"context"
	v1 "htp-platform/api/htpp/interface/v1"
	"htp-platform/app/htpp/interface/internal/biz"
)

func (s *InterfaceService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.UserReply, error) {
	u := in.GetUser()
	rv, err := s.uu.Login(ctx, u.GetEmail(), u.GetEmail())
	if err != nil {
		return nil, err
	}

	return &v1.UserReply{User: &v1.UserReply_User{
		Id:       rv.Id,
		Email:    rv.Email,
		Token:    rv.Token,
		Username: rv.Username,
	}}, nil
}

func (s *InterfaceService) Register(ctx context.Context, in *v1.RegisterRequest) (*v1.UserReply, error) {
	u := in.GetUser()
	rv, err := s.uu.Register(ctx, u.GetEmail(), u.GetEmail(), u.GetPassword())
	if err != nil {
		return nil, err
	}

	return &v1.UserReply{
		User: &v1.UserReply_User{
			Id:       rv.Id,
			Email:    rv.Email,
			Token:    rv.Token,
			Username: rv.Username,
		}}, nil
}

func (s *InterfaceService) GetCurrentUser(ctx context.Context, in *v1.GetCurrentUserRequest) (*v1.UserReply, error) {
	rv, err := s.uu.Get(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	return &v1.UserReply{User: &v1.UserReply_User{
		Id:       rv.Id,
		Email:    rv.Email,
		Token:    "",
		Username: rv.Username,
	}}, nil
}

func (s *InterfaceService) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.UserReply, error) {
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

	return &v1.UserReply{User: &v1.UserReply_User{
		Id:       rv.Id,
		Email:    rv.Email,
		Token:    rv.Token,
		Username: rv.Username,
	}}, nil
}
