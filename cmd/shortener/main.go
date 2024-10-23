package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/SAgamyradov/yandexService.git/internal/app/config"
	handler "github.com/SAgamyradov/yandexService.git/internal/app/handler"
	"github.com/SAgamyradov/yandexService.git/internal/app/repository"
	"github.com/SAgamyradov/yandexService.git/internal/app/service"
)

func main() {

	appConfig := config.InitConfig()

	repo := repository.NewInMemoryStorage()

	urlService := service.NewURLService(repo, appConfig)

	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		handler.ShortenURL(c, urlService)
	})
	r.GET("/:id", func(c *gin.Context) {
		handler.Redirect(c, urlService)
	})

	if err := r.Run(appConfig.Address); err != nil {
		fmt.Println(err)
	}

}
