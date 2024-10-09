package main

import (
	"fmt"
	"log"

	"github.com/SAgamyradov/yandexService.git/internal/app/handler"
	"github.com/SAgamyradov/yandexService.git/internal/app/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	repo := repository.NewInMemoryStorage()

	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		handler.ShortenURL(c, repo)
	})
	r.GET("/:id", func(c *gin.Context) {
		handler.Redirect(c, repo)
	})

	fmt.Println("Started server on http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
