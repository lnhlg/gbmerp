package data

import (
	"context"

	"gbmerp/app/department/internal/biz"
	"gbmerp/app/department/internal/data/model"
	"github.com/go-kratos/kratos/v2/errors"
	xerrors "github.com/pkg/errors"
)

type deptRepo struct {
	data *Data
}

func NewDeptRepo(data *Data) biz.DeptRepo {
	return &deptRepo{
		data: data,
	}
}

func (dp *deptRepo) GetDepartment(
	ctx context.Context,
	id int64,
) (*biz.Department, error) {

	dpt := model.DeptModel{
		ID: id,
	}
	if err := dp.data.db.
		WithContext(ctx).
		Find(&dpt).Error;
		err != nil {

		err = errors.InternalServer("GETDEPARTMENT_FAILED", err.Error())
		return nil, xerrors.Wrap(err, "")
	}

	return &biz.Department{
		ID:   dpt.ID,
		No:   dpt.No,
		Name: dpt.Name,
	}, nil
}
