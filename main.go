package main

import (
	"context"
	"fmt"
	"golang-chi-rest/config"
	"golang-chi-rest/controller"
	"golang-chi-rest/entity"
	"golang-chi-rest/helper"
	"golang-chi-rest/repository"
	"golang-chi-rest/service"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	context := context.Background()
	credis := config.RedisConnect()

	r := helper.NewRedis(context, *credis)

	r.SetRedis("helo", "ehlo")

	fmt.Println(r.GetRedis("helo"))

	db := config.DatabaseConnect()
	db.AutoMigrate(&entity.App{})

	app := chi.NewRouter()

	appRepository := repository.NewAppRepository(db)
	appService := service.NewAppService(appRepository)
	appController := controller.NewAppController(appService)

	appController.Route(app)

	log.Fatal(http.ListenAndServe(config.ConfigField("app", "url"), app))
}
