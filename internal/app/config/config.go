package config

import (
	"flag"
	"os"
	"strings"
)

type Config struct {
	Address string
	BaseURL string
}

func InitConfig() *Config {

	defaultAddr := "localhost:8080"
	defaultBaseURL := "http://localhost:8080/"

	addrEnv := os.Getenv("SERVER_ADDRESS")
	baseURLEnv := os.Getenv("BASE_URL")

	if addrEnv == "" {
		addrEnv = defaultAddr
	}
	if baseURLEnv == "" {
		baseURLEnv = defaultBaseURL
	}

	addr := flag.String("a", addrEnv, "Адрес запуска HTTP-сервера")
	baseURL := flag.String("b", baseURLEnv, "baseURL address for short")

	flag.Parse()
	if !strings.HasSuffix(*baseURL, "/") {
		*baseURL += "/"
	}
	return &Config{
		Address: *addr,
		BaseURL: *baseURL,
	}
}
