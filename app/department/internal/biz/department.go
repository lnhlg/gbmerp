package biz

import "context"

type Department struct {
	ID   int64
	No   string
	Name string
}

type DeptRepo interface {
	GetDepartment(context.Context, int64) (*Department, error)
}

type DeptUsecase struct {
	repo DeptRepo
}

func NewDeptUsecase(repo DeptRepo) *DeptUsecase {
	return &DeptUsecase{
		repo: repo,
	}
}

func (uc *DeptUsecase) GetDepartment(
	ctx context.Context,
	id int64,
) (*Department, error) {

	return uc.repo.GetDepartment(ctx, id)
}
