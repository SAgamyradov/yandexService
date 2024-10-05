package repository

import (
	"fmt"
)

type UrlMapping struct {
	LongURL  string
	ShortURL string
}

var urlMap = make(map[string]UrlMapping)

func GenerateShortURL(longURL string) (string, error) {
	shortURL := fmt.Sprintf("%x", len(urlMap)+1)
	urlMap[shortURL] = UrlMapping{LongURL: longURL, ShortURL: shortURL}
	return shortURL, nil
}

func GetLongURL(id string) (string, error) {
	if mapping, found := urlMap[id]; found {
		return mapping.LongURL, nil
	}
	return "", nil
}
