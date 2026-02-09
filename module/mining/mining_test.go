package mining

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/tests/mocks"
)

type MiningTestSuite struct {
	client  *mocks.Client
	service *Service
	suite.Suite
}

func (s *MiningTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
}

func (s *MiningTestSuite) TestGetRewards() {
	endpoint := newRewardsEndpoint(10, 0)

	byteResponse := []byte(`{
		"total": 2,
		"records": [
			{"id": 1, "ticker": "BTC", "amount": "0.00001", "timestamp": 1700000000},
			{"id": 2, "ticker": "BTC", "amount": "0.00002", "timestamp": 1700086400}
		],
		"limit": 10,
		"offset": 0
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/mining/rewards", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.GetRewards(10, 0)

	s.NoError(err)
	s.Equal(2, result.Total)
	s.Equal(2, len(result.Records))
	s.Equal(int64(1), result.Records[0].ID)
	s.Equal("BTC", result.Records[0].Ticker)
	s.Equal("0.00001", result.Records[0].Amount)
}

func (s *MiningTestSuite) TestGetHashrate() {
	endpoint := newHashrateEndpoint(10, 0)

	byteResponse := []byte(`{
		"total": 2,
		"records": [
			{"hashrate": "100.5", "timestamp": 1700000000},
			{"hashrate": "102.3", "timestamp": 1700086400}
		],
		"limit": 10,
		"offset": 0
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/mining/hashrate", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.GetHashrate(10, 0)

	s.NoError(err)
	s.Equal(2, result.Total)
	s.Equal(2, len(result.Records))
	s.Equal("100.5", result.Records[0].Hashrate)
	s.Equal(int64(1700000000), result.Records[0].Timestamp)
}

func (s *MiningTestSuite) TestGetRewardsInvalidResponse() {
	byteResponse := []byte(`invalid json`)

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	_, err := s.service.GetRewards(10, 0)

	s.Error(err)
	s.Contains(err.Error(), "invalid character")
}

func (s *MiningTestSuite) TestGetHashrateInvalidResponse() {
	byteResponse := []byte(`invalid json`)

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	_, err := s.service.GetHashrate(10, 0)

	s.Error(err)
	s.Contains(err.Error(), "invalid character")
}

func TestMiningTestSuite(t *testing.T) {
	suite.Run(t, new(MiningTestSuite))
}
