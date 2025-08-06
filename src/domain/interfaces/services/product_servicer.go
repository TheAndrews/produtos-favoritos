package services

import "produtos-favoritos/src/domain/models"

type ProductServicer interface {
	GetProducts() ([]models.Product, error)
	GetProductByID(productID int32) (*models.Product, error)
}
