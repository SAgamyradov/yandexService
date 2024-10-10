package config

import (
	"flag"
	"strings"
)

type Config struct {
	Address string
	BaseURL string
}

func InitConfig() *Config {
	addr := flag.String("a", "localhost:8888", "Адрес запуска HTTP-сервера")
	baseURL := flag.String("b", "http://localhost:8888", "Базовый адрес сокращенного URL")

	flag.Parse()
	if !strings.HasSuffix(*baseURL, "/") {
		*baseURL += "/"
	}
	return &Config{
		Address: *addr,
		BaseURL: *baseURL,
	}
}
