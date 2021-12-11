package service

import (
	"context"
	"gbmerp/app/duty/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	pb "gbmerp/api/duty/v1"
)

type DutyService struct {
	pb.UnimplementedDutyServer

	uc  *biz.DutyUsecase
	log *log.Helper
}

func NewDutyService(
	uc *biz.DutyUsecase,
	logger log.Logger,
) *DutyService {
	return &DutyService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (ds *DutyService) GetDuty(
	ctx context.Context,
	req *pb.GetDutyRequest,
) (*pb.GetDutyReply, error) {

	ds.log.WithContext(ctx).Infof("start call DutyService.GetDuty: %v", req)

	if req.GetId() == 0 {
		err := errors.BadRequest("DUTYID_MISSING", "missing parameter")
		ds.log.WithContext(ctx).Error(err)
		return nil, err
	}

	duty, err := ds.uc.GetDuty(ctx, req.GetId())
	if err != nil {
		ds.log.WithContext(ctx).Error(err)
		return nil, err
	}

	return &pb.GetDutyReply{
		Id:    duty.ID,
		No:    duty.No,
		Name:  duty.Name,
		Level: duty.Level,
	}, nil
}
