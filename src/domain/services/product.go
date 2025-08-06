package services

import (
	"encoding/json"
	"produtos-favoritos/src/domain/interfaces/services"
	"produtos-favoritos/src/domain/models"
)

type ProductService struct {
	fakeProductApiClient services.FakeProductApiClientServicer
}

func NewProductService(fakeProductApiClientServicer services.FakeProductApiClientServicer) services.ProductServicer {
	return &ProductService{
		fakeProductApiClient: fakeProductApiClientServicer,
	}
}

func (ps *ProductService) GetProducts() ([]models.Product, error) {
	body, err := ps.fakeProductApiClient.ListProducts()
	if err != nil {
		return nil, err
	}

	var products []models.Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (ps *ProductService) GetProductByID(productID int32) (*models.Product, error) {
	body, err := ps.fakeProductApiClient.GetProduct(productID)
	if err != nil {
		return nil, err
	}

	var product models.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
