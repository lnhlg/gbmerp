package model

type DeptModel struct {
	ID   int64  `gorm:"column:ID;primaryKey"`
	No   string `gorm:"column:Department_ID"`
	Name string `gorm:"column:Department_Name"`
}

func (d DeptModel) TableName() string {
	return "tbsf_department"
}
