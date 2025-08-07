package services

import (
	"encoding/json"
	"errors"
	"produtos-favoritos/src/domain/models"
	"produtos-favoritos/src/internals/mocks"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProducts_Success(t *testing.T) {
	mockClient := new(mocks.FakeProductApiClientServicer)

	// Prepare mock response
	products := []models.Product{
		{ID: 1, Title: "Product A"},
		{ID: 2, Title: "Product B"},
	}
	body, _ := json.Marshal(products)

	mockClient.On("ListProducts").Return(body, nil)

	service := NewProductService(mockClient)
	result, err := service.GetProducts()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, int32(1), result[0].ID)
	assert.Equal(t, "Product A", result[0].Title)

	mockClient.AssertExpectations(t)
}

func TestGetProducts_ErrorFromAPI(t *testing.T) {
	mockClient := new(mocks.FakeProductApiClientServicer)

	mockClient.On("ListProducts").Return([]byte(nil), errors.New("API error"))

	service := NewProductService(mockClient)
	result, err := service.GetProducts()

	assert.Error(t, err)
	assert.Nil(t, result)
	mockClient.AssertExpectations(t)
}

func TestGetProducts_UnmarshalError(t *testing.T) {
	mockClient := new(mocks.FakeProductApiClientServicer)

	mockClient.On("ListProducts").Return([]byte("invalid json"), nil)

	service := NewProductService(mockClient)
	result, err := service.GetProducts()

	assert.Error(t, err)
	assert.Nil(t, result)
	mockClient.AssertExpectations(t)
}

func TestGetProductByID_Success(t *testing.T) {
	mockClient := new(mocks.FakeProductApiClientServicer)

	product := models.Product{ID: 1, Title: "Product A"}
	body, _ := json.Marshal(product)

	mockClient.On("GetProduct", int32(1)).Return(body, nil)

	service := NewProductService(mockClient)
	result, err := service.GetProductByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int32(1), result.ID)
	assert.Equal(t, "Product A", result.Title)

	mockClient.AssertExpectations(t)
}

func TestGetProductByID_ErrorFromAPI(t *testing.T) {
	mockClient := new(mocks.FakeProductApiClientServicer)

	mockClient.On("GetProduct", int32(1)).Return([]byte(nil), errors.New("API error"))

	service := NewProductService(mockClient)
	result, err := service.GetProductByID(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockClient.AssertExpectations(t)
}

func TestGetProductByID_UnmarshalError(t *testing.T) {
	mockClient := new(mocks.FakeProductApiClientServicer)

	mockClient.On("GetProduct", int32(1)).Return([]byte("not a product json"), nil)

	service := NewProductService(mockClient)
	result, err := service.GetProductByID(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockClient.AssertExpectations(t)
}
