package service

import (
	"context"
	"gbmerp/app/staff/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	pb "gbmerp/api/staff/v1"
)

type StaffService struct {
	pb.UnimplementedStaffServer

	uc  *biz.StaffUsecase
	log *log.Helper
}

func NewStaffService(
	uc *biz.StaffUsecase,
	logger log.Logger,
) *StaffService {
	return &StaffService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *StaffService) GetStaff(
	ctx context.Context,
	req *pb.GetStaffRequest,
) (*pb.GetStaffReply, error) {

	s.log.WithContext(ctx).Infof("start call StaffService.GetStaff: %v", req)

	if req.GetStaffno() == "" {

		err := pb.ErrorStaffnoMissing("missing parameter")
		s.log.WithContext(ctx).Error(err)
		return nil, err
	}

	staff, err := s.uc.GetStaff(ctx, req.GetStaffno())
	if err != nil {

		s.log.WithContext(ctx).Error(err)
		return nil, err
	}

	return &pb.GetStaffReply{
		Id:     staff.ID,
		No:     staff.No,
		Name:   staff.Name,
		Gender: staff.Gender,
		Deptid: staff.DeptID,
		Dutyid: staff.DutyID,
	}, nil
}
