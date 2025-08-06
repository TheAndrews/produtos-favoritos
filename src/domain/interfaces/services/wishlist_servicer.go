package services

type WishlistServicer interface {
	WishlistProduct(productID int32, customerID string) error
	RemoveProductFromWishlist(customerID string, productID int32) error
}
