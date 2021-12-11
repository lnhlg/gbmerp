package service

import (
	"context"
	"gbmerp/app/department/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	pb "gbmerp/api/department/v1"
)

type DeptService struct {
	pb.UnimplementedDepartmentServer

	uc  *biz.DeptUsecase
	log *log.Helper
}

func NewDeptService(
	uc *biz.DeptUsecase,
	logger log.Logger,
) *DeptService {
	return &DeptService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (ds *DeptService) GetDepartment(
	ctx context.Context,
	req *pb.GetDepartmentRequest,
) (*pb.GetDepartmentReply, error) {

	ds.log.WithContext(ctx).Infof("start call DeptService.GetDepartment: %v", req)

	if req.GetId() == 0 {
		
		err := errors.BadRequest("DEPTID_MISSING", "missing parameter")
		ds.log.WithContext(ctx).Error(err)
		return nil, err
	}

	dpt, err := ds.uc.GetDepartment(ctx, req.GetId())
	if err != nil {

		ds.log.WithContext(ctx).Error(err)
		return nil, err
	}

	return &pb.GetDepartmentReply{
		Id:   dpt.ID,
		No:   dpt.No,
		Name: dpt.Name,
	}, nil
}
