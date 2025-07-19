package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"inforce-task/internal/dto"

	"github.com/nordew/go-errx"
	"go.uber.org/zap"
)

func (r *Client) GetOwnershipsByID(id string) (map[string]any, error) {
	url := fmt.Sprintf("%s/ownerships/%s", r.raribleBaseURL, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		r.logger.Error("client.GetOwnershipsByID: ", zap.Error(err))
		return nil, errx.NewInternal().WithDescriptionAndCause("client.GetOwnershipsByID: ", err)
	}

	req.Header.Set("X-API-KEY", r.raribleAPIKey)
	resp, err := r.client.Do(req)
	if err != nil {
		r.logger.Error("client.GetOwnershipsByID: ", zap.Error(err))
		return nil, errx.NewInternal().WithDescriptionAndCause("client.GetOwnershipsByID: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errx.NewInternal().WithDescription(
			fmt.Sprintf("client.GetOwnershipsByID: status code %d", resp.StatusCode),
		)
	}

	var result map[string]any
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		r.logger.Error("client.GetOwnershipsByID: failed to decode response", zap.Error(err))
		return nil, errx.NewInternal().WithDescriptionAndCause("client.GetOwnershipsByID: failed to decode response", err)
	}

	return result, err
}

func (r *Client) GetTraitRarities(body dto.CollectionTrait) (map[string]any, error) {
	url := fmt.Sprintf("%s/items/traits/rarity", r.raribleBaseURL)
	payloadBytes, err := json.Marshal(body)
	if err != nil {
		r.logger.Error("client.GetTraitRarities: failed to marshal body", zap.Error(err))
		return nil, err
	}
	payload := strings.NewReader(string(payloadBytes))

	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		r.logger.Error("client.GetTraitRarities: ", zap.Error(err))
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-API-KEY", r.raribleAPIKey)
	resp, err := r.client.Do(req)
	if err != nil {
		r.logger.Error("client.GetTraitRarities: ", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]any
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, err
}
