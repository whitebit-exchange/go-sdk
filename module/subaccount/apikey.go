package subaccount

import "github.com/whitebit-exchange/go-sdk"

const (
	apiKeyCreateURL    = "/api/v4/sub-account/api-key/create"
	apiKeyEditURL      = "/api/v4/sub-account/api-key/edit"
	apiKeyDeleteURL    = "/api/v4/sub-account/api-key/delete"
	apiKeyListURL      = "/api/v4/sub-account/api-key/list"
	apiKeyResetURL     = "/api/v4/sub-account/api-key/reset"
	ipAddressListURL   = "/api/v4/sub-account/api-key/ip-address/list"
	ipAddressCreateURL = "/api/v4/sub-account/api-key/ip-address/create"
	ipAddressDeleteURL = "/api/v4/sub-account/api-key/ip-address/delete"
)

type apiKeyCreateEndpoint struct {
	whitebit.AuthParams
	SubAccountID string   `json:"subAccountUid"`
	Label        string   `json:"label"`
	Permissions  []string `json:"permissions,omitempty"`
}

func newAPIKeyCreateEndpoint(subAccountID, label string, permissions []string) *apiKeyCreateEndpoint {
	return &apiKeyCreateEndpoint{
		AuthParams:   whitebit.NewAuthParams(apiKeyCreateURL),
		SubAccountID: subAccountID,
		Label:        label,
		Permissions:  permissions,
	}
}

type apiKeyEditEndpoint struct {
	whitebit.AuthParams
	SubAccountID string   `json:"subAccountUid"`
	APIKeyID     string   `json:"apiKeyId"`
	Label        string   `json:"label,omitempty"`
	Permissions  []string `json:"permissions,omitempty"`
}

func newAPIKeyEditEndpoint(subAccountID, apiKeyID, label string, permissions []string) *apiKeyEditEndpoint {
	return &apiKeyEditEndpoint{
		AuthParams:   whitebit.NewAuthParams(apiKeyEditURL),
		SubAccountID: subAccountID,
		APIKeyID:     apiKeyID,
		Label:        label,
		Permissions:  permissions,
	}
}

type apiKeyDeleteEndpoint struct {
	whitebit.AuthParams
	SubAccountID string `json:"subAccountUid"`
	APIKeyID     string `json:"apiKeyId"`
}

func newAPIKeyDeleteEndpoint(subAccountID, apiKeyID string) *apiKeyDeleteEndpoint {
	return &apiKeyDeleteEndpoint{
		AuthParams:   whitebit.NewAuthParams(apiKeyDeleteURL),
		SubAccountID: subAccountID,
		APIKeyID:     apiKeyID,
	}
}

type apiKeyListEndpoint struct {
	whitebit.AuthParams
	SubAccountID string `json:"subAccountUid"`
	Limit        int    `json:"limit,omitempty"`
	Offset       int    `json:"offset,omitempty"`
}

func newAPIKeyListEndpoint(subAccountID string, limit, offset int) *apiKeyListEndpoint {
	return &apiKeyListEndpoint{
		AuthParams:   whitebit.NewAuthParams(apiKeyListURL),
		SubAccountID: subAccountID,
		Limit:        limit,
		Offset:       offset,
	}
}

type apiKeyResetEndpoint struct {
	whitebit.AuthParams
	SubAccountID string `json:"subAccountUid"`
	APIKeyID     string `json:"apiKeyId"`
}

func newAPIKeyResetEndpoint(subAccountID, apiKeyID string) *apiKeyResetEndpoint {
	return &apiKeyResetEndpoint{
		AuthParams:   whitebit.NewAuthParams(apiKeyResetURL),
		SubAccountID: subAccountID,
		APIKeyID:     apiKeyID,
	}
}

type ipAddressListEndpoint struct {
	whitebit.AuthParams
	SubAccountID string `json:"subAccountUid"`
	APIKeyID     string `json:"apiKeyId"`
}

func newIPAddressListEndpoint(subAccountID, apiKeyID string) *ipAddressListEndpoint {
	return &ipAddressListEndpoint{
		AuthParams:   whitebit.NewAuthParams(ipAddressListURL),
		SubAccountID: subAccountID,
		APIKeyID:     apiKeyID,
	}
}

type ipAddressCreateEndpoint struct {
	whitebit.AuthParams
	SubAccountID string `json:"subAccountUid"`
	APIKeyID     string `json:"apiKeyId"`
	IPAddress    string `json:"ip"`
}

func newIPAddressCreateEndpoint(subAccountID, apiKeyID, ip string) *ipAddressCreateEndpoint {
	return &ipAddressCreateEndpoint{
		AuthParams:   whitebit.NewAuthParams(ipAddressCreateURL),
		SubAccountID: subAccountID,
		APIKeyID:     apiKeyID,
		IPAddress:    ip,
	}
}

type ipAddressDeleteEndpoint struct {
	whitebit.AuthParams
	SubAccountID string `json:"subAccountUid"`
	APIKeyID     string `json:"apiKeyId"`
	IPAddress    string `json:"ip"`
}

func newIPAddressDeleteEndpoint(subAccountID, apiKeyID, ip string) *ipAddressDeleteEndpoint {
	return &ipAddressDeleteEndpoint{
		AuthParams:   whitebit.NewAuthParams(ipAddressDeleteURL),
		SubAccountID: subAccountID,
		APIKeyID:     apiKeyID,
		IPAddress:    ip,
	}
}
