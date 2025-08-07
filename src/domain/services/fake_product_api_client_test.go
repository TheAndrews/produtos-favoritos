package services

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"produtos-favoritos/src/infrastructure/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mockRoundTripper implements http.RoundTripper for testing
type mockRoundTripper struct {
	response *http.Response
	err      error
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.response, m.err
}

func makeHTTPClient(responseBody string, statusCode int, err error) *http.Client {
	return &http.Client{
		Transport: &mockRoundTripper{
			response: &http.Response{
				StatusCode: statusCode,
				Body:       ioutil.NopCloser(bytes.NewBufferString(responseBody)),
			},
			err: err,
		},
	}
}

func TestListProducts_Success(t *testing.T) {
	// Setup config for base URL if needed
	config.PRODUCTS_BASE_URL = "http://fakeapi.test"

	client := makeHTTPClient(`[{"id":1,"title":"Prod1"}]`, 200, nil)

	service := NewFakeProductApiClientService(client)
	body, err := service.ListProducts()

	assert.NoError(t, err)
	assert.JSONEq(t, `[{"id":1,"title":"Prod1"}]`, string(body))
}

func TestListProducts_HTTPError(t *testing.T) {
	client := makeHTTPClient("", 0, errors.New("network error"))

	service := NewFakeProductApiClientService(client)
	body, err := service.ListProducts()

	assert.Error(t, err)
	assert.Nil(t, body)
}

func TestListProducts_ReadBodyError(t *testing.T) {
	// Simulate response body read error by giving a Body that returns error on Read
	client := &http.Client{
		Transport: &mockRoundTripper{
			response: &http.Response{
				StatusCode: 200,
				Body:       &errorReadCloser{},
			},
			err: nil,
		},
	}

	service := NewFakeProductApiClientService(client)
	body, err := service.ListProducts()

	assert.Error(t, err)
	assert.Nil(t, body)
}

func TestGetProduct_Success(t *testing.T) {
	config.PRODUCTS_BASE_URL = "http://fakeapi.test"
	client := makeHTTPClient(`{"id":1,"title":"Prod1"}`, 200, nil)

	service := NewFakeProductApiClientService(client)
	body, err := service.GetProduct(1)

	assert.NoError(t, err)
	assert.JSONEq(t, `{"id":1,"title":"Prod1"}`, string(body))
}

func TestGetProduct_HTTPError(t *testing.T) {
	client := makeHTTPClient("", 0, errors.New("network error"))

	service := NewFakeProductApiClientService(client)
	body, err := service.GetProduct(1)

	assert.Error(t, err)
	assert.Nil(t, body)
}

func TestGetProduct_ReadBodyError(t *testing.T) {
	client := &http.Client{
		Transport: &mockRoundTripper{
			response: &http.Response{
				StatusCode: 200,
				Body:       &errorReadCloser{},
			},
			err: nil,
		},
	}

	service := NewFakeProductApiClientService(client)
	body, err := service.GetProduct(1)

	assert.Error(t, err)
	assert.Nil(t, body)
}

// errorReadCloser mocks Read error for response.Body
type errorReadCloser struct{}

func (*errorReadCloser) Read(p []byte) (int, error) {
	return 0, errors.New("read error")
}
func (*errorReadCloser) Close() error {
	return nil
}
