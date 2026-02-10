package tests

import (
	"net/http"
	"reflect"
	"testing"
	"time"
	"unsafe"

	whitebit "github.com/whitebit-exchange/go-sdk"
	wbSDK "github.com/whitebit-exchange/go-sdk/sdk"
)

// getHTTPClientPtr extracts the unexported httpClient field from *whitebit.Whitebit
func getHTTPClientPtr(t *testing.T, c *whitebit.Whitebit) *http.Client {
	t.Helper()
	if c == nil {
		t.Fatal("whitebit client is nil")
	}
	v := reflect.ValueOf(c).Elem()
	f := v.FieldByName("httpClient")
	if !f.IsValid() {
		t.Fatalf("field httpClient not found in %T", c)
	}
	f = reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
	cli, ok := f.Interface().(*http.Client)
	if !ok || cli == nil {
		t.Fatalf("httpClient is not *http.Client (got %T)", f.Interface())
	}
	return cli
}

func TestSDK_New_Defaults(t *testing.T) {
	s := wbSDK.New("key", "secret")
	if s == nil || s.Client == nil {
		t.Fatal("SDK or SDK.Client is nil")
	}
	if s.Client.BaseURL != "https://whitebit.com" {
		t.Fatalf("unexpected BaseURL: %s", s.Client.BaseURL)
	}
	hc := getHTTPClientPtr(t, s.Client)
	// default in sdk.New is cfg.Timeout=15s
	if hc.Timeout != 15*time.Second {
		t.Fatalf("unexpected http.Client timeout: %v", hc.Timeout)
	}
	if hc.Transport == nil {
		t.Fatalf("expected non-nil Transport on default client")
	}
}

func TestSDK_New_WithBaseURL(t *testing.T) {
	base := "https://example.test"
	s := wbSDK.New("k", "s", wbSDK.WithBaseURL(base))
	if s.Client.BaseURL != base {
		t.Fatalf("BaseURL not applied, got: %s", s.Client.BaseURL)
	}
}

func TestSDK_New_WithTimeout_NoHTTPClient(t *testing.T) {
	s := wbSDK.New("k", "s", wbSDK.WithTimeout(3*time.Second))
	hc := getHTTPClientPtr(t, s.Client)
	if hc.Timeout != 3*time.Second {
		t.Fatalf("timeout not applied, got: %v", hc.Timeout)
	}
}

func TestSDK_New_WithHTTPClient_Overrides(t *testing.T) {
	custom := &http.Client{Timeout: 1 * time.Second}
	s := wbSDK.New("k", "s", wbSDK.WithHTTPClient(custom), wbSDK.WithTimeout(10*time.Second))
	hc := getHTTPClientPtr(t, s.Client)
	if hc != custom {
		t.Fatalf("expected SDK to use provided http.Client instance")
	}
	if hc.Timeout != 1*time.Second {
		t.Fatalf("custom client timeout should be preserved, got: %v", hc.Timeout)
	}
}
