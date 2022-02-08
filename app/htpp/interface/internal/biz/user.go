package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"htp-platform/app/htpp/interface/internal/conf"
	"htp-platform/app/htpp/interface/internal/pkg/middleware/auth"
	"htp-platform/pkg/hash"
)

const (
	// unauthorized holds the unauthorized reason.
	unauthorized string = "UNAUTHORIZED"

	// notFound holds the notFound reason.
	notFound string = "NOTFOUND"
)

var (
	ErrVerifyFailed = errors.Unauthorized(unauthorized, "Verify password failed")
	ErrUserNotFound = errors.NotFound(notFound, "user not found")
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
	ph, err := hash.HashPassword(password)
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

	if !hash.VerifyPassword(u.PasswordHash, password) {
		return nil, ErrVerifyFailed
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
