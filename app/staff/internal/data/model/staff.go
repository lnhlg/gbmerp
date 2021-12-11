package model

type StaffModel struct {
	ID       int64  `gorm:"column:ID;primaryKey"`
	No       string `gorm:"column:Staff_ID"`
	Name     string `gorm:"column:Staff_Name"`
	Gender   string `gorm:"column:Staff_Sex"`
	Status   string `gorm:"column:Staff_Status"`
	DeptID   int64  `gorm:"column:Department_ID1"`
	DutyID   int64  `gorm:"column:Staff_Duty_ID3"`
	ShStatus int64  `gorm:"column:Sh_Status"`
}

func (s StaffModel) TableName() string {
	return "tbsf_staff"
}
