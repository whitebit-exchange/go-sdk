package tests

import (
	"reflect"
	"testing"
	"unsafe"

	whitebit "github.com/whitebit-exchange/go-sdk"
	wbSDK "github.com/whitebit-exchange/go-sdk/sdk"
)

// getServiceClientPtr uses reflection+unsafe to extract the unexported `client` field
// from any service struct pointer and returns the underlying *whitebit.Whitebit pointer.
func getServiceClientPtr(t *testing.T, svc any) *whitebit.Whitebit {
	t.Helper()
	v := reflect.ValueOf(svc)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		t.Fatalf("service must be non-nil pointer, got %T", svc)
	}
	elem := v.Elem()
	if elem.Kind() != reflect.Struct {
		t.Fatalf("service must point to struct, got %v", elem.Kind())
	}
	f := elem.FieldByName("client")
	if !f.IsValid() {
		t.Fatalf("service has no field 'client' (type: %T)", svc)
	}
	// Make unexported field addressable and readable
	f = reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
	cli, ok := f.Interface().(*whitebit.Whitebit)
	if !ok || cli == nil {
		t.Fatalf("client field is not *whitebit.Whitebit (got %T)", f.Interface())
	}
	return cli
}

func TestSDK_AllServicesShareSameClientPointer(t *testing.T) {
	s := wbSDK.New("key", "secret")

	// Basic non-nil checks
	if s.Client == nil {
		t.Fatal("SDK Client is nil")
	}
	nonNilServices := []struct {
		name string
		v    any
	}{
		{"Server", s.Server},
		{"Market", s.Market},
		{"Depth", s.Depth},
		{"Tickers", s.Tickers},
		{"Symbols", s.Symbols},
		{"Deals", s.Deals},
		{"Assets", s.Assets},
		{"Fee", s.Fee},
		{"Futures", s.Futures},
		{"Kline", s.Kline},
		{"OrdersSpot", s.OrdersSpot},
		{"OrdersCollateral", s.OrdersCollateral},
		{"AccountTrade", s.AccountTrade},
		{"AccountCollateral", s.AccountCollateral},
		{"AccountMain", s.AccountMain},
	}
	for _, it := range nonNilServices {
		if it.v == nil {
			t.Fatalf("service %s is nil", it.name)
		}
	}

	// All services must hold the exact same *whitebit.Whitebit pointer as SDK.Client
	for _, it := range nonNilServices {
		cli := getServiceClientPtr(t, it.v)
		if cli != s.Client {
			t.Fatalf("service %s has different client pointer: %p != %p", it.name, cli, s.Client)
		}
	}
}
