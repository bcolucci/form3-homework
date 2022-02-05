package integration_test

import (
	"context"
	"testing"

	"github.com/bcolucci/form3-homework/account"
	cus_http "github.com/bcolucci/form3-homework/http"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var orgID = uuid.NewString()

var api = account.NewAccountAPI("http://accountapi:8080/v1", &cus_http.TimeoutHttpClient{})

var briceColucciAccount = account.AccountCreate{
	Name:    []string{"Brice Colucci"},
	Country: "CA",
	Bic:     "NWBKGB22",
}

func TestCreate(t *testing.T) {
	account := createAccount(t)
	fetchAccount(t, account.ID)
	deleteAccount(t, account)
}

func createAccount(t *testing.T) *account.AccountData {
	t.Log("creating account ...")

	account, err := api.CreateAccount(context.Background(), orgID, briceColucciAccount)
	assert.Nil(t, err)
	assert.NotNil(t, account)
	t.Logf("account %s created", account.ID)

	return account
}

func fetchAccount(t *testing.T, accountID string) {
	t.Logf("fetching account %s ...", accountID)

	account, err := api.FetchAccount(context.Background(), accountID)
	assert.Nil(t, err)
	assert.NotNil(t, account)

	assert.Equal(t, accountID, account.ID)
	assert.Equal(t, briceColucciAccount.Name[0], account.Attributes.Name[0])
}

func deleteAccount(t *testing.T, account *account.AccountData) {
	t.Logf("deleting account %s ...", account.ID)

	err := api.DeleteAccount(context.Background(), account.ID, *account.Version)
	assert.Nil(t, err)

	// check that the account has been really deleted
	foundAccount, _ := api.FetchAccount(context.Background(), account.ID)
	assert.Nil(t, foundAccount)
}
