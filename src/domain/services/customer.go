package services

import (
	"fmt"
	"time"

	"produtos-favoritos/src/domain/interfaces/repositories"
	"produtos-favoritos/src/domain/interfaces/services"
	"produtos-favoritos/src/domain/models"
	"produtos-favoritos/src/internals/exceptions"
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
	existingCustomer, err := s.repository.GetByEmail(customer.Email)
	if err != nil {
		return err
	}
	if existingCustomer != nil {
		return &exceptions.EmailAlreadyRegisteredErr{
			Reason: "this email is already registered",
		}
	}

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

	existingCustomerWithEmail, err := s.repository.GetByEmail(updatedCustomer.Email)
	if err != nil {
		return nil, err
	}
	if existingCustomerWithEmail != nil && updatedCustomer.Email != existingCustomer.Email {
		return nil, &exceptions.EmailAlreadyRegisteredErr{
			Reason: "this email is already registered",
		}
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
