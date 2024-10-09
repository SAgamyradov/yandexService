package config

import "flag"

type Config struct {
	Addr    string
	BaseURL string
}

func InitConfig() *Config {
	addr := flag.String("a", "localhost:8888", "Port for starting HTTP-server")
	baseURL := flag.String("b", "http://localhost:8000", "Basic addr for short URL")

	flag.Parse()

	return &Config{
		Addr:    *addr,
		BaseURL: *baseURL,
	}
}
