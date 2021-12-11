package data

import (
	"context"
	"gbmerp/app/duty/internal/biz"
	"gbmerp/app/duty/internal/data/model"
	"github.com/go-kratos/kratos/v2/errors"
	xerrors "github.com/pkg/errors"
)

type dutyRepo struct {
	data *Data
}

func NewDutyRepo(data *Data) biz.DutyRepo {
	return &dutyRepo{
		data: data,
	}
}

func (dp *dutyRepo) GetDuty(
	ctx context.Context,
	id int64,
) (*biz.Duty, error) {

	duty := model.DutyModel{
		ID: id,
	}
	if err := dp.data.db.
		WithContext(ctx).
		Find(&duty).Error;
		err != nil {

		err = errors.InternalServer("GETDUTY_FAILED", err.Error())
		return nil, xerrors.Wrap(err, "")
	}

	return &biz.Duty{
		ID:    duty.ID,
		No:    duty.No,
		Name:  duty.Name,
		Level: duty.Level,
	}, nil
}
