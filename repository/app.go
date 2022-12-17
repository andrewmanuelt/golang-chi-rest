package repository

import (
	"golang-chi-rest/entity"

	"gorm.io/gorm"
)

type AppRepository interface {
	Get() entity.App
}

type AppRepositoryImpl struct {
	db *gorm.DB
}

// Get implements AppRepository
func (repository AppRepositoryImpl) Get() entity.App {
	var app entity.App

	repository.db.First(&app)

	return entity.App{
		AppName: app.AppName,
		AppVer:  app.AppVer,
	}
}

func NewAppRepository(db *gorm.DB) AppRepository {
	return AppRepositoryImpl{
		db: db,
	}
}
