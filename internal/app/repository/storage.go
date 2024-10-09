package repository

import "fmt"

type InMemoryStorage struct {
	urlMap map[string]string
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		urlMap: make(map[string]string),
	}
}

func (s *InMemoryStorage) GenerateShortURL(longURL string) (string, error) {
	shortURL := fmt.Sprintf("%d", len(s.urlMap)+1)
	s.urlMap[shortURL] = longURL
	return shortURL, nil
}

func (s *InMemoryStorage) GetLongURL(shortURL string) (string, error) {
	longURL, ok := s.urlMap[shortURL]
	if !ok {
		return "", fmt.Errorf("short URL not found")
	}
	return longURL, nil
}
