package container

import (
	"go.uber.org/dig"

	"github.com/gin-gonic/gin"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// inject web framework
	container.Provide(gin.Default)

	// inject db
	container.Provide(ProvideGormDB)

	// inject Repositories
	container.Provide(ProvideCustomerRepository)

	// inject Services
	container.Provide(ProvideCustomerService)
	container.Provide(ProvideFakeApiClient)
	container.Provide(ProvideProductService)
	container.Provide(ProvideWishlistService)

	// inject Controllers
	container.Provide(ProvideCustomerController)
	container.Provide(ProvideProductController)
	container.Provide(ProvideWishlisController)

	return container
}
