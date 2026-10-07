package services

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProducts_Success(t *testing.T) {
	body, err := os.ReadFile("../../internals/mocks/products.json")
	if err != nil {
		t.Fatalf("failed to read test fixture: %v", err)
	}

	service := NewFakeProductApiClientService(body)

	result, err := service.ListProducts()

	assert.NoError(t, err)
	assert.JSONEq(t, string(body), string(result))
}

func TestGetProduct_Success(t *testing.T) {
	body, err := os.ReadFile("../../internals/mocks/products.json")
	if err != nil {
		t.Fatalf("failed to read test fixture: %v", err)
	}

	var products []json.RawMessage
	err = json.Unmarshal(body, &products)
	if err != nil {
		t.Fatalf("failed to unmarshal test fixture: %v", err)
	}

	service := NewFakeProductApiClientService(body)

	result, err := service.GetProduct(1)

	assert.NoError(t, err)
	assert.JSONEq(t, string(products[0]), string(result))
}