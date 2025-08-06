package services

import (
	"fmt"
	"time"

	"produtos-favoritos/src/domain/interfaces/repositories"
	"produtos-favoritos/src/domain/interfaces/services"
	"produtos-favoritos/src/domain/models"
)

type CustomerService struct {
	repository repositories.CustomerQuerier
}

// Constructor
func NewCustomerService(querier repositories.CustomerQuerier) services.CustomerServicer {
	return &CustomerService{
		repository: querier,
	}
}

func (s *CustomerService) CreateCustomer(customer *models.Customer) error {
	// Add business logic here if needed (e.g., validation)
	return s.repository.Create(customer)
}

func (s *CustomerService) GetCustomerByID(id string) (*models.Customer, error) {
	return s.repository.GetByID(id)
}

func (s *CustomerService) UpdateCustomer(id string, updatedCustomer *models.Customer) (*models.Customer, error) {
	existingCustomer, err := s.GetCustomerByID(id)
	if err != nil {
		return nil, fmt.Errorf("customer not found: %w", err)
	}

	existingCustomer.Name = updatedCustomer.Name
	existingCustomer.Email = updatedCustomer.Email
	existingCustomer.UpdatedAt = time.Now()

	return s.repository.Update(existingCustomer)
}

func (s *CustomerService) DeleteCustomer(id string) error {
	return s.repository.Delete(id)
}

func (s *CustomerService) ListCustomers() ([]models.Customer, error) {
	return s.repository.List()
}
