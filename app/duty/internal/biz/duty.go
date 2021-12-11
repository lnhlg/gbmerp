package biz

import "context"

type Duty struct {
	ID    int64
	No    string
	Name  string
	Level int64
}

type DutyRepo interface {
	GetDuty(context.Context, int64) (*Duty, error)
}

type DutyUsecase struct {
	repo DutyRepo
}

func NewDutyUsecase(repo DutyRepo) *DutyUsecase {
	return &DutyUsecase{
		repo: repo,
	}
}

func (uc *DutyUsecase) GetDuty(
	ctx context.Context,
	id int64,
) (*Duty, error) {
	return uc.repo.GetDuty(ctx, id)
}
