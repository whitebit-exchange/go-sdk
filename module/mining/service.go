package mining

import (
	"encoding/json"

	"github.com/whitebit-exchange/go-sdk"
)

const (
	rewardsURL  = "/api/v4/mining/rewards"
	hashrateURL = "/api/v4/mining/hashrate"
)

type rewardsEndpoint struct {
	whitebit.AuthParams
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

func newRewardsEndpoint(limit, offset int) *rewardsEndpoint {
	return &rewardsEndpoint{
		AuthParams: whitebit.NewAuthParams(rewardsURL),
		Limit:      limit,
		Offset:     offset,
	}
}

type hashrateEndpoint struct {
	whitebit.AuthParams
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

func newHashrateEndpoint(limit, offset int) *hashrateEndpoint {
	return &hashrateEndpoint{
		AuthParams: whitebit.NewAuthParams(hashrateURL),
		Limit:      limit,
		Offset:     offset,
	}
}

type Reward struct {
	ID        int64  `json:"id"`
	Ticker    string `json:"ticker"`
	Amount    string `json:"amount"`
	Timestamp int64  `json:"timestamp"`
}

type RewardsResponse struct {
	Total   int      `json:"total"`
	Records []Reward `json:"records"`
	Limit   int      `json:"limit"`
	Offset  int      `json:"offset"`
}

type HashrateRecord struct {
	Hashrate  string `json:"hashrate"`
	Timestamp int64  `json:"timestamp"`
}

type HashrateResponse struct {
	Total   int              `json:"total"`
	Records []HashrateRecord `json:"records"`
	Limit   int              `json:"limit"`
	Offset  int              `json:"offset"`
}

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

// GetRewards returns mining rewards history.
func (s *Service) GetRewards(limit, offset int) (RewardsResponse, error) {
	resp, err := s.client.SendRequest(newRewardsEndpoint(limit, offset))
	if err != nil {
		return RewardsResponse{}, err
	}

	var result RewardsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return RewardsResponse{}, err
	}
	return result, nil
}

// GetHashrate returns hashrate history.
func (s *Service) GetHashrate(limit, offset int) (HashrateResponse, error) {
	resp, err := s.client.SendRequest(newHashrateEndpoint(limit, offset))
	if err != nil {
		return HashrateResponse{}, err
	}

	var result HashrateResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return HashrateResponse{}, err
	}
	return result, nil
}
