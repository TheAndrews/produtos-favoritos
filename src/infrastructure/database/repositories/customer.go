package repositories

import (
	interfaces "produtos-favoritos/src/domain/interfaces/repositories"
	"produtos-favoritos/src/domain/models"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) interfaces.CustomerQuerier {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Create(customer *models.Customer) error {
	return r.db.Create(customer).Error
}

func (r *CustomerRepository) GetByID(id string) (*models.Customer, error) {
	var customer models.Customer
	if err := r.db.Preload("Wishlist").First(&customer, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) Update(customer *models.Customer) (*models.Customer, error) {
	err := r.db.Save(customer).Error
	return customer, err
}

func (r *CustomerRepository) Delete(id string) error {
	return r.db.Delete(&models.Customer{}, "id = ?", id).Error
}

func (r *CustomerRepository) List() ([]models.Customer, error) {
	var customers []models.Customer
	if err := r.db.Preload("Wishlist").Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *CustomerRepository) RemoveProductFromWishlist(customerID string, productID int32) error {
	var customer models.Customer
	if err := r.db.First(&customer, "id = ?", customerID).Error; err != nil {
		return err
	}

	product := models.Product{ID: productID}

	return r.db.Model(&customer).Association("Wishlist").Delete(&product)
}
