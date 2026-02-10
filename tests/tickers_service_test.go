package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk/module/tickers"
	testmocks "github.com/whitebit-exchange/go-sdk/tests/mocks"
)

func TestTickers_GetTickers_Success(t *testing.T) {
	m := new(testmocks.Client)
	svc := tickers.NewService(m)

	jsonResp := []byte(`{
        "success": true,
        "message": {},
        "result": {
            "BTC_USDT": {
                "ticker": {
                    "open": "40000",
                    "bid": "40500",
                    "ask": "40600",
                    "low": "39500",
                    "high": "41000",
                    "last": "40550",
                    "volume": "123.45",
                    "deal": "5000000",
                    "change": "+1.2%"
                },
                "at": 1710000000
            }
        }
    }`)
	m.On("SendRequest", mock.Anything).Return(jsonResp, nil).Once()

	res, err := svc.GetTickers()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !res.Success {
		t.Fatalf("expected success=true")
	}
	if _, ok := res.Result["BTC_USDT"]; !ok {
		t.Fatalf("expected BTC_USDT ticker in result")
	}
	m.AssertExpectations(t)
}

func TestTickers_GetTickers_SendRequestError(t *testing.T) {
	m := new(testmocks.Client)
	svc := tickers.NewService(m)
	m.On("SendRequest", mock.Anything).Return(nil, errors.New("network")).Once()

	_, err := svc.GetTickers()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	m.AssertExpectations(t)
}

func TestTickers_GetTickers_BadJSON(t *testing.T) {
	m := new(testmocks.Client)
	svc := tickers.NewService(m)
	m.On("SendRequest", mock.Anything).Return([]byte(`{oops`), nil).Once()

	_, err := svc.GetTickers()
	if err == nil {
		t.Fatalf("expected json unmarshal error, got nil")
	}
	m.AssertExpectations(t)
}
