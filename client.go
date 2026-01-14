package whitebit

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Whitebit struct {
	ApiKey     string
	ApiSecret  string
	BaseURL    string
	httpClient *http.Client
}

type Client interface {
	SendRequest(endpoint Endpoint) ([]byte, error)
}

func NewClient(apiKey string, apiSecret string) *Whitebit {
	return &Whitebit{
		ApiKey:     apiKey,
		ApiSecret:  apiSecret,
		BaseURL:    "https://whitebit.com",
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
		ApiKey:     apiKey,
		ApiSecret:  apiSecret,
		BaseURL:    baseURL,
		httpClient: httpClient,
	}
}

func (c *Whitebit) call(request *http.Request) ([]byte, int, error) {
	client := c.httpClient

	response, err := client.Do(request)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer response.Body.Close()

	//receiving data
	responseBody, err := io.ReadAll(response.Body)

	return responseBody, response.StatusCode, err
}

func (c *Whitebit) SendRequest(endpoint Endpoint) ([]byte, error) {
	url := c.BaseURL + endpoint.Url()

	var req *http.Request
	var err error

	if endpoint.IsAuthed() {
		requestBody, err := json.Marshal(endpoint)

		if err != nil {
			return nil, err
		}

		req, err = CreateAuthedRequest(url, requestBody, c.ApiKey, c.ApiSecret)

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
		_ = json.Unmarshal(response, &validationError)
		return nil, validationError
	}

	return response, nil
}
