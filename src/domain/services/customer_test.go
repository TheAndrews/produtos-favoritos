package services

import (
	"errors"
	"produtos-favoritos/src/domain/models"
	"produtos-favoritos/src/internals/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCustomer_Success(t *testing.T) {
	mockRepo := new(mocks.CustomerQuerier)

	customer := &models.Customer{
		BaseModel: models.BaseModel{ID: uuid.New()},
		Name:      "Customer Create",
		Email:     "customer@create.com",
	}

	mockRepo.On("Create", customer).Return(nil)
	mockRepo.On("GetByEmail", customer.Email).Return(nil, nil)

	service := NewCustomerService(mockRepo)
	err := service.CreateCustomer(customer)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetCustomerByID_Success(t *testing.T) {
	mockRepo := new(mocks.CustomerQuerier)

	customerID := uuid.New().String()
	expectedCustomer := &models.Customer{
		BaseModel: models.BaseModel{ID: uuid.MustParse(customerID)},
		Name:      "Customer 1",
		Email:     "customer@ig.com",
	}

	mockRepo.On("GetByID", customerID).Return(expectedCustomer, nil)

	service := NewCustomerService(mockRepo)
	result, err := service.GetCustomerByID(customerID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, result)
	mockRepo.AssertExpectations(t)
}

func TestGetCustomerByID_NotFound(t *testing.T) {
	mockRepo := new(mocks.CustomerQuerier)

	customerID := uuid.New().String()

	mockRepo.On("GetByID", customerID).Return(nil, errors.New("record not found"))

	service := NewCustomerService(mockRepo)
	result, err := service.GetCustomerByID(customerID)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateCustomer_Success(t *testing.T) {
	mockRepo := new(mocks.CustomerQuerier)

	customerID := uuid.New().String()
	existingCustomer := &models.Customer{
		BaseModel: models.BaseModel{ID: uuid.MustParse(customerID)},
		Name:      "Customer",
		Email:     "customer@bol.com",
	}
	updatedCustomer := &models.Customer{
		Name:  "Customer Updated",
		Email: "customerupdated@uol.com",
	}

	mockRepo.On("GetByID", customerID).Return(existingCustomer, nil)
	mockRepo.On("GetByEmail", updatedCustomer.Email).Return(nil, nil)
	// Update returns updated customer and nil error
	mockRepo.On("Update", mock.AnythingOfType("*models.Customer")).
		Return(existingCustomer, nil).
		Run(func(args mock.Arguments) {
			cust := args.Get(0).(*models.Customer)
			cust.Name = updatedCustomer.Name
			cust.Email = updatedCustomer.Email
			cust.UpdatedAt = time.Now()
		})

	service := NewCustomerService(mockRepo)
	result, err := service.UpdateCustomer(customerID, updatedCustomer)

	assert.NoError(t, err)
	assert.Equal(t, updatedCustomer.Name, result.Name)
	assert.Equal(t, updatedCustomer.Email, result.Email)
	assert.NotZero(t, result.UpdatedAt)
	mockRepo.AssertExpectations(t)
}

func TestUpdateCustomer_NotFound(t *testing.T) {
	mockRepo := new(mocks.CustomerQuerier)

	customerID := uuid.New().String()
	updatedCustomer := &models.Customer{
		Name:  "Customer",
		Email: "customer@uol.com",
	}

	mockRepo.On("GetByID", customerID).Return(nil, errors.New("record not found"))

	service := NewCustomerService(mockRepo)
	result, err := service.UpdateCustomer(customerID, updatedCustomer)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteCustomer_Success(t *testing.T) {
	mockRepo := new(mocks.CustomerQuerier)

	customerID := uuid.New().String()

	mockRepo.On("Delete", customerID).Return(nil)

	service := NewCustomerService(mockRepo)
	err := service.DeleteCustomer(customerID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteCustomer_Error(t *testing.T) {
	mockRepo := new(mocks.CustomerQuerier)

	customerID := uuid.New().String()

	mockRepo.On("Delete", customerID).Return(errors.New("delete failed"))

	service := NewCustomerService(mockRepo)
	err := service.DeleteCustomer(customerID)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestListCustomers_Success(t *testing.T) {
	mockRepo := new(mocks.CustomerQuerier)

	customers := []models.Customer{
		{Name: "Customer One", Email: "customer@google.com"},
		{Name: "Customer Two", Email: "customer@msn.com"},
	}

	mockRepo.On("List").Return(customers, nil)

	service := NewCustomerService(mockRepo)
	result, err := service.ListCustomers()

	assert.NoError(t, err)
	assert.Equal(t, customers, result)
	mockRepo.AssertExpectations(t)
}

func TestListCustomers_Error(t *testing.T) {
	mockRepo := new(mocks.CustomerQuerier)

	mockRepo.On("List").Return(nil, errors.New("list failed"))

	service := NewCustomerService(mockRepo)
	result, err := service.ListCustomers()

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
