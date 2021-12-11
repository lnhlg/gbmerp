package model

type DutyModel struct {
	ID    int64  `gorm:"column:ID;primaryKey"`
	No    string `gorm:"column:Duty_ID"`
	Name  string `gorm:"column:Duty_Name"`
	Level int64  `gorm:"column:Duty_Level"`
}

func (d DutyModel) TableName() string {
	return "tblBase_HR_Duty"
}
