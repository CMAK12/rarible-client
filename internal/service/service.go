package service

import (
	"inforce-task/internal/dto"

	"go.uber.org/zap"
)

type Client interface {
	GetOwnershipsByID(id string) (map[string]interface{}, error)
	GetTraitRarities(body dto.CollectionTrait) (map[string]interface{}, error)
}

type Service struct {
	client Client
	logger *zap.Logger
}

func NewService(c Client, log *zap.Logger) *Service {
	return &Service{
		client: c,
		logger: log,
	}
}
