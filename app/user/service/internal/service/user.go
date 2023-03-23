package service

import (
	"context"

	v1 "github.com/ynu-e-young/apis-go/htpp/user/service/v1"

	"github.com/ynu-e-young/htp-platform/app/user/service/internal/biz"
)

func (s *UserService) FindByUsername(ctx context.Context, in *v1.FindByUsernameRequest) (*v1.UserReply, error) {
	u, err := s.uu.FindByUsername(ctx, in.GetUsername())
	if err != nil {
		return nil, err
	}

	return &v1.UserReply{User: &v1.UserStruct{
		Id:           u.Id,
		Email:        u.Email,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}}, nil
}

func (s *UserService) FindByEmail(ctx context.Context, in *v1.FindByEmailRequest) (*v1.UserReply, error) {
	u, err := s.uu.FindByEmail(ctx, in.GetEmail())
	if err != nil {
		return nil, err
	}

	return &v1.UserReply{User: &v1.UserStruct{
		Id:           u.Id,
		Email:        u.Email,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}}, nil
}

func (s *UserService) Create(ctx context.Context, in *v1.CreateRequest) (*v1.UserReply, error) {
	inUser := in.GetUser()
	u, err := s.uu.Create(ctx, &biz.User{
		Id:           inUser.GetId(),
		Email:        inUser.GetEmail(),
		Username:     inUser.GetUsername(),
		PasswordHash: inUser.GetPasswordHash(),
	})
	if err != nil {
		return nil, err
	}

	return &v1.UserReply{User: &v1.UserStruct{
		Id:           u.Id,
		Email:        u.Email,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}}, nil
}

func (s *UserService) Update(ctx context.Context, in *v1.UpdateRequest) (*v1.UserReply, error) {
	inUser := in.GetUser()
	u, err := s.uu.Update(ctx, &biz.User{
		Id:           inUser.GetId(),
		Email:        inUser.GetEmail(),
		Username:     inUser.GetUsername(),
		PasswordHash: inUser.GetPasswordHash(),
	})
	if err != nil {
		return nil, err
	}

	return &v1.UserReply{User: &v1.UserStruct{
		Id:           u.Id,
		Email:        u.Email,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}}, nil
}

func (s *UserService) Get(ctx context.Context, in *v1.GetRequest) (*v1.UserReply, error) {
	u, err := s.uu.Get(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	return &v1.UserReply{User: &v1.UserStruct{
		Id:           u.Id,
		Email:        u.Email,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}}, nil
}
