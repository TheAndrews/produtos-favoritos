package services

import (
	querier "produtos-favoritos/src/domain/interfaces/repositories"
	servicers "produtos-favoritos/src/domain/interfaces/services"
	"produtos-favoritos/src/internals/exceptions"
)

type WishlistService struct {
	CustomerRepository querier.CustomerQuerier
	ProductService     servicers.ProductServicer
}

func NewWishlistService(customerRepository querier.CustomerQuerier,
	productService servicers.ProductServicer) servicers.WishlistServicer {
	return &WishlistService{customerRepository, productService}
}

func (ws *WishlistService) WishlistProduct(productID int32, customerID string) error {
	customer, err := ws.CustomerRepository.GetByID(customerID)
	if err != nil {
		return &exceptions.NotFoundEntityError{
			Reason: "customer not found",
		}
	}

	product, err := ws.ProductService.GetProductByID(productID)
	if err != nil {
		return &exceptions.NotFoundEntityError{
			Reason: "product not found",
		}
	}

	// Check if product already in wishlist
	for _, p := range customer.Wishlist {
		if p.ID == product.ID {
			return &exceptions.AlreadyWishlistedErr{
				Reason: "product already in wishlist",
			}
		}
	}

	// Add product to wishlist
	customer.Wishlist = append(customer.Wishlist, product)

	// Save updated customer
	ws.CustomerRepository.Update(customer)

	return err
}

func (ws *WishlistService) RemoveProductFromWishlist(customerID string, productID int32) error {
	customer, err := ws.CustomerRepository.GetByID(customerID)
	if err != nil {
		return &exceptions.NotFoundEntityError{
			Reason: "customer not found",
		}
	}

	// Check if product exists in wishlist
	index := -1
	for i, p := range customer.Wishlist {
		if p.ID == productID {
			index = i
			break
		}
	}
	if index == -1 {
		return &exceptions.NotFoundEntityError{
			Reason: "product not in wishlist",
		}
	}

	_, err = ws.ProductService.GetProductByID(productID)
	if err != nil {
		return &exceptions.NotFoundEntityError{
			Reason: "product not found",
		}
	}

	return ws.CustomerRepository.RemoveProductFromWishlist(customerID, productID)
}
