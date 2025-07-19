package service

import (
	"inforce-task/internal/dto"

	"github.com/nordew/go-errx"
)

func (s *Service) GetOwnerships(id string) (map[string]any, error) {
	if id == "" {
		return nil, errx.NewBadRequest().WithDescription("ID is required")
	}

	return s.client.GetOwnershipsByID(id)
}

func (s *Service) GetTraits(body dto.CollectionTrait) (map[string]any, error) {
	if body.CollectionID == "" || len(body.Properties) == 0 {
		return nil, errx.NewBadRequest().WithDescription("CollectionID and Properties are required")
	}

	return s.client.GetTraitRarities(body)
}
