package biz

import (
	"context"
)

type User struct {
	ID       string
	Name     string
	Password string
	StaffID  string
	UserType string
}

type UserRepo interface {
	GetUser(context.Context, string) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (uc *UserUsecase) GetUser(ctx context.Context, name string) (*User, error) {
	return uc.repo.GetUser(ctx, name)
}
