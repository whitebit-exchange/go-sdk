package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk/module/depth"
	testmocks "github.com/whitebit-exchange/go-sdk/tests/mocks"
)

func TestDepth_GetDepth_Success(t *testing.T) {
	m := new(testmocks.Client)
	svc := depth.NewService(m)

	jsonResp := []byte(`{
        "success": true,
        "message": {},
        "result": {
            "lastUpdateTimestamp": "1710000000",
            "asks": [["40600","1.2"]],
            "bids": [["40500","0.8"]]
        }
    }`)
	m.On("SendRequest", mock.Anything).Return(jsonResp, nil).Once()

	res, err := svc.GetDepth("BTC_USDT")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !res.Success {
		t.Fatalf("expected success=true")
	}
	if len(res.Result.Asks) == 0 || len(res.Result.Bids) == 0 {
		t.Fatalf("expected non-empty asks/bids")
	}
	m.AssertExpectations(t)
}

func TestDepth_GetDepth_SendRequestError(t *testing.T) {
	m := new(testmocks.Client)
	svc := depth.NewService(m)
	m.On("SendRequest", mock.Anything).Return(nil, errors.New("network")).Once()

	_, err := svc.GetDepth("BTC_USDT")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	m.AssertExpectations(t)
}

func TestDepth_GetDepth_BadJSON(t *testing.T) {
	m := new(testmocks.Client)
	svc := depth.NewService(m)
	m.On("SendRequest", mock.Anything).Return([]byte(`{oops`), nil).Once()

	_, err := svc.GetDepth("BTC_USDT")
	if err == nil {
		t.Fatalf("expected json unmarshal error, got nil")
	}
	m.AssertExpectations(t)
}
