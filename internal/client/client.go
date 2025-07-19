package client

import (
	"inforce-task/internal/config"
	"net/http"

	"go.uber.org/zap"
)

type Client struct {
	client         *http.Client
	logger         *zap.Logger
	raribleBaseURL string
	raribleAPIKey  string
}

func NewClient(client *http.Client, logger *zap.Logger, rarible *config.RaribleAPIConfig) *Client {
	return &Client{
		client:         client,
		logger:         logger,
		raribleBaseURL: rarible.BaseURL,
		raribleAPIKey:  rarible.APIKey,
	}
}
