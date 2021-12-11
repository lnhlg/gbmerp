package data

import (
	"context"

	pb "gbmerp/api/staff/v1"

	"gbmerp/app/staff/internal/biz"
	"gbmerp/app/staff/internal/data/model"
	xerrors "github.com/pkg/errors"
)

type staffRepo struct {
	data *Data
}

func NewStaffRepo(data *Data) biz.StaffRepo {
	return &staffRepo{
		data: data,
	}
}

func (sr *staffRepo) GetStaff(
	ctx context.Context,
	no string,
) (*biz.Staff, error) {

	staff := model.StaffModel{
		No:       no,
		Status:   "2",
		ShStatus: 1,
	}
	if err := sr.data.db.
		WithContext(ctx).
		Where(&staff).
		Find(&staff).Error; err != nil {

		err = pb.ErrorGetstaffFailed(err.Error())
		return nil, xerrors.Wrap(err, "")
	}

	return &biz.Staff{
		ID:     staff.ID,
		No:     staff.No,
		Name:   staff.Name,
		Gender: staff.Gender,
		DeptID: staff.DeptID,
		DutyID: staff.DutyID,
	}, nil
}
