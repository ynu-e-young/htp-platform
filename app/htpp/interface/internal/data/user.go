package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	interfaceV1 "github.com/ynu-e-young/apis-go/htpp/htpp/interface/v1"
	userV1 "github.com/ynu-e-young/apis-go/htpp/user/service/v1"

	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}

func (r *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	target, err := r.data.uc.FindByUsername(ctx, &userV1.FindByUsernameRequest{Username: username})
	if err != nil {
		return nil, interfaceV1.ErrorNotFoundError("find username: %s not found", username)
	}
	u := target.GetUser()

	return &biz.User{
		Id:           u.GetId(),
		Email:        u.GetEmail(),
		Username:     u.GetUsername(),
		PasswordHash: u.GetPasswordHash(),
	}, nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*biz.User, error) {
	target, err := r.data.uc.FindByEmail(ctx, &userV1.FindByEmailRequest{Email: email})
	if err != nil {
		return nil, interfaceV1.ErrorNotFoundError("find email: %s not found", email)
	}
	u := target.GetUser()

	return &biz.User{
		Id:           u.GetId(),
		Email:        u.GetEmail(),
		Username:     u.GetUsername(),
		PasswordHash: u.GetPasswordHash(),
	}, nil
}

func (r *userRepo) Create(ctx context.Context, user *biz.User) (*biz.User, error) {
	po, err := r.data.uc.Create(ctx, &userV1.CreateRequest{User: &userV1.UserStruct{
		Id:           user.Id,
		Email:        user.Email,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
	}})
	if err != nil {
		return nil, interfaceV1.ErrorUsernameConflict("create user conflict")
	}

	u := po.GetUser()

	return &biz.User{
		Id:           u.GetId(),
		Email:        u.GetEmail(),
		Username:     u.GetUsername(),
		PasswordHash: u.GetPasswordHash(),
	}, nil
}

func (r *userRepo) Update(ctx context.Context, user *biz.User) (*biz.User, error) {
	po, err := r.data.uc.Update(ctx, &userV1.UpdateRequest{User: &userV1.UserStruct{
		Id:           user.Id,
		Email:        user.Email,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
	}})
	if err != nil {
		return nil, interfaceV1.ErrorUsernameConflict("update user conflict")
	}

	u := po.GetUser()

	return &biz.User{
		Id:           u.GetId(),
		Email:        u.GetEmail(),
		Username:     u.GetUsername(),
		PasswordHash: u.GetPasswordHash(),
	}, nil
}

func (r *userRepo) Get(ctx context.Context, id int64) (*biz.User, error) {
	po, err := r.data.uc.Get(ctx, &userV1.GetRequest{Id: id})
	if err != nil {
		return nil, interfaceV1.ErrorNotFoundError("find id: %s not found, err: %v", id, err)
	}

	u := po.GetUser()

	return &biz.User{
		Id:           u.GetId(),
		Email:        u.GetEmail(),
		Username:     u.GetUsername(),
		PasswordHash: u.GetPasswordHash(),
	}, nil
}
