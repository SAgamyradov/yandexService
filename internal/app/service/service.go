// internal/app/service/url_service.go
package service

import (
	"fmt"

	"github.com/SAgamyradov/yandexService.git/internal/app/config"
	"github.com/SAgamyradov/yandexService.git/internal/app/repository"
)

// URLService интерфейс для обработки URLs

// URLService реализация интерфейса URLService
type URLServiceImpl struct {
	Repo    repository.URLRepository
	BaseURL string
}

// Конструктор для URLServiceImpl
func NewURLService(repo repository.URLRepository, cfg *config.Config) *URLServiceImpl {
	return &URLServiceImpl{
		Repo:    repo,
		BaseURL: cfg.BaseURL,
	}
}

func (s *URLServiceImpl) ShortenURL(longURL string) (string, error) {
	shortURLId, err := s.Repo.GenerateShortURL(longURL)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", s.BaseURL, shortURLId), nil
}

func (s *URLServiceImpl) GetLongURL(shortURL string) (string, error) {
	return s.Repo.GetLongURL(shortURL)
}
