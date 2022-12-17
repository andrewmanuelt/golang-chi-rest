package service

import (
	"golang-chi-rest/model"
	"golang-chi-rest/repository"
)

type AppService interface {
	Get() model.AppResponse
}

type AppServiceImpl struct {
	repository repository.AppRepository
}

func (service AppServiceImpl) Get() model.AppResponse {
	result := service.repository.Get()

	return model.AppResponse{
		AppName: result.AppName,
		AppVer:  result.AppVer,
	}
}

func NewAppService(repository repository.AppRepository) AppService {
	return AppServiceImpl{
		repository: repository,
	}
}
