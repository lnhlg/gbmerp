package data

import (
	"context"
	dpv1 "gbmerp/api/department/v1"
	duv1 "gbmerp/api/duty/v1"
	stv1 "gbmerp/api/staff/v1"
	usv1 "gbmerp/api/user/v1"
	"gbmerp/app/interface/internal/biz"
	"gbmerp/app/interface/internal/conf"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/sync/errgroup"
	"time"
)

type userRepo struct {
	data   *Data
	expire int64
}

func NewUserRepo(data *Data, conf *conf.Auth) biz.UserRepo {
	return &userRepo{
		data:   data,
		expire: conf.Expire,
	}
}

func (us *userRepo) Login(
	ctx context.Context,
	usName string,
	pwd string,
) (int64, error) {

	key := "user:" + usName

	usInfo, err := us.getUser(ctx, usName, pwd)
	if err != nil {
		return 0, err
	}

	////存为json格式
	//json := jsoniter.ConfigCompatibleWithStandardLibrary
	//usJs, err := json.Marshal(&usInfo)
	//if err != nil {
	//	return 0, err
	//}
	//
	//if err := us.data.rdb.Set(
	//	ctx,
	//	key,
	//	usJs,
	//	time.Hour*time.Duration(us.expire),
	//).Err(); err != nil {
	//	return 0, err
	//}

	//存为哈希表
	_, err = us.data.rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "ID", usInfo.ID)
		rdb.HSet(ctx, key, "Name", usInfo.Name)
		rdb.HSet(ctx, key, "UserType", usInfo.UserType)
		rdb.HSet(ctx, key+":staff", "ID", usInfo.Staff.ID)
		rdb.HSet(ctx, key+":staff", "No", usInfo.Staff.No)
		rdb.HSet(ctx, key+":staff", "Name", usInfo.Staff.Name)
		rdb.HSet(ctx, key+":staff", "Gender", usInfo.Staff.Gender)
		rdb.HSet(ctx, key+":Staff"+":duty", "ID", usInfo.Staff.Duty.ID)
		rdb.HSet(ctx, key+":Staff"+":duty", "No", usInfo.Staff.Duty.No)
		rdb.HSet(ctx, key+":Staff"+":duty", "Name", usInfo.Staff.Duty.Name)
		rdb.HSet(ctx, key+":Staff"+":duty", "Level", usInfo.Staff.Duty.Level)
		rdb.HSet(ctx, key+":Staff"+":dept", "Level", usInfo.Staff.Dept.ID)
		rdb.HSet(ctx, key+":Staff"+":dept", "Level", usInfo.Staff.Dept.No)
		rdb.HSet(ctx, key+":Staff"+":dept", "Level", usInfo.Staff.Dept.Name)
		return nil
	})
	if err != nil {
		return 0, err
	}

	return time.Now().
		Add(time.Hour * time.Duration(us.expire)).
		Unix(), nil
}

func (us *userRepo) getUser(
	ctx context.Context,
	usName string,
	pwd string,
) (*biz.User, error) {
	//获取用户
	usReply, err := us.data.usc.GetUser(ctx, &usv1.GetUserRequest{
		UserName: usName,
		Password: &pwd,
	})
	if err != nil {
		return nil, err
	}

	//获取用户关联的员工资料
	stfReply, err := us.data.stc.GetStaff(ctx, &stv1.GetStaffRequest{
		Staffno: usReply.GetStaffID(),
	})
	if err != nil {
		return nil, err
	}

	//并发获取员工职务与部门
	g, ctx := errgroup.WithContext(ctx)
	var dutReply *duv1.GetDutyReply
	var dptReply *dpv1.GetDepartmentReply

	g.Go(func() error {
		//获取该员工的职务
		dutReply, err = us.data.duc.GetDuty(ctx, &duv1.GetDutyRequest{
			Id: stfReply.GetDutyid(),
		})

		return err
	})
	g.Go(func() error {
		//获取员工所在部门
		dptReply, err = us.data.dpc.GetDepartment(ctx, &dpv1.GetDepartmentRequest{
			Id: stfReply.GetDeptid(),
		})

		return err
	})

	if err = g.Wait(); err != nil {
		return nil, err
	}

	usInfo := &biz.User{
		ID:       usReply.GetUserID(),
		Name:     usReply.GetUserName(),
		UserType: usReply.GetUserType(),
		Staff: &biz.Staff{
			ID:     stfReply.GetId(),
			No:     stfReply.GetNo(),
			Name:   stfReply.GetName(),
			Gender: stfReply.GetGender(),
			Duty: &biz.Duty{
				ID:    dutReply.GetId(),
				No:    dutReply.GetNo(),
				Name:  dutReply.GetName(),
				Level: dutReply.GetLevel(),
			},
			Dept: &biz.Department{
				ID:   dptReply.GetId(),
				No:   dptReply.GetNo(),
				Name: dptReply.GetName(),
			},
		},
	}

	return usInfo, nil
}
