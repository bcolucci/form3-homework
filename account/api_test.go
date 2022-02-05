package account

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// used to fake the API host URL
const host = "http://localhost:8080"

// ramdom organisation ID used for create tests
var organisationID = uuid.NewString()

// use for create tests
var briceColucciAccount = AccountCreate{
	Name:    []string{"Brice Colucci"},
	Country: "CA",
	Bic:     "NWBKGB22",
}

// used to test request error
var errReq = errors.New("some error")

// the HTTP client mock in which you can pass whatever you want to do with the request
type httpClientMock struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (c *httpClientMock) Do(req *http.Request) (*http.Response, error) {
	// the test decides what to return here, based on the request
	return c.DoFunc(req)
}

// create tests ---------------------------------------------------------------

func TestAccountCreateSuccess(t *testing.T) {
	accountID := uuid.NewString()

	client := &httpClientMock{func(req *http.Request) (*http.Response, error) {

		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, host+"/organisation/accounts", req.URL.String())
		assert.NotNil(t, req.Body)

		var data struct{ Data accountCreateData }
		assert.Nil(t, json.NewDecoder(req.Body).Decode(&data))
		assert.NotEmpty(t, data.Data.ID)
		assert.Equal(t, organisationID, data.Data.OrganizationID)
		assert.Equal(t, "accounts", data.Data.Type)

		resBody, _ := json.Marshal(struct{ Data AccountData }{AccountData{ID: accountID}})

		return &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(bytes.NewBuffer(resBody)),
		}, nil
	}}

	api := NewAccountAPI(host, client)

	account, err := api.CreateAccount(context.Background(), organisationID, briceColucciAccount)
	assert.Nil(t, err)
	assert.Equal(t, accountID, account.ID)
}

func TestAccountCreateReqErr(t *testing.T) {
	api := NewAccountAPI(host, clientWithError())

	account, err := api.CreateAccount(context.Background(), organisationID, briceColucciAccount)
	assert.Nil(t, account)
	assert.Equal(t, "error executing request: "+errReq.Error(), err.Error())
}

func TestAccountCreateInvalidStatus(t *testing.T) {
	testAccountCreateInvalidStatus(t, http.StatusInternalServerError)
}

// fetch tests ----------------------------------------------------------------

func TestAccountFetchSuccess(t *testing.T) {
	accountID := uuid.NewString()

	client := &httpClientMock{func(req *http.Request) (*http.Response, error) {

		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, host+"/organisation/accounts/"+accountID, req.URL.String())
		assert.Nil(t, req.Body)

		resBody, _ := json.Marshal(struct{ Data AccountData }{AccountData{ID: accountID}})

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(resBody)),
		}, nil
	}}

	api := NewAccountAPI(host, client)

	account, err := api.FetchAccount(context.Background(), accountID)
	assert.Nil(t, err)
	assert.Equal(t, accountID, account.ID)
}

func TestAccountFetchReqErr(t *testing.T) {
	api := NewAccountAPI(host, clientWithError())

	account, err := api.FetchAccount(context.Background(), uuid.NewString())
	assert.Nil(t, account)
	assert.Equal(t, "error executing request: "+errReq.Error(), err.Error())
}

func TestAccountFetchInvalidStatus(t *testing.T) {
	testAccountFetchInvalidStatus(t, http.StatusInternalServerError)
}

// delete tests ---------------------------------------------------------------

func TestAccountDeleteSuccess(t *testing.T) {
	accountID := uuid.NewString()

	client := &httpClientMock{func(req *http.Request) (*http.Response, error) {

		assert.Equal(t, http.MethodDelete, req.Method)
		assert.Equal(t, host+"/organisation/accounts/"+accountID+"?version=42", req.URL.String())
		assert.Nil(t, req.Body)

		return &http.Response{StatusCode: http.StatusNoContent}, nil
	}}

	api := NewAccountAPI(host, client)

	err := api.DeleteAccount(context.Background(), accountID, 42)
	assert.Nil(t, err)
}

func TestAccountDeleteReqErr(t *testing.T) {
	api := NewAccountAPI(host, clientWithError())

	err := api.DeleteAccount(context.Background(), uuid.NewString(), 0)
	assert.Equal(t, "error executing request: "+errReq.Error(), err.Error())
}

func TestAccountDeleteNotFound(t *testing.T) {
	testAccountDeleteInvalidStatus(t, http.StatusNotFound)
}

func TestAccountDeleteConflict(t *testing.T) {
	testAccountDeleteInvalidStatus(t, http.StatusConflict)
}

// utils ----------------------------------------------------------------------

func testAccountCreateInvalidStatus(t *testing.T, status int) {
	api := NewAccountAPI(host, clientWithStatus(status))

	account, err := api.CreateAccount(context.Background(), organisationID, briceColucciAccount)
	assert.Nil(t, account)
	assert.Equal(t, fmt.Sprintf("unpextected status %d", status), err.Error())
}

func testAccountFetchInvalidStatus(t *testing.T, status int) {
	api := NewAccountAPI(host, clientWithStatus(status))

	account, err := api.FetchAccount(context.Background(), uuid.NewString())
	assert.Nil(t, account)
	assert.Equal(t, fmt.Sprintf("unpextected status %d", status), err.Error())
}

func testAccountDeleteInvalidStatus(t *testing.T, status int) {
	api := NewAccountAPI(host, clientWithStatus(status))

	err := api.DeleteAccount(context.Background(), uuid.NewString(), 0)
	assert.Equal(t, fmt.Sprintf("unpextected status %d", status), err.Error())
}

// returns an HTTP client moch which just answers with the specified HTTP status
func clientWithStatus(status int) *httpClientMock {
	return &httpClientMock{func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: status}, nil
	}}
}

// returns an HTTP client mock which just answers with a defined error
func clientWithError() *httpClientMock {
	return &httpClientMock{func(req *http.Request) (*http.Response, error) {
		return nil, errReq
	}}
}
