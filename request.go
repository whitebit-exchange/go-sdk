package whitebit

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"net/http"
)

func CreateAuthedRequest(url string, body []byte, key string, secret string) (*http.Request, error) {
	//calculating payload
	payload := base64.StdEncoding.EncodeToString(body)

	//calculating signature using sha512
	h := hmac.New(sha512.New, []byte(secret))
	h.Write([]byte(payload))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-type", "application/json")
	request.Header.Set("X-TXC-APIKEY", key)
	request.Header.Set("X-TXC-PAYLOAD", payload)
	request.Header.Set("X-TXC-SIGNATURE", signature)

	return request, nil
}

func CreateRequest(url string) (*http.Request, error) {
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-type", "application/json")

	return request, nil
}
