package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/SAgamyradov/yandexService.git/internal/app/service"
)

// POST request
func ShortenURL(c *gin.Context, urlService *service.URLServiceImpl) {
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

	shortURL, err := urlService.ShortenURL(longURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error saving URL"})
		return
	}

	c.Header("Content-Type", "text/plain")
	c.String(http.StatusCreated, shortURL)
}

// GET request
func Redirect(c *gin.Context, urlService *service.URLServiceImpl) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
		return
	}

	longURL, err := urlService.GetLongURL(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed GET long URL"})
		return
	}

	c.Writer.Header().Set("Location", longURL)
	fmt.Printf("Setting Location header to: %s\n", longURL)
	c.Writer.WriteHeader(http.StatusTemporaryRedirect)
	fmt.Printf("Writing status code: %d\n", http.StatusTemporaryRedirect) // Add logging

}
