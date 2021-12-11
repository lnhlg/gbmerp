package biz

import (
	"context"
)

type User struct {
	ID       string
	Name     string
	UserType string
	Staff    *Staff
}

type Staff struct {
	ID     int64
	No     string
	Name   string
	Gender string
	Duty   *Duty
	Dept   *Department
}

type Duty struct {
	ID    int64
	No    string
	Name  string
	Level int64
}

type Department struct {
	ID   int64
	No   string
	Name string
}

type UserRepo interface {
	Login(context.Context, string, string) (int64, error)
}

type UserUseCase struct {
	repo   UserRepo
	authUc *AuthUseCase
}

func NewUserUseCase(repo UserRepo, authUc *AuthUseCase) *UserUseCase {
	return &UserUseCase{
		repo:   repo,
		authUc: authUc,
	}
}

func (us *UserUseCase) Login(
	ctx context.Context,
	usName string,
	pwd string,
) (*Token, error) {
	expire, err := us.repo.Login(ctx, usName, pwd)
	if err != nil {
		return nil, err
	}

	return us.authUc.Auth(usName, expire)
}
