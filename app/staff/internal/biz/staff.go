package biz

import "context"

type Staff struct {
	ID     int64
	No     string
	Name   string
	Gender string
	DeptID int64
	DutyID int64
}

type StaffRepo interface {
	GetStaff(context.Context, string) (*Staff, error)
}

type StaffUsecase struct {
	repo StaffRepo
}

func NewStaffUsecase(repo StaffRepo) *StaffUsecase {
	return &StaffUsecase{
		repo: repo,
	}
}

func (uc *StaffUsecase) GetStaff(
	ctx context.Context,
	no string,
) (*Staff, error) {
	return uc.repo.GetStaff(ctx, no)
}
