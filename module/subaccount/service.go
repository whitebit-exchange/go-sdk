package subaccount

import (
	"encoding/json"

	"github.com/whitebit-exchange/go-sdk"
)

type SubAccount struct {
	ID        string `json:"id"`
	Alias     string `json:"alias"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt int64  `json:"createdAt"`
}

type ListResponse struct {
	Total   int          `json:"total"`
	Records []SubAccount `json:"records"`
	Limit   int          `json:"limit"`
	Offset  int          `json:"offset"`
}

type Balance struct {
	MainBalance string `json:"main_balance"`
}

type BalancesMap map[string]Balance

type TransferRecord struct {
	ID        string `json:"id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Ticker    string `json:"ticker"`
	Amount    string `json:"amount"`
	CreatedAt int64  `json:"createdAt"`
}

type TransferHistoryResponse struct {
	Total   int              `json:"total"`
	Records []TransferRecord `json:"records"`
	Limit   int              `json:"limit"`
	Offset  int              `json:"offset"`
}

type APIKey struct {
	ID          string   `json:"id"`
	APIKey      string   `json:"apiKey"`
	SecretKey   string   `json:"secretKey,omitempty"`
	Label       string   `json:"label"`
	Permissions []string `json:"permissions"`
	CreatedAt   int64    `json:"createdAt"`
}

type APIKeyListResponse struct {
	Total   int      `json:"total"`
	Records []APIKey `json:"records"`
	Limit   int      `json:"limit"`
	Offset  int      `json:"offset"`
}

type IPAddress struct {
	IP        string `json:"ip"`
	CreatedAt int64  `json:"createdAt"`
}

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

// Create creates a new sub-account.
func (s *Service) Create(alias, email string) (SubAccount, error) {
	resp, err := s.client.SendRequest(newCreateEndpoint(alias, email))
	if err != nil {
		return SubAccount{}, err
	}

	var result SubAccount
	if err := json.Unmarshal(resp, &result); err != nil {
		return SubAccount{}, err
	}
	return result, nil
}

// Delete deletes a sub-account by ID.
func (s *Service) Delete(id string) error {
	_, err := s.client.SendRequest(newDeleteEndpoint(id))
	return err
}

// Edit updates a sub-account alias.
func (s *Service) Edit(id, alias string) (SubAccount, error) {
	resp, err := s.client.SendRequest(newEditEndpoint(id, alias))
	if err != nil {
		return SubAccount{}, err
	}

	var result SubAccount
	if err := json.Unmarshal(resp, &result); err != nil {
		return SubAccount{}, err
	}
	return result, nil
}

// List returns a list of sub-accounts.
func (s *Service) List(limit, offset int) (ListResponse, error) {
	resp, err := s.client.SendRequest(newListEndpoint(limit, offset))
	if err != nil {
		return ListResponse{}, err
	}

	var result ListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return ListResponse{}, err
	}
	return result, nil
}

// Transfer transfers funds between accounts.
// Use empty string for fromSubAccountID or toSubAccountID to transfer from/to main account.
func (s *Service) Transfer(fromSubAccountID, toSubAccountID, ticker, amount string) error {
	_, err := s.client.SendRequest(newTransferEndpoint(fromSubAccountID, toSubAccountID, ticker, amount))
	return err
}

// Block blocks a sub-account.
func (s *Service) Block(id string) error {
	_, err := s.client.SendRequest(newBlockEndpoint(id))
	return err
}

// Unblock unblocks a sub-account.
func (s *Service) Unblock(id string) error {
	_, err := s.client.SendRequest(newUnblockEndpoint(id))
	return err
}

// GetBalances returns balances for a sub-account.
func (s *Service) GetBalances(id, ticker string) (BalancesMap, error) {
	resp, err := s.client.SendRequest(newBalancesEndpoint(id, ticker))
	if err != nil {
		return nil, err
	}

	result := make(BalancesMap)
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetTransferHistory returns transfer history.
func (s *Service) GetTransferHistory(subAccountID string, limit, offset int) (TransferHistoryResponse, error) {
	resp, err := s.client.SendRequest(newTransferHistoryEndpoint(subAccountID, limit, offset))
	if err != nil {
		return TransferHistoryResponse{}, err
	}

	var result TransferHistoryResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return TransferHistoryResponse{}, err
	}
	return result, nil
}

// CreateAPIKey creates a new API key for a sub-account.
func (s *Service) CreateAPIKey(subAccountID, label string, permissions []string) (APIKey, error) {
	resp, err := s.client.SendRequest(newAPIKeyCreateEndpoint(subAccountID, label, permissions))
	if err != nil {
		return APIKey{}, err
	}

	var result APIKey
	if err := json.Unmarshal(resp, &result); err != nil {
		return APIKey{}, err
	}
	return result, nil
}

// EditAPIKey updates an API key.
func (s *Service) EditAPIKey(subAccountID, apiKeyID, label string, permissions []string) (APIKey, error) {
	resp, err := s.client.SendRequest(newAPIKeyEditEndpoint(subAccountID, apiKeyID, label, permissions))
	if err != nil {
		return APIKey{}, err
	}

	var result APIKey
	if err := json.Unmarshal(resp, &result); err != nil {
		return APIKey{}, err
	}
	return result, nil
}

// DeleteAPIKey deletes an API key.
func (s *Service) DeleteAPIKey(subAccountID, apiKeyID string) error {
	_, err := s.client.SendRequest(newAPIKeyDeleteEndpoint(subAccountID, apiKeyID))
	return err
}

// ListAPIKeys returns API keys for a sub-account.
func (s *Service) ListAPIKeys(subAccountID string, limit, offset int) (APIKeyListResponse, error) {
	resp, err := s.client.SendRequest(newAPIKeyListEndpoint(subAccountID, limit, offset))
	if err != nil {
		return APIKeyListResponse{}, err
	}

	var result APIKeyListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return APIKeyListResponse{}, err
	}
	return result, nil
}

// ResetAPIKeySecret resets the secret key for an API key.
func (s *Service) ResetAPIKeySecret(subAccountID, apiKeyID string) (APIKey, error) {
	resp, err := s.client.SendRequest(newAPIKeyResetEndpoint(subAccountID, apiKeyID))
	if err != nil {
		return APIKey{}, err
	}

	var result APIKey
	if err := json.Unmarshal(resp, &result); err != nil {
		return APIKey{}, err
	}
	return result, nil
}

// ListIPAddresses returns IP addresses for an API key.
func (s *Service) ListIPAddresses(subAccountID, apiKeyID string) ([]IPAddress, error) {
	resp, err := s.client.SendRequest(newIPAddressListEndpoint(subAccountID, apiKeyID))
	if err != nil {
		return nil, err
	}

	var result []IPAddress
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// AddIPAddress adds an IP address to an API key whitelist.
func (s *Service) AddIPAddress(subAccountID, apiKeyID, ip string) error {
	_, err := s.client.SendRequest(newIPAddressCreateEndpoint(subAccountID, apiKeyID, ip))
	return err
}

// RemoveIPAddress removes an IP address from an API key whitelist.
func (s *Service) RemoveIPAddress(subAccountID, apiKeyID, ip string) error {
	_, err := s.client.SendRequest(newIPAddressDeleteEndpoint(subAccountID, apiKeyID, ip))
	return err
}
