package config

import "flag"

var (
	Addr    string
	BaseURL string
)

func ConfigInit() {
	flag.StringVar(&Addr, "a", "localhost:8888", "Addr for starting server")
	flag.StringVar(&BaseURL, "b", "http://localhost:9999", "short basic URL")

	flag.Parse()
}
