package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id           int64
	Email        string
	Username     string
	PasswordHash string
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

	log *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUsecase) FindByUsername(ctx context.Context, username string) (*User, error) {
	return uc.repo.FindByUsername(ctx, username)
}

func (uc *UserUsecase) FindByEmail(ctx context.Context, email string) (*User, error) {
	return uc.repo.FindByEmail(ctx, email)
}

func (uc *UserUsecase) Create(ctx context.Context, user *User) (*User, error) {
	return uc.repo.Create(ctx, user)
}

func (uc *UserUsecase) Update(ctx context.Context, user *User) (*User, error) {
	return uc.repo.Update(ctx, user)
}

func (uc *UserUsecase) Get(ctx context.Context, id int64) (*User, error) {
	return uc.repo.Get(ctx, id)
}
