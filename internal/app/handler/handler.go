package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/SAgamyradov/yandexService.git/internal/app/repository"
)

// POST request
func ShortenURL(repo repository.URLRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Incorrect method", http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading", http.StatusBadRequest)
			return
		}
		longURL := strings.TrimSpace(string(body))
		if _, err = url.ParseRequestURI(longURL); err != nil {
			http.Error(w, "Incorrect URL", http.StatusBadRequest)
			return
		}
		shortURL, err := repo.GenerateShortURL(longURL)
		if err != nil {
			http.Error(w, "Error saving URL", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "http://localhost:8080/%s", shortURL)
	})
}

func Redirect(repo repository.URLRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/"):]
		if id == "" {
			http.Error(w, "Missing id", http.StatusBadRequest)
			return
		}
		longURL, err := repo.GetLongURL(id)
		if err != nil {
			http.Error(w, "failed GET long URL", http.StatusBadRequest)
			return
		}
		w.Header().Set("Location", longURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
}
