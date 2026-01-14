package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk/module/fee"
	testmocks "github.com/whitebit-exchange/go-sdk/tests/mocks"
)

func TestFee_GetTradingFee_Success(t *testing.T) {
	m := new(testmocks.Client)
	svc := fee.NewService(m)

	jsonResp := []byte(`{"success":true,"message":{},"result":{"makerFee":"0.1","takerFee":"0.1"}}`)
	m.On("SendRequest", mock.Anything).Return(jsonResp, nil).Once()

	res, err := svc.GetTradingFee()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !res.Success || res.Result.MakerFee == "" {
		t.Fatalf("unexpected result: %#v", res)
	}
	m.AssertExpectations(t)
}

func TestFee_GetTradingFee_SendRequestError(t *testing.T) {
	m := new(testmocks.Client)
	svc := fee.NewService(m)
	m.On("SendRequest", mock.Anything).Return(nil, errors.New("network")).Once()

	_, err := svc.GetTradingFee()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	m.AssertExpectations(t)
}

func TestFee_GetTradingFee_BadJSON(t *testing.T) {
	m := new(testmocks.Client)
	svc := fee.NewService(m)
	m.On("SendRequest", mock.Anything).Return([]byte(`{oops`), nil).Once()

	_, err := svc.GetTradingFee()
	if err == nil {
		t.Fatalf("expected json unmarshal error, got nil")
	}
	m.AssertExpectations(t)
}
