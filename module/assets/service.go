package assets

import (
	"encoding/json"
	"github.com/whitebit-exchange/whitebit"
)

type Asset struct {
	Name              string         `json:"name"`
	UCID              int            `json:"unified_cryptoasset_id"`
	CanWithdraw       bool           `json:"can_withdraw"`
	CanDeposit        bool           `json:"can_deposit"`
	MinWithdraw       string         `json:"min_withdraw"`
	MaxWithdraw       string         `json:"max_withdraw"`
	MakerFee          string         `json:"maker_fee"`
	TakerFee          string         `json:"taker_fee"`
	MinDeposit        string         `json:"min_deposit"`
	MaxDeposit        string         `json:"max_deposit"`
	Networks          AssetNetworks  `json:"networks,omitempty"`
	Confirmations     map[string]int `json:"confirmations,omitempty"`
	Providers         AssetProviders `json:"providers,omitempty"`
	Limits            AssetLimits    `json:"limits,omitempty"`
	CurrencyPrecision int            `json:"currency_precision"`
	IsMemo            bool           `json:"is_memo"`
}

type AssetNetworks struct {
	AssetProviders
	Default string `json:"default,omitempty"`
}

type AssetProviders struct {
	Deposits  []string `json:"deposits,omitempty"`
	Withdraws []string `json:"withdraws,omitempty"`
}

type AssetLimits struct {
	Deposit     map[string]AssetLimitItem `json:"deposit,omitempty"`
	Withdraw    map[string]AssetLimitItem `json:"withdraw,omitempty"`
	MaxDeposit  float64                   `json:"-"`
	MinDeposit  float64                   `json:"-"`
	MaxWithdraw float64                   `json:"-"`
	MinWithdraw float64                   `json:"-"`
}

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

func (service *Service) GetAssets() (map[string]Asset, error) {
	response, err := service.client.SendRequest(newAssetsEndpoint())
	if err != nil {
		return nil, err
	}

	var result map[string]Asset
	err = json.Unmarshal(response, &result)

	if err != nil {
		return nil, err
	}

	return result, nil

}
