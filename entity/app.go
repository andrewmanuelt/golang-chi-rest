package entity

import "gorm.io/gorm"

type App struct {
	gorm.Model

	AppName string `gorm:"column:app_name"`
	AppVer  string `gorm:"column:app_ver"`
}
