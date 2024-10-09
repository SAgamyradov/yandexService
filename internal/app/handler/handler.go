package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/SAgamyradov/yandexService.git/internal/app/repository"
	"github.com/gin-gonic/gin"
)

// POST request
func ShortenURL(c *gin.Context, repo repository.URLRepository) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect method"})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading"})
		return
	}

	longURL := strings.TrimSpace(string(body))
	if _, err = url.ParseRequestURI(longURL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect URL"})
		return
	}

	shortURL, err := repo.GenerateShortURL(longURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error saving URL"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"shortURL": fmt.Sprintf("http://localhost:8080/%s", shortURL)})
}

// GET request
func Redirect(c *gin.Context, repo repository.URLRepository) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
		return
	}

	longURL, err := repo.GetLongURL(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed GET long URL"})
		return
	}

	c.Writer.Header().Set("Location", longURL)
	c.Writer.WriteHeader(http.StatusTemporaryRedirect)
}
