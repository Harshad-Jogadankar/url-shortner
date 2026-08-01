package main

import (
	"log"
	"net/http"
	"url-shortner/internal/controller"
	"url-shortner/internal/repository"
	"url-shortner/internal/router"
	"url-shortner/internal/service"
)

func main() {
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	controller := controller.NewController(svc)
	mux := router.NewRouter(controller)

	log.Println("server starting")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
