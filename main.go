package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tharun-d/blog/core"
	"github.com/tharun-d/blog/handlers"
	"github.com/tharun-d/blog/repository"
	"github.com/tharun-d/blog/service"
)

func main() {
	config, err := core.NewConfig()
	if err != nil {
		log.Fatalf("error while getting data from config: %s", err.Error())
	}

	db, err := repository.NewRepository(config)
	if err != nil {
		log.Fatalf("error while connecting to db: %s", err.Error())
	}

	svc := service.NewService(db)

	router := handlers.NewHandlers(svc)
	r := mux.NewRouter()

	r.HandleFunc("/articles", router.SaveBlog).Methods(http.MethodPost)
	r.HandleFunc("/articles/{id}", router.GetBlogByID).Methods(http.MethodGet)
	r.HandleFunc("/articles", router.GetAll).Methods(http.MethodGet)

	log.Printf("server started at %s", config.Listener)
	http.ListenAndServe(config.Listener, r)

}
