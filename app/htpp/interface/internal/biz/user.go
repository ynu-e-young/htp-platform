package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	interfaceV1 "github.com/ynu-e-young/apis-go/htpp/htpp/interface/v1"

	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/conf"
	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/pkg/middleware/auth"
	"github.com/ynu-e-young/htp-platform/pkg"
)

type User struct {
	Id           int64
	Email        string
	Username     string
	PasswordHash string
	Token        string
}

type UserRepo interface {
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Get(ctx context.Context, id int64) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
	jc   *conf.Jwt

	log *log.Helper
}

func NewUserUsecase(repo UserRepo, jc *conf.Jwt, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		jc:   jc,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUsecase) generateTokenHeader(username string) (string, error) {
	return auth.GenerateToken(uc.jc.GetSecret(), username)
}

func (uc *UserUsecase) Register(ctx context.Context, username, email, password string) (*User, error) {
	ph, err := pkg.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Email:        email,
		Username:     username,
		PasswordHash: ph,
	}

	u, err := uc.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	token, err := uc.generateTokenHeader(username)
	if err != nil {
		return nil, err
	}

	u.Token = token
	return u, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*User, error) {
	u, err := uc.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !pkg.VerifyPassword(u.PasswordHash, password) {
		return nil, interfaceV1.ErrorLoginFailed("Verify password failed")
	}

	token, err := uc.generateTokenHeader(u.Username)
	if err != nil {
		return nil, err
	}

	u.Token = token
	return u, nil
}

func (uc *UserUsecase) Get(ctx context.Context, id int64) (*User, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *UserUsecase) Update(ctx context.Context, user *User) (*User, error) {
	ph, err := pkg.HashPassword(user.PasswordHash)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = ph
	u, err := uc.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	token, err := uc.generateTokenHeader(u.Username)
	if err != nil {
		return nil, err
	}

	u.Token = token
	return u, nil
}
