package services

import (
	"fmt"
	"io"
	"net/http"
	"produtos-favoritos/src/domain/interfaces/services"
	"produtos-favoritos/src/infrastructure/config"
)

type FakeProductApiClientService struct {
	HTTP *http.Client
}

func NewFakeProductApiClientService(httpClient *http.Client) services.FakeProductApiClientServicer {
	return &FakeProductApiClientService{
		HTTP: httpClient,
	}
}

func (fp *FakeProductApiClientService) ListProducts() ([]byte, error) {
	listProductsUrl := fmt.Sprintf("%s%s", config.PRODUCTS_BASE_URL, "/products")
	request, _ := http.NewRequest("GET", listProductsUrl, nil)
	response, err := fp.HTTP.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func (fp *FakeProductApiClientService) GetProduct(productID int32) ([]byte, error) {
	getProductUrl := fmt.Sprintf("%s%s%d", config.PRODUCTS_BASE_URL, "/products/", productID)
	request, _ := http.NewRequest("GET", getProductUrl, nil)
	response, err := fp.HTTP.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
