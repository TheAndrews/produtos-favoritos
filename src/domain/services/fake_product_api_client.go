package services

import (
	"encoding/json"
	"fmt"
	"produtos-favoritos/src/domain/interfaces/services"

	"produtos-favoritos/src/domain/models"
)

type FakeProductApiClientService struct {
	Products []byte
}

func NewFakeProductApiClientService(products []byte) services.FakeProductApiClientServicer {
	return &FakeProductApiClientService{
		Products: products,
	}
}

func (fp *FakeProductApiClientService) ListProducts() ([]byte, error) {

	return fp.Products, nil
}

func (fp *FakeProductApiClientService) GetProduct(productID int32) ([]byte, error) {
	var products []models.Product
	err := json.Unmarshal(fp.Products, &products)
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