package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk/module/assets"
	testmocks "github.com/whitebit-exchange/go-sdk/tests/mocks"
)

func TestAssets_GetAssets_Success(t *testing.T) {
	m := new(testmocks.Client)
	svc := assets.NewService(m)

	jsonResp := []byte(`{"BTC":{"name":"Bitcoin","unified_cryptoasset_id":1}}`)
	m.On("SendRequest", mock.Anything).Return(jsonResp, nil).Once()

	res, err := svc.GetAssets()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := res["BTC"]; !ok {
		t.Fatalf("expected BTC in assets result")
	}
	m.AssertExpectations(t)
}

func TestAssets_GetAssets_SendRequestError(t *testing.T) {
	m := new(testmocks.Client)
	svc := assets.NewService(m)
	m.On("SendRequest", mock.Anything).Return(nil, errors.New("network")).Once()

	_, err := svc.GetAssets()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	m.AssertExpectations(t)
}

func TestAssets_GetAssets_BadJSON(t *testing.T) {
	m := new(testmocks.Client)
	svc := assets.NewService(m)
	m.On("SendRequest", mock.Anything).Return([]byte(`{oops`), nil).Once()

	_, err := svc.GetAssets()
	if err == nil {
		t.Fatalf("expected json unmarshal error, got nil")
	}
	m.AssertExpectations(t)
}
