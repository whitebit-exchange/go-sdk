package whitebit

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Whitebit struct {
	ApiKey    string
	ApiSecret string
	BaseURL   string
}

type Client interface {
	SendRequest(endpoint Endpoint) ([]byte, error)
}

func NewClient(apiKey string, apiSecret string) *Whitebit {
	return &Whitebit{ApiKey: apiKey, ApiSecret: apiSecret, BaseURL: "https://whitebit.com"}
}

func (c *Whitebit) call(request *http.Request) ([]byte, int, error) {
	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer response.Body.Close()

	//receiving data
	responseBody, err := ioutil.ReadAll(response.Body)

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

	if status != http.StatusOK {
		var validationError Error
		_ = json.Unmarshal(response, &validationError)
		return nil, validationError
	}

	return response, nil
}
