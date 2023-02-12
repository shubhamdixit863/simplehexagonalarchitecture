package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"task4/cmd/web/handlers"
	"task4/cmd/web/models"
	"task4/cmd/web/repository"
)

func main() {
	repo := &repository.Stub{Items: []models.Item{}}
	us := handlers.Handler{
		Repo: repo,
	}
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	h := c.Handler(r)
	r.HandleFunc("/items", us.GetItems)

	address := "localhost"
	port := "8080"
	log.Println("Server Starting .... At Port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), h))

}
