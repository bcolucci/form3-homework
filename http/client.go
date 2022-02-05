package http

import (
	"net/http"
	"time"
)

// HttpClient is an HTTP client interface so we can create mock for tests
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// if the context does not set any timeout, we should put one there
var client = http.Client{Timeout: 10 * time.Second}

// TimeoutHttpClient is the exposed http client with a default timeout
type TimeoutHttpClient struct{}

func (c *TimeoutHttpClient) Do(req *http.Request) (*http.Response, error) {
	return client.Do(req)
}
