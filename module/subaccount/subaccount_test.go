package subaccount

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/tests/mocks"
)

type SubAccountTestSuite struct {
	client  *mocks.Client
	service *Service
	suite.Suite
}

func (s *SubAccountTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
}

func (s *SubAccountTestSuite) TestCreate() {
	endpoint := newCreateEndpoint("test-alias", "test@example.com")

	byteResponse := []byte(`{
		"id": "sub-123",
		"alias": "test-alias",
		"email": "test@example.com",
		"status": "active",
		"createdAt": 1700000000
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/sub-account/create", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.Create("test-alias", "test@example.com")

	s.NoError(err)
	s.Equal("sub-123", result.ID)
	s.Equal("test-alias", result.Alias)
	s.Equal("test@example.com", result.Email)
}

func (s *SubAccountTestSuite) TestDelete() {
	endpoint := newDeleteEndpoint("sub-123")

	byteResponse := []byte(`{}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/sub-account/delete", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	err := s.service.Delete("sub-123")

	s.NoError(err)
}

func (s *SubAccountTestSuite) TestEdit() {
	endpoint := newEditEndpoint("sub-123", "new-alias")

	byteResponse := []byte(`{
		"id": "sub-123",
		"alias": "new-alias",
		"email": "test@example.com",
		"status": "active",
		"createdAt": 1700000000
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/sub-account/edit", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.Edit("sub-123", "new-alias")

	s.NoError(err)
	s.Equal("sub-123", result.ID)
	s.Equal("new-alias", result.Alias)
}

func (s *SubAccountTestSuite) TestList() {
	endpoint := newListEndpoint(10, 0)

	byteResponse := []byte(`{
		"total": 2,
		"records": [
			{"id": "sub-1", "alias": "alias1", "email": "a@b.com", "status": "active", "createdAt": 1700000000},
			{"id": "sub-2", "alias": "alias2", "email": "c@d.com", "status": "active", "createdAt": 1700000001}
		],
		"limit": 10,
		"offset": 0
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/sub-account/list", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.List(10, 0)

	s.NoError(err)
	s.Equal(2, result.Total)
	s.Equal(2, len(result.Records))
	s.Equal("sub-1", result.Records[0].ID)
}

func (s *SubAccountTestSuite) TestTransfer() {
	endpoint := newTransferEndpoint("sub-1", "sub-2", "USDT", "100")

	byteResponse := []byte(`{}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/sub-account/transfer", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	err := s.service.Transfer("sub-1", "sub-2", "USDT", "100")

	s.NoError(err)
}

func (s *SubAccountTestSuite) TestBlock() {
	endpoint := newBlockEndpoint("sub-123")

	byteResponse := []byte(`{}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/sub-account/block", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	err := s.service.Block("sub-123")

	s.NoError(err)
}

func (s *SubAccountTestSuite) TestUnblock() {
	endpoint := newUnblockEndpoint("sub-123")

	byteResponse := []byte(`{}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/sub-account/unblock", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	err := s.service.Unblock("sub-123")

	s.NoError(err)
}

func (s *SubAccountTestSuite) TestGetBalances() {
	endpoint := newBalancesEndpoint("sub-123", "USDT")

	byteResponse := []byte(`{
		"USDT": {"main_balance": "1000.50"}
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/sub-account/balances", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.GetBalances("sub-123", "USDT")

	s.NoError(err)
	s.Equal("1000.50", result["USDT"].MainBalance)
}

func (s *SubAccountTestSuite) TestGetTransferHistory() {
	endpoint := newTransferHistoryEndpoint("sub-123", 10, 0)

	byteResponse := []byte(`{
		"total": 1,
		"records": [
			{"id": "tx-1", "from": "sub-1", "to": "sub-2", "ticker": "USDT", "amount": "100", "createdAt": 1700000000}
		],
		"limit": 10,
		"offset": 0
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/sub-account/transfer/history", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.GetTransferHistory("sub-123", 10, 0)

	s.NoError(err)
	s.Equal(1, result.Total)
	s.Equal("tx-1", result.Records[0].ID)
}

func (s *SubAccountTestSuite) TestCreateAPIKey() {
	endpoint := newAPIKeyCreateEndpoint("sub-123", "test-key", []string{"read", "trade"})

	byteResponse := []byte(`{
		"id": "key-1",
		"apiKey": "abc123",
		"secretKey": "secret123",
		"label": "test-key",
		"permissions": ["read", "trade"],
		"createdAt": 1700000000
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/sub-account/api-key/create", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.CreateAPIKey("sub-123", "test-key", []string{"read", "trade"})

	s.NoError(err)
	s.Equal("key-1", result.ID)
	s.Equal("abc123", result.APIKey)
	s.Equal("secret123", result.SecretKey)
}

func (s *SubAccountTestSuite) TestListAPIKeys() {
	endpoint := newAPIKeyListEndpoint("sub-123", 10, 0)

	byteResponse := []byte(`{
		"total": 1,
		"records": [
			{"id": "key-1", "apiKey": "abc123", "label": "test-key", "permissions": ["read"], "createdAt": 1700000000}
		],
		"limit": 10,
		"offset": 0
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/sub-account/api-key/list", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.ListAPIKeys("sub-123", 10, 0)

	s.NoError(err)
	s.Equal(1, result.Total)
	s.Equal("key-1", result.Records[0].ID)
}

func (s *SubAccountTestSuite) TestDeleteAPIKey() {
	endpoint := newAPIKeyDeleteEndpoint("sub-123", "key-1")

	byteResponse := []byte(`{}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/sub-account/api-key/delete", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	err := s.service.DeleteAPIKey("sub-123", "key-1")

	s.NoError(err)
}

func (s *SubAccountTestSuite) TestListIPAddresses() {
	endpoint := newIPAddressListEndpoint("sub-123", "key-1")

	byteResponse := []byte(`[
		{"ip": "192.168.1.1", "createdAt": 1700000000}
	]`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/sub-account/api-key/ip-address/list", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.ListIPAddresses("sub-123", "key-1")

	s.NoError(err)
	s.Equal(1, len(result))
	s.Equal("192.168.1.1", result[0].IP)
}

func (s *SubAccountTestSuite) TestInvalidResponse() {
	byteResponse := []byte(`invalid json`)

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	_, err := s.service.Create("test", "test@test.com")

	s.Error(err)
	s.Contains(err.Error(), "invalid character")
}

func TestSubAccountTestSuite(t *testing.T) {
	suite.Run(t, new(SubAccountTestSuite))
}
