package data

import (
	"context"
	"gbmerp/app/user/internal/biz"
	"gbmerp/app/user/internal/data/model"
	"github.com/go-kratos/kratos/v2/errors"
	xerrors "github.com/pkg/errors"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (ur *userRepo) GetUser(
	ctx context.Context,
	userName string,
) (*biz.User, error) {

	us := model.UserModel{Name: userName}

	if err := ur.data.db.
		WithContext(ctx).
		Find(&us).Error;
		err != nil {

		err = errors.InternalServer("GETUSER_FAILED", err.Error())
		return nil, xerrors.Wrap(err, "")
	}

	return &biz.User{
		ID:       us.ID,
		Name:     us.Name,
		Password: us.Password,
		StaffID:  us.StaffID,
		UserType: us.Type,
	}, nil
}
