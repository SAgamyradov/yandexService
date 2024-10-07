package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SAgamyradov/yandexService.git/internal/app/handler"
	"github.com/SAgamyradov/yandexService.git/internal/app/repository"
	"github.com/gorilla/mux"
)

func main() {

	repo := repository.NewInMemoryStorage()

	r := mux.NewRouter()

	shortenHandler := handler.ShortenURL(repo)
	redirectHandler := handler.Redirect(repo)

	r.HandleFunc("/", shortenHandler.ServeHTTP).Methods(http.MethodPost)
	r.HandleFunc("/{id}", redirectHandler.ServeHTTP).Methods(http.MethodGet)

	fmt.Println("Started server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
