package services

type FakeProductApiClientServicer interface {
	ListProducts() ([]byte, error)
	GetProduct(productID int32) ([]byte, error)
}
