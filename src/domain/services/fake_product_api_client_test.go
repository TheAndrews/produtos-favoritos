package services

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProducts_Success(t *testing.T) {
	body, err := os.ReadFile("../../infrastructure/mocks/products.json")
	if err != nil {
		t.Fatalf("failed to read test fixture: %v", err)
	}

	service := NewFakeProductApiClientService()

	result, err := service.ListProducts()

	assert.NoError(t, err)
	assert.JSONEq(t, string(body), string(result))
}

func TestGetProduct_Success(t *testing.T) {
	body, err := os.ReadFile("../../infrastructure/mocks/products.json")
	if err != nil {
		t.Fatalf("failed to read test fixture: %v", err)
	}

	service := NewFakeProductApiClientService()

	result, err := service.GetProduct(1)

	assert.NoError(t, err)
	assert.JSONEq(t, string(body), string(result))
}
