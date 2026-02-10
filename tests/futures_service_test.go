package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk/module/futures"
	testmocks "github.com/whitebit-exchange/go-sdk/tests/mocks"
)

func TestFutures_GetFuturesMarkets_Success(t *testing.T) {
	m := new(testmocks.Client)
	svc := futures.NewService(m)

	jsonResp := []byte(`{"success":true,"message":{},"result":[{"ticker_id":"BTC_USDT","stock_currency":"BTC","money_currency":"USDT"}]}`)
	m.On("SendRequest", mock.Anything).Return(jsonResp, nil).Once()

	res, err := svc.GetFuturesMarkets()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !res.Success || len(res.Result) == 0 || res.Result[0].TickerId != "BTC_USDT" {
		t.Fatalf("unexpected result: %#v", res)
	}
	m.AssertExpectations(t)
}

func TestFutures_GetFuturesMarkets_SendRequestError(t *testing.T) {
	m := new(testmocks.Client)
	svc := futures.NewService(m)
	m.On("SendRequest", mock.Anything).Return(nil, errors.New("network")).Once()

	_, err := svc.GetFuturesMarkets()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	m.AssertExpectations(t)
}

func TestFutures_GetFuturesMarkets_BadJSON(t *testing.T) {
	m := new(testmocks.Client)
	svc := futures.NewService(m)
	m.On("SendRequest", mock.Anything).Return([]byte(`{oops`), nil).Once()

	_, err := svc.GetFuturesMarkets()
	if err == nil {
		t.Fatalf("expected json unmarshal error, got nil")
	}
	m.AssertExpectations(t)
}
