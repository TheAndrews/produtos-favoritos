package repositories

import (
	"produtos-favoritos/src/domain/models"
)

type CustomerQuerier interface {
	Create(customer *models.Customer) error
	GetByID(id string) (*models.Customer, error)
	Update(customer *models.Customer) (*models.Customer, error)
	Delete(id string) error
	List() ([]models.Customer, error)
	RemoveProductFromWishlist(customerID string, productID int32) error
}
