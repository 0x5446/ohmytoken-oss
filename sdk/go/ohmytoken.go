// Package ohmytoken provides a lightweight client for tracking LLM token consumption.
package ohmytoken

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const DefaultEndpoint = "https://api.ohmytoken.dev/api/v1/ingest"

type Client struct {
	APIKey   string
	Endpoint string
	client   *http.Client
}

type TrackRequest struct {
	Model            string  `json:"model"`
	PromptTokens     int     `json:"prompt_tokens"`
	CompletionTokens int     `json:"completion_tokens"`
	Cost             float64 `json:"cost,omitempty"`
	Label            string  `json:"label,omitempty"`
}

type TrackResponse struct {
	Status      string `json:"status"`
	TotalTokens int    `json:"total_tokens,omitempty"`
	BoardTotal  int64  `json:"board_total,omitempty"`
}

func New(apiKey string) *Client {
	return &Client{
		APIKey:   apiKey,
		Endpoint: DefaultEndpoint,
		client:   &http.Client{Timeout: 5 * time.Second},
	}
}

func (c *Client) Track(model string, promptTokens, completionTokens int) (*TrackResponse, error) {
	return c.TrackWithOptions(TrackRequest{
		Model:            model,
		PromptTokens:     promptTokens,
		CompletionTokens: completionTokens,
	})
}

func (c *Client) TrackWithOptions(req TrackRequest) (*TrackResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("X-API-Key", c.APIKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TrackResponse
	json.NewDecoder(resp.Body).Decode(&result)
	return &result, nil
}
