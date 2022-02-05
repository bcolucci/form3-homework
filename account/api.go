package account

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	cus_http "github.com/bcolucci/form3-homework/http"
	"github.com/google/uuid"
)

// AccountApi defines the Organisation>Account API (partial) interface
type AccountApi interface {

	// CreateAccount creates an account for the given organisation
	CreateAccount(ctx context.Context, organizationID string, create AccountCreate) (*AccountData, error)

	// FetchAccount fetches a given account
	FetchAccount(ctx context.Context, accountID string) (*AccountData, error)

	// DeleteAccount deletes a given account (and its version)
	DeleteAccount(ctx context.Context, accountID string, version int64) error
}

// this is what has to be provided to create an account
type accountCreateData struct {
	Type           string        `json:"type"`
	ID             string        `json:"id"`
	OrganizationID string        `json:"organisation_id"`
	Attributes     AccountCreate `json:"attributes"`
}

// basic AccountApi implementation
type accountApiSimpleImpl struct {
	host   string
	client cus_http.HttpClient
}

// NewAccountAPI creates an AccountApi instance
func NewAccountAPI(host string, client cus_http.HttpClient) AccountApi {
	return &accountApiSimpleImpl{host, client}
}

func (api *accountApiSimpleImpl) CreateAccount(ctx context.Context, organizationID string, create AccountCreate) (*AccountData, error) {

	b, err := json.Marshal(struct{ Data accountCreateData }{
		accountCreateData{
			Type:           "accounts",
			ID:             uuid.NewString(),
			OrganizationID: organizationID,
			Attributes:     create,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error serializing: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, api.host+"/organisation/accounts", bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := api.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unpextected status %d", res.StatusCode)
	}
	defer res.Body.Close()

	var d struct{ Data AccountData }
	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
		return nil, fmt.Errorf("error unserializing %w", err)
	}

	return &d.Data, nil
}

func (api *accountApiSimpleImpl) FetchAccount(ctx context.Context, accountID string) (*AccountData, error) {

	req, err := http.NewRequest(http.MethodGet, api.host+"/organisation/accounts/"+accountID, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := api.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unpextected status %d", res.StatusCode)
	}
	defer res.Body.Close()

	var d struct{ Data AccountData }
	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
		return nil, fmt.Errorf("error unserializing %w", err)
	}

	return &d.Data, nil
}

func (api *accountApiSimpleImpl) DeleteAccount(ctx context.Context, accountID string, version int64) error {

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/organisation/accounts/%s?version=%d", api.host, accountID, version), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	res, err := api.client.Do(req.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("error executing request: %w", err)
	}

	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unpextected status %d", res.StatusCode)
	}

	return nil
}
