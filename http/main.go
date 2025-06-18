package main

import (
	"beginning_http/application"
	"log"
	"net/http"
)

func main() {
	api := &application.Api{Addr: ":8080"}
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    api.Addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /", api.GetHome)
	mux.HandleFunc("GET /users", api.GetUsers)
	mux.HandleFunc("POST /users" , api.CreateUsers)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
