package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type GormDB struct {
	db *gorm.DB
}

func NewGormDB(driver, dsn string) *gorm.DB {
	var dialector gorm.Dialector
	switch driver {
	case "mysql":
		dialector = mysql.Open(dsn)
	case "mssql":
		dialector = sqlserver.Open(dsn)
	default:
		dialector = sqlserver.Open(dsn)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = db.Debug()

	return db
}
