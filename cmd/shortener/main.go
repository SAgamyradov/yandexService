package main

import (
	"fmt"
	"net/http"

	"github.com/SAgamyradov/yandexService.git/internal/app/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.ShortenURL)
	r.HandleFunc("/{id}", handler.Redirect)

	fmt.Println("Started server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
