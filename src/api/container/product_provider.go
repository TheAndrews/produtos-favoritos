package container

import (
	"os"
	controllers "produtos-favoritos/src/api/controllers"
	handlers "produtos-favoritos/src/domain/interfaces/controllers"
	servicers "produtos-favoritos/src/domain/interfaces/services"
	services "produtos-favoritos/src/domain/services"
)

func ProvideFakeApiClient() servicers.FakeProductApiClientServicer {
	body, err := os.ReadFile("src/internals/mocks/products.json")
	if err != nil {
		panic("failed to read products mock: " + err.Error())
	}
	return services.NewFakeProductApiClientService(body)
}

func ProvideProductService(fakeApiClient servicers.FakeProductApiClientServicer) servicers.ProductServicer {
	return services.NewProductService(fakeApiClient)
}

func ProvideProductController(productService servicers.ProductServicer) handlers.ProductHandler {
	return controllers.NewProductController(productService)
}
