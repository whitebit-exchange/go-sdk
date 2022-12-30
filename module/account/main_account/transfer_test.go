package main_account

import (
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *MainAccountTestSuite) TestTransferWithResult() {
	endpoint := newBalanceEndpoint("")
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/main-account/balance"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	err := s.service.Transfer(TransferParams{
		Ticker: "BTC",
		Amount: "0.001",
		From:   Main,
		To:     Collateral,
	})

	s.Equal(err, error(nil))

}
