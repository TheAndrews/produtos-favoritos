package container

import (
	controllers "produtos-favoritos/src/api/controllers"
	services "produtos-favoritos/src/domain/services"
	repositories "produtos-favoritos/src/infrastructure/database/repositories"

	handlers "produtos-favoritos/src/domain/interfaces/controllers"
	queriers "produtos-favoritos/src/domain/interfaces/repositories"
	servicers "produtos-favoritos/src/domain/interfaces/services"

	"gorm.io/gorm"
)

func ProvideCustomerController(service servicers.CustomerServicer) handlers.CustomerHandler {
	return controllers.NewCustomerController(service) // returns *CustomerController implements CustomerHandler
}

func ProvideCustomerService(repo queriers.CustomerQuerier) servicers.CustomerServicer {
	return services.NewCustomerService(repo) // returns *CustomerService implements CustomerServicer
}

func ProvideCustomerRepository(db *gorm.DB) queriers.CustomerQuerier {
	return repositories.NewCustomerRepository(db) // returns *CustomerRepository implements CustomerQuerier
}
