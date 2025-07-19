package app

import (
	"fmt"
	"net/http"

	"inforce-task/internal/client"
	"inforce-task/internal/config"
	http_v1 "inforce-task/internal/controller/http/v1"
	"inforce-task/internal/service"

	"go.uber.org/zap"
)

func MustRun() error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}

	cfg, err := config.Load()
	if err != nil {
		logger.Error("Failed to load config", zap.Error(err))
		return err
	}

	httpClient := &http.Client{
		Timeout: cfg.HTTPClient.Timeout,
		Transport: &http.Transport{
			IdleConnTimeout: cfg.HTTPClient.ConnectTimeout,
			MaxIdleConns:    cfg.HTTPClient.MaxIdleConns,
		},
	}

	client := client.NewClient(httpClient, logger, &cfg.Rarible)
	service := service.NewService(client, logger)
	handler := http_v1.NewHandler(service, logger)

	app := handler.InitRoutes()

	return app.Listen(fmt.Sprintf(":%s", cfg.Server.Port))
}
