package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

type APIClient struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

func NewAPIClient(baseURL, apiKey string) *APIClient {
	transport := &http2.Transport{}

	return &APIClient{
		BaseURL: baseURL,
		APIKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout:   60 * time.Second,
			Transport: transport,
		},
	}
}

func (c *APIClient) makeRequest(method, endpoint string, body any) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)

	var requestBody []byte
	var err error
	if body != nil {
		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.APIKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("status: %d, error: %s", resp.StatusCode, resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func (c *APIClient) Get(endpoint string) ([]byte, error) {
	return c.makeRequest(http.MethodGet, endpoint, nil)
}

func (c *APIClient) Post(endpoint string, body any) ([]byte, error) {
	return c.makeRequest(http.MethodPost, endpoint, body)
}

func (c *APIClient) Delete(endpoint string, payload map[string]string) ([]byte, error) {
	return c.makeRequest(http.MethodPost, endpoint, payload)
}
