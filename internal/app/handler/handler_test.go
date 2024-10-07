package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SAgamyradov/yandexService.git/internal/app/repository"
)

var mockRepo *repository.InMemoryStorage

func TestShortenURL(t *testing.T) {
	t.Run("Should return 201 Created and shortURL", func(t *testing.T) {

		longURL := "https://www.google.com"
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(longURL))
		w := httptest.NewRecorder()

		ShortenURL(mockRepo).ServeHTTP(w, req) //  DI for mockRepo and ServeHTTP

		if w.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
		}
		if !strings.Contains(w.Body.String(), "http://localhost:8080/") {
			t.Errorf("Response body should contain shortURL")
		}
	})
	t.Run("Should return 400 Bad Request for invalid URL", func(t *testing.T) {

		longURL := "invalid-url"
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(longURL))
		w := httptest.NewRecorder()

		ShortenURL(mockRepo).ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}

func TestRedirect(t *testing.T) {
	t.Run("Should return 307 Temporary Redirect and redirect to longURL", func(t *testing.T) {

		longURL := "https://www.google.com"
		shortURL := "1"
		mockRepo.GenerateShortURL(longURL)
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", shortURL), nil)
		w := httptest.NewRecorder()

		Redirect(mockRepo).ServeHTTP(w, req)

		if w.Code != http.StatusTemporaryRedirect {
			t.Errorf("Expected status code %d, got %d", http.StatusTemporaryRedirect, w.Code)
		}
		if w.Header().Get("Location") != longURL {
			t.Errorf("Expected Location header '%s', got '%s'", longURL, w.Header().Get("Location"))
		}
	})
	t.Run("Should return 400 Bad Request for missing id", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		Redirect(mockRepo).ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
	t.Run("Should return 400 Bad Request for non-existent id", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/non-existent-id", nil)
		w := httptest.NewRecorder()

		Redirect(mockRepo).ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}

func init() {
	mockRepo = repository.NewInMemoryStorage()
}
