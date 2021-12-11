package service

import (
	"context"
	"gbmerp/app/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	pb "gbmerp/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {

	s.log.WithContext(ctx).Infof("start call UserService.GetUser: %v", req)

	if req.GetUserName() == "" {

		err := pb.ErrorUserNameMissing("missing parameter")
		s.log.WithContext(ctx).Error(err)
		return nil, err
	}

	us, err := s.uc.GetUser(ctx, req.GetUserName())
	if err != nil {
		s.log.WithContext(ctx).Error(err)
		return nil, err
	}

	if req.GetPassword() != "" && us.Password != req.GetPassword() {

		err = pb.ErrorPasswordWrong("wrong password")
		s.log.WithContext(ctx).Error(err)
		return nil, err
	}

	return &pb.GetUserReply{
		UserID:   us.ID,
		UserName: us.Name,
		StaffID:  us.StaffID,
		UserType: us.UserType,
	}, nil
}
