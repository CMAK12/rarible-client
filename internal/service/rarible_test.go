package service_test

import (
	"errors"
	"inforce-task/internal/dto"
	"inforce-task/internal/service"
	"inforce-task/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestService_GetOwnerships_Success(t *testing.T) {
	mockClient := new(mocks.Client)
	svc := service.NewService(mockClient, zap.NewNop())

	expected := map[string]interface{}{"id": "result"}
	mockClient.On("GetOwnershipsByID", "some-id").Return(expected, nil)

	result, err := svc.GetOwnerships("some-id")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockClient.AssertExpectations(t)
}

func TestService_GetOwnerships_EmptyID(t *testing.T) {
	mockClient := new(mocks.Client)
	svc := service.NewService(mockClient, zap.NewNop())

	result, err := svc.GetOwnerships("")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "missing ID param", err.Error())
}

func TestService_GetTraits_Success(t *testing.T) {
	mockClient := new(mocks.Client)
	svc := service.NewService(mockClient, zap.NewNop())

	body := dto.CollectionTrait{
		CollectionID: "abc",
		Properties:   []dto.TraitProperty{{Key: "rarity", Value: "common"}},
	}
	expected := map[string]interface{}{"rarity": "common"}

	mockClient.On("GetTraitRarities", body).Return(expected, nil)

	result, err := svc.GetTraits(body)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockClient.AssertExpectations(t)
}

func TestService_GetTraits_Error(t *testing.T) {
	mockClient := new(mocks.Client)
	svc := service.NewService(mockClient, zap.NewNop())

	body := dto.CollectionTrait{}
	mockClient.On("GetTraitRarities", body).Return(nil, errors.New("api error"))

	result, err := svc.GetTraits(body)
	assert.Error(t, err)
	assert.Nil(t, result)
}
