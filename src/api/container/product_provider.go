package container

import (
	"net/http"
	controllers "produtos-favoritos/src/api/controllers"
	handlers "produtos-favoritos/src/domain/interfaces/controllers"
	servicers "produtos-favoritos/src/domain/interfaces/services"
	services "produtos-favoritos/src/domain/services"
)

func ProvideFakeApiClient() servicers.FakeProductApiClientServicer {
	return services.NewFakeProductApiClientService(&http.Client{})
}

func ProvideProductService(fakeApiClient servicers.FakeProductApiClientServicer) servicers.ProductServicer {
	return services.NewProductService(fakeApiClient)
}

func ProvideProductController(productService servicers.ProductServicer) handlers.ProductHandler {
	return controllers.NewProductController(productService)
}
