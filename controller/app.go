package controller

import (
	"golang-chi-rest/service"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type AppController struct {
	service service.AppService
}

func (controller AppController) Route(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		result := controller.service.Get()

		render.JSON(w, r, result)
	})
}

func NewAppController(service service.AppService) AppController {
	return AppController{
		service: service,
	}
}
