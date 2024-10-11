package main

import (
	"fmt"

	"github.com/SAgamyradov/yandexService.git/internal/app/config"
	"github.com/SAgamyradov/yandexService.git/internal/app/handler"
	"github.com/SAgamyradov/yandexService.git/internal/app/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	appConfig := config.InitConfig()

	repo := repository.NewInMemoryStorage()

	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		handler.ShortenURL(c, repo, appConfig.BaseURL)
	})
	r.GET("/:id", func(c *gin.Context) {
		handler.Redirect(c, repo)
	})

	if err := r.Run(appConfig.Address); err != nil {
		fmt.Println(err)
	}

}
