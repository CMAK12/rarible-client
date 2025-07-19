package client_test

import (
	"bytes"
	"encoding/json"
	"inforce-task/internal/client"
	"inforce-task/internal/config"
	"inforce-task/internal/dto"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func newTestClient(handler http.HandlerFunc) *client.Client {
	ts := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) *http.Response {
			w := &fakeResponseWriter{HeaderMap: http.Header{}}
			handler(w, req)
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(w.Body.Bytes())),
				Header:     w.HeaderMap,
			}
		}),
	}
	cfg := &config.RaribleAPIConfig{
		BaseURL: "https://api.rarible.org",
		APIKey:  "test-key",
	}
	logger := zap.NewNop()
	return client.NewClient(ts, logger, cfg)
}

func TestGetOwnershipsByID_Success(t *testing.T) {
	expected := map[string]interface{}{"key": "value"}

	c := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "test-key", r.Header.Get("X-API-KEY"))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expected)
	})

	result, err := c.GetOwnershipsByID("some-id")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGetTraitRarities_Success(t *testing.T) {
	expected := map[string]interface{}{"rarity": "rare"}

	c := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expected)
	})

	body := dto.CollectionTrait{
		CollectionID: "my-collection",
		Properties:   []dto.TraitProperty{{Key: "color", Value: "blue"}},
	}

	result, err := c.GetTraitRarities(body)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

type fakeResponseWriter struct {
	HeaderMap http.Header
	Body      bytes.Buffer
}

func (f *fakeResponseWriter) Header() http.Header         { return f.HeaderMap }
func (f *fakeResponseWriter) Write(b []byte) (int, error) { return f.Body.Write(b) }
func (f *fakeResponseWriter) WriteHeader(statusCode int)  {}
