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
//	    option.WithAPIKey("YOUR_API_KEY"),
//	    option.WithTxcPayload(""),
//	    option.WithTxcSignature(""),
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

	// request/nonce: fill if falsy, not just if absent. Some endpoints' generated
	// request structs declare request/nonce as required (OpenAPI marks them
	// required for the wire contract, since they're mandatory for non-SDK
	// callers), so Fern emits non-optional fields that always serialize as their
	// zero value ("", 0) when the caller doesn't set them — which a strict
	// "== nil" presence check would never catch. "" and 0 are never legitimate
	// caller-supplied values for these two fields, so treating them as absent is
	// safe and still lets an explicitly-set nonce survive.
	if v, ok := body["request"].(string); !ok || v == "" {
		body["request"] = req.URL.Path
	}
	if v, ok := body["nonce"].(float64); !ok || v == 0 {
		body["nonce"] = time.Now().UnixMilli()
	}
	// nonceWindow is genuinely optional in every spec and defaults to false
	// server-side — an explicit `nonceWindow: false` from the caller must
	// survive, so only fill it in when the key is missing entirely.
	if _, ok := body["nonceWindow"]; !ok {
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
