package whitebit

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Whitebit struct {
	apiKey     string
	apiSecret  string
	baseURL    string
	httpClient *http.Client
}

type Client interface {
	SendRequest(endpoint Endpoint) ([]byte, error)
}

func NewClient(apiKey string, apiSecret string) *Whitebit {
	return &Whitebit{
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		baseURL:    "https://whitebit.com",
		httpClient: &http.Client{Timeout: 15 * time.Second},
	}
}

// NewClientWithHTTPClient allows constructing a Whitebit client with a custom http.Client and baseURL.
// If httpClient is nil, a default one will be used.
func NewClientWithHTTPClient(apiKey, apiSecret string, httpClient *http.Client, baseURL string) *Whitebit {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 15 * time.Second}
	}
	if baseURL == "" {
		baseURL = "https://whitebit.com"
	}
	return &Whitebit{
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

func (c *Whitebit) call(request *http.Request) ([]byte, int, error) {
	client := c.httpClient

	response, err := client.Do(request)
	if err != nil {
		return nil, 0, fmt.Errorf("network error: %w", err)
	}
	defer response.Body.Close()

	//receiving data
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, fmt.Errorf("failed to read response body: %w", err)
	}

	return responseBody, response.StatusCode, nil
}

func (c *Whitebit) SendRequest(endpoint Endpoint) ([]byte, error) {
	url := c.baseURL + endpoint.Url()

	var req *http.Request
	var err error

	if endpoint.IsAuthed() {
		requestBody, err := json.Marshal(endpoint)

		if err != nil {
			return nil, err
		}

		req, err = CreateAuthedRequest(url, requestBody, c.apiKey, c.apiSecret)

		if err != nil {
			return nil, err
		}
	} else {
		req, err = CreateRequest(url)

		if err != nil {
			return nil, err
		}
	}

	response, status, err := c.call(req)

	if err != nil {
		return nil, err
	}

	if status != http.StatusOK && status != http.StatusCreated {
		var validationError Error
		if err := json.Unmarshal(response, &validationError); err != nil {
			// Return more informative error with raw response
			return nil, fmt.Errorf("API error (status %d): %s (failed to parse error response: %w)",
				status, string(response), err)
		}
		return nil, validationError
	}

	return response, nil
}
