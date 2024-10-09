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

	cfg := config.InitConfig()

	repo := repository.NewInMemoryStorage()

	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		handler.ShortenURL(c, repo)
	})
	r.GET("/:id", func(c *gin.Context) {
		handler.Redirect(c, repo)
	})

	fmt.Printf("Started server on http://%s\n", cfg.Addr)
	log.Fatal(r.Run(cfg.Addr))
}
