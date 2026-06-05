package whitebitauth

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// HmacTransport is an http.RoundTripper that adds WhiteBit HMAC authentication
// headers (X-TXC-PAYLOAD, X-TXC-SIGNATURE) to every POST request.
//
// Usage:
//
//	import whitebitauth "github.com/whitebit-exchange/go-sdk/auth"
//
//	c := client.NewClient(
//	    option.WithTxcApikey("YOUR_API_KEY"),
//	    option.WithHTTPClient(whitebitauth.NewHmacClient("YOUR_API_SECRET")),
//	)
type HmacTransport struct {
	APISecret string
	Next      http.RoundTripper
}

func (t *HmacTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	next := t.Next
	if next == nil {
		next = http.DefaultTransport
	}

	if req.Method != http.MethodPost || req.Body == nil || !strings.HasPrefix(req.URL.Path, "/api/v4/") {
		return next.RoundTrip(req)
	}

	raw, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	_ = req.Body.Close()

	var body map[string]any
	if err := json.Unmarshal(raw, &body); err != nil {
		req.Body = io.NopCloser(bytes.NewReader(raw))
		return next.RoundTrip(req)
	}

	if body["nonce"] == nil {
		body["nonce"] = time.Now().UnixMilli()
	}
	if body["request"] == nil {
		body["request"] = req.URL.Path
	}
	if body["nonceWindow"] == nil {
		body["nonceWindow"] = true
	}

	newRaw, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	payload := base64.StdEncoding.EncodeToString(newRaw)
	mac := hmac.New(sha512.New, []byte(t.APISecret))
	mac.Write([]byte(payload))
	signature := fmt.Sprintf("%x", mac.Sum(nil))

	newReq := req.Clone(req.Context())
	newReq.Body = io.NopCloser(bytes.NewReader(newRaw))
	newReq.ContentLength = int64(len(newRaw))
	newReq.Header.Set("X-TXC-PAYLOAD", payload)
	newReq.Header.Set("X-TXC-SIGNATURE", signature)

	return next.RoundTrip(newReq)
}

// NewHmacClient returns an *http.Client whose transport computes HMAC headers
// for every outgoing POST request.
func NewHmacClient(apiSecret string) *http.Client {
	return &http.Client{
		Transport: &HmacTransport{APISecret: apiSecret},
	}
}
