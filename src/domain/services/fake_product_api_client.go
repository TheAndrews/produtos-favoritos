package services

import (
	"encoding/json"
	"fmt"
	"os"
	"produtos-favoritos/src/domain/interfaces/services"

	"produtos-favoritos/src/domain/models"
)

type FakeProductApiClientService struct {
}

func NewFakeProductApiClientService() services.FakeProductApiClientServicer {
	return &FakeProductApiClientService{}
}

func (fp *FakeProductApiClientService) ListProducts() ([]byte, error) {
	body, err := os.ReadFile("src/internals/mocks/products.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read products mock: %w", err)
	}

	return body, nil
}

func (fp *FakeProductApiClientService) GetProduct(productID int32) ([]byte, error) {
	body, err := os.ReadFile("src/internals/mocks/products.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read products mock: %w", err)
	}

	var products []models.Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, fmt.Errorf("failed to parse products mock: %w", err)
	}

	for _, product := range products {
		if product.ID == productID {
			return json.Marshal(product)
		}
	}

	return nil, fmt.Errorf("product with ID %d not found", productID)
}