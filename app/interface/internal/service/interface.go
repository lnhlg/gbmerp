package service

import (
	"context"
	pb "gbmerp/api/interface/v1"
	"gbmerp/app/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type InterfaceService struct {
	pb.UnimplementedInterfaceServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewInterfaceService(uc *biz.UserUseCase) *InterfaceService {
	return &InterfaceService{
		uc: uc,
	}
}

func (s *InterfaceService) Login(
	ctx context.Context,
	req *pb.LoginReq,
) (*pb.LoginReply, error) {

	if req.GetUsername() == "" || req.GetPassword() == "" {

		err := errors.BadRequest("Login", "missing parameter")
		s.log.WithContext(ctx).Error(err)
		return nil, err
	}

	token, err := s.uc.Login(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {

		s.log.WithContext(ctx).Error(err)
		return nil, err
	}

	return &pb.LoginReply{
		Token:  token.Ciphertext,
		Expire: token.Expire,
	}, nil
}
