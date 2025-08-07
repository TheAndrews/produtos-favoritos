package services

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"produtos-favoritos/src/domain/models"
	"produtos-favoritos/src/internals/exceptions"
	"produtos-favoritos/src/internals/mocks"
)

func createCustomer(id uuid.UUID, wishlist []*models.Product) *models.Customer {
	return &models.Customer{
		BaseModel: models.BaseModel{ID: id, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:      "Test User",
		Email:     "test@test.com",
		Wishlist:  wishlist,
	}
}

func createProduct(id int32) *models.Product {
	return &models.Product{
		ID:    id,
		Title: "Test Product",
	}
}

func TestWishlistProduct_Success(t *testing.T) {
	customerRepo := new(mocks.CustomerQuerier)
	productSvc := new(mocks.ProductServicer)

	customerID := uuid.New()
	productID := int32(1)

	customer := createCustomer(customerID, []*models.Product{})
	product := createProduct(productID)

	customerRepo.On("GetByID", customerID.String()).Return(customer, nil)
	productSvc.On("GetProductByID", productID).Return(product, nil)
	customerRepo.On("Update", mock.Anything).Maybe().Return(customer, nil)

	service := NewWishlistService(customerRepo, productSvc)

	err := service.WishlistProduct(productID, customerID.String())

	assert.NoError(t, err)
	assert.Len(t, customer.Wishlist, 1)
	assert.Equal(t, productID, customer.Wishlist[0].ID)
	customerRepo.AssertExpectations(t)
	productSvc.AssertExpectations(t)
}

func TestWishlistProduct_AlreadyWishlisted(t *testing.T) {
	customerRepo := new(mocks.CustomerQuerier)
	productSvc := new(mocks.ProductServicer)

	customerID := uuid.New()
	productID := int32(1)
	product := createProduct(productID)
	customer := createCustomer(customerID, []*models.Product{product})

	customerRepo.On("GetByID", customerID.String()).Return(customer, nil)
	productSvc.On("GetProductByID", productID).Return(product, nil)

	service := NewWishlistService(customerRepo, productSvc)

	err := service.WishlistProduct(productID, customerID.String())

	assert.Error(t, err)
	assert.IsType(t, &exceptions.AlreadyWishlistedErr{}, err)
}

func TestWishlistProduct_CustomerNotFound(t *testing.T) {
	customerRepo := new(mocks.CustomerQuerier)
	productSvc := new(mocks.ProductServicer)

	customerID := uuid.New()
	customerRepo.On("GetByID", customerID.String()).Return(nil, errors.New("not found"))

	service := NewWishlistService(customerRepo, productSvc)

	err := service.WishlistProduct(1, customerID.String())

	assert.Error(t, err)
	assert.IsType(t, &exceptions.NotFoundEntityError{}, err)
}

func TestWishlistProduct_ProductNotFound(t *testing.T) {
	customerRepo := new(mocks.CustomerQuerier)
	productSvc := new(mocks.ProductServicer)

	customerID := uuid.New()
	customer := createCustomer(customerID, []*models.Product{})
	customerRepo.On("GetByID", customerID.String()).Return(customer, nil)
	productSvc.On("GetProductByID", int32(1)).Return(nil, errors.New("not found"))

	service := NewWishlistService(customerRepo, productSvc)

	err := service.WishlistProduct(1, customerID.String())

	assert.Error(t, err)
	assert.IsType(t, &exceptions.NotFoundEntityError{}, err)
}

func TestRemoveProductFromWishlist_Success(t *testing.T) {
	customerRepo := new(mocks.CustomerQuerier)
	productSvc := new(mocks.ProductServicer)

	customerID := uuid.New()
	productID := int32(1)
	product := createProduct(productID)
	customer := createCustomer(customerID, []*models.Product{product})

	customerRepo.On("GetByID", customerID.String()).Return(customer, nil)
	productSvc.On("GetProductByID", productID).Return(product, nil)
	customerRepo.On("RemoveProductFromWishlist", customerID.String(), productID).Return(nil)

	service := NewWishlistService(customerRepo, productSvc)

	err := service.RemoveProductFromWishlist(customerID.String(), productID)

	assert.NoError(t, err)
}

func TestRemoveProductFromWishlist_CustomerNotFound(t *testing.T) {
	customerRepo := new(mocks.CustomerQuerier)
	productSvc := new(mocks.ProductServicer)

	customerID := uuid.New()
	customerRepo.On("GetByID", customerID.String()).Return(nil, errors.New("not found"))

	service := NewWishlistService(customerRepo, productSvc)

	err := service.RemoveProductFromWishlist(customerID.String(), 1)

	assert.Error(t, err)
	assert.IsType(t, &exceptions.NotFoundEntityError{}, err)
}

func TestRemoveProductFromWishlist_ProductNotInWishlist(t *testing.T) {
	customerRepo := new(mocks.CustomerQuerier)
	productSvc := new(mocks.ProductServicer)

	customerID := uuid.New()
	customer := createCustomer(customerID, []*models.Product{})

	customerRepo.On("GetByID", customerID.String()).Return(customer, nil)

	service := NewWishlistService(customerRepo, productSvc)

	err := service.RemoveProductFromWishlist(customerID.String(), 1)

	assert.Error(t, err)
	assert.IsType(t, &exceptions.NotFoundEntityError{}, err)
}
