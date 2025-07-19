package config

import "time"

type Config struct {
	Server     Server           `envPrefix:"SERVER_"`
	Rarible    RaribleAPIConfig `envPrefix:"RARIBLE_"`
	HTTPClient HTTPClientConfig `envPrefix:"HTTP_CLIENT_"`
}

type Server struct {
	Port string `env:"PORT,required"`
}

type RaribleAPIConfig struct {
	BaseURL string `env:"BASE_URL,required"`
	APIKey  string `env:"API_KEY,required"`
}

type HTTPClientConfig struct {
	Timeout        time.Duration `env:"TIMEOUT" envDefault:"30s"`
	ConnectTimeout time.Duration `env:"CONNECT_TIMEOUT" envDefault:"10s"`
	MaxIdleConns   int           `env:"MAX_IDLE_CONNS" envDefault:"100"`
}
