package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk/module/deal"
	testmocks "github.com/whitebit-exchange/go-sdk/tests/mocks"
)

func TestDeals_GetDeals_Success(t *testing.T) {
	m := new(testmocks.Client)
	svc := deal.NewService(m)

	jsonResp := []byte(`[
        {"tradeID":1,"price":"40500","quote_volume":"1.0","base_volume":"1.0","trade_timestamp":1710000000,"type":"buy"}
    ]`)
	m.On("SendRequest", mock.Anything).Return(jsonResp, nil).Once()

	res, err := svc.GetDeals(deal.Options{Market: "BTC_USDT"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res) != 1 || res[0].TradeID != 1 {
		t.Fatalf("unexpected result: %#v", res)
	}
	m.AssertExpectations(t)
}

func TestDeals_GetDeals_SendRequestError(t *testing.T) {
	m := new(testmocks.Client)
	svc := deal.NewService(m)
	m.On("SendRequest", mock.Anything).Return(nil, errors.New("network")).Once()

	_, err := svc.GetDeals(deal.Options{Market: "BTC_USDT"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	m.AssertExpectations(t)
}

func TestDeals_GetDeals_BadJSON(t *testing.T) {
	m := new(testmocks.Client)
	svc := deal.NewService(m)
	m.On("SendRequest", mock.Anything).Return([]byte(`{oops`), nil).Once()

	_, err := svc.GetDeals(deal.Options{Market: "BTC_USDT"})
	if err == nil {
		t.Fatalf("expected json unmarshal error, got nil")
	}
	m.AssertExpectations(t)
}
