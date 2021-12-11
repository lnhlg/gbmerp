package model

type UserModel struct {
	ID       string `gorm:"column:User_ID"`
	Name     string `gorm:"column:User_Name;primaryKey"`
	StaffID  string `gorm:"column:Staff_ID"`
	Password string `gorm:"column:User_Password"`
	Type     string `gorm:"column:User_Type"`
}

func (u UserModel) TableName() string {
	return "tbsf_user"
}
