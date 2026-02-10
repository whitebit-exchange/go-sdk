package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk/module/kline"
	testmocks "github.com/whitebit-exchange/go-sdk/tests/mocks"
)

func TestKline_GetKline_Success(t *testing.T) {
	m := new(testmocks.Client)
	svc := kline.NewService(m)

	jsonResp := []byte(`{"success":true,"message":{},"result":[["1710000000","40500","40600","40000","41000","100.5"]]}`)
	m.On("SendRequest", mock.Anything).Return(jsonResp, nil).Once()

	res, err := svc.GetKline(kline.Options{Market: "BTC_USDT", Interval: "1m"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !res.Success || len(res.Result) == 0 {
		t.Fatalf("unexpected result: %#v", res)
	}
	m.AssertExpectations(t)
}

func TestKline_GetKline_SendRequestError(t *testing.T) {
	m := new(testmocks.Client)
	svc := kline.NewService(m)
	m.On("SendRequest", mock.Anything).Return(nil, errors.New("network")).Once()

	_, err := svc.GetKline(kline.Options{Market: "BTC_USDT"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	m.AssertExpectations(t)
}

func TestKline_GetKline_BadJSON(t *testing.T) {
	m := new(testmocks.Client)
	svc := kline.NewService(m)
	m.On("SendRequest", mock.Anything).Return([]byte(`{oops`), nil).Once()

	_, err := svc.GetKline(kline.Options{Market: "BTC_USDT"})
	if err == nil {
		t.Fatalf("expected json unmarshal error, got nil")
	}
	m.AssertExpectations(t)
}
