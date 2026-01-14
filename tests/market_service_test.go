package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk/module/market"
	testmocks "github.com/whitebit-exchange/go-sdk/tests/mocks"
)

func TestMarket_GetMarkets_Success(t *testing.T) {
	m := new(testmocks.Client)
	svc := market.NewService(m)

	jsonResp := []byte(`{
        "success": true,
        "message": {},
        "result": [
            {"name":"BTC_USDT","stock":"BTC","money":"USDT","tradesEnabled":true}
        ]
    }`)
	m.On("SendRequest", mock.Anything).Return(jsonResp, nil).Once()

	res, err := svc.GetMarkets()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !res.Success {
		t.Fatalf("expected success=true")
	}
	if len(res.Result) != 1 || res.Result[0].Name != "BTC_USDT" {
		t.Fatalf("unexpected result: %#v", res.Result)
	}
	m.AssertExpectations(t)
}

func TestMarket_GetMarkets_SendRequestError(t *testing.T) {
	m := new(testmocks.Client)
	svc := market.NewService(m)
	m.On("SendRequest", mock.Anything).Return(nil, errors.New("network")).Once()

	_, err := svc.GetMarkets()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	m.AssertExpectations(t)
}

func TestMarket_GetMarkets_BadJSON(t *testing.T) {
	m := new(testmocks.Client)
	svc := market.NewService(m)
	m.On("SendRequest", mock.Anything).Return([]byte(`{oops`), nil).Once()

	_, err := svc.GetMarkets()
	if err == nil {
		t.Fatalf("expected json unmarshal error, got nil")
	}
	m.AssertExpectations(t)
}
