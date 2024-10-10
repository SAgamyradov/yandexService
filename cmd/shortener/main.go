package main

import (
	"fmt"
	"log"

	"github.com/SAgamyradov/yandexService.git/internal/app/config"
	"github.com/SAgamyradov/yandexService.git/internal/app/handler"
	"github.com/SAgamyradov/yandexService.git/internal/app/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	repo := repository.NewInMemoryStorage()
	appConfig := config.InitConfig()

	r := gin.Default()

	// Middleware to add config to context
	r.Use(func(c *gin.Context) {
		c.Set("BaseURL", appConfig.BaseURL) // Add BaseURL to the context
		c.Next()
	})

	r.POST("/", func(c *gin.Context) {
		handler.ShortenURL(c, repo)
	})
	r.GET("/:id", func(c *gin.Context) {
		handler.Redirect(c, repo)
	})

	fmt.Printf("Started server on http://%s\n", appConfig.BaseURL)
	log.Fatal(r.Run(appConfig.Address))
}
