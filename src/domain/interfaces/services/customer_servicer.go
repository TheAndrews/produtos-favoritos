package services

import "produtos-favoritos/src/domain/models"

type CustomerServicer interface {
	CreateCustomer(customer *models.Customer) error
	GetCustomerByID(id string) (*models.Customer, error)
	UpdateCustomer(id string, customer *models.Customer) (*models.Customer, error)
	DeleteCustomer(id string) error
	ListCustomers() ([]models.Customer, error)
}
