package container

import (
	controllers "produtos-favoritos/src/api/controllers"
	handlers "produtos-favoritos/src/domain/interfaces/controllers"
	querier "produtos-favoritos/src/domain/interfaces/repositories"
	servicers "produtos-favoritos/src/domain/interfaces/services"
	services "produtos-favoritos/src/domain/services"
)

func ProvideWishlistService(customerRepository querier.CustomerQuerier,
	productService servicers.ProductServicer) servicers.WishlistServicer {
	return services.NewWishlistService(customerRepository, productService)
}

func ProvideWishlisController(service servicers.WishlistServicer) handlers.WishlistHandler {
	return controllers.NewWishlistController(service)
}
