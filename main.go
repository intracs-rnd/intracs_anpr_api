package main

import (
	"fmt"
	"log"
	"net/http"

	// Third party
	_ "github.com/go-sql-driver/mysql"

	// Project
	"intracs_anpr_api/api"
	"intracs_anpr_api/controller"
	"intracs_anpr_api/database"
	"intracs_anpr_api/internal/env"
	"intracs_anpr_api/repository"
	"intracs_anpr_api/route"
)

func main() {
	db := database.GetDatabase()
	repos := repository.InitRepositories(db)
	controllers := controller.InitControllers(repos)
	api := api.InitApi(controllers)
	router := route.GetAll(api)

	appUrl := env.Get("APP_URL")
	if appUrl == "" {
		appUrl = "http://localhost:8881"
	}

	fmt.Printf("Starting API server at %s\n", appUrl)
	log.Fatal(
		http.ListenAndServe(appUrl, router),
	)
}
