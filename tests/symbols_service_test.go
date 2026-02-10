package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk/module/symbol"
	testmocks "github.com/whitebit-exchange/go-sdk/tests/mocks"
)

func TestSymbols_GetSymbols_Success(t *testing.T) {
	// Arrange
	m := new(testmocks.Client)
	svc := symbol.NewService(m)

	jsonResp := []byte(`{"success":true,"message":{},"result":["BTC_USDT","ETH_USDT"]}`)
	m.On("SendRequest", mock.Anything).Return(jsonResp, nil).Once()

	// Act
	res, err := svc.GetSymbols()

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !res.Success {
		t.Fatalf("expected success=true")
	}
	if len(res.Result) != 2 || res.Result[0] != "BTC_USDT" || res.Result[1] != "ETH_USDT" {
		t.Fatalf("unexpected result: %#v", res.Result)
	}
	m.AssertExpectations(t)
}

func TestSymbols_GetSymbols_SendRequestError(t *testing.T) {
	m := new(testmocks.Client)
	svc := symbol.NewService(m)

	m.On("SendRequest", mock.Anything).Return(nil, errors.New("network"))

	_, err := svc.GetSymbols()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	m.AssertExpectations(t)
}

func TestSymbols_GetSymbols_BadJSON(t *testing.T) {
	m := new(testmocks.Client)
	svc := symbol.NewService(m)

	m.On("SendRequest", mock.Anything).Return([]byte(`{not-json`), nil)

	_, err := svc.GetSymbols()
	if err == nil {
		t.Fatalf("expected json unmarshal error, got nil")
	}
	m.AssertExpectations(t)
}
