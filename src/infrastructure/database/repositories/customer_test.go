package repositories

import (
	"testing"

	queriers "produtos-favoritos/src/domain/interfaces/repositories"
	"produtos-favoritos/src/domain/models"

	"github.com/stretchr/testify/assert"
)

func SetupCustomerTest(t *testing.T) queriers.CustomerQuerier {
	err := TestDB.Migrator().DropTable(&models.Customer{}, &models.Product{})
	assert.NoError(t, err)

	err = TestDB.AutoMigrate(&models.Customer{}, &models.Product{})
	assert.NoError(t, err)

	return NewCustomerRepository(TestDB)
}

func TestCustomerRepository_CreateAndGetByID(t *testing.T) {
	repo := SetupCustomerTest(t)

	c := &models.Customer{
		Name:  "Customer 1",
		Email: "customer@aol.com",
	}

	err := repo.Create(c)
	assert.NoError(t, err)

	fetched, err := repo.GetByID(c.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, c.Email, fetched.Email)
	assert.Equal(t, c.Name, fetched.Name)
}

func TestCustomerRepository_Update(t *testing.T) {
	repo := SetupCustomerTest(t)

	c := &models.Customer{
		Name:  "Customer 2",
		Email: "secondcustomer@bol.com",
	}

	_ = repo.Create(c)

	c.Name = "Customer 2"
	updated, err := repo.Update(c)
	assert.NoError(t, err)
	assert.Equal(t, "Customer 2", updated.Name)
}

func TestCustomerRepository_Delete(t *testing.T) {
	repo := SetupCustomerTest(t)

	c := &models.Customer{
		Name:  "Customer to Delete",
		Email: "customerson@example.com",
	}
	_ = repo.Create(c)

	err := repo.Delete(c.ID.String())
	assert.NoError(t, err)
}

func TestCustomerRepository_List(t *testing.T) {
	repo := SetupCustomerTest(t)

	c1 := &models.Customer{Name: "Customer1", Email: "customer1@yahoo.com"}
	c2 := &models.Customer{Name: "Customer2", Email: "customer2@yahoo.com"}

	_ = repo.Create(c1)
	_ = repo.Create(c2)

	customers, err := repo.List()
	assert.NoError(t, err)
	assert.Len(t, customers, 2)
}

func TestCustomerRepository_RemoveProductFromWishlist(t *testing.T) {
	repo := SetupCustomerTest(t)

	product := &models.Product{ID: 1,
		Title:       "Produto 1",
		Price:       10.1,
		Description: "Produto Test",
		Category:    "Cat 1",
		Image:       ""}
	customer := &models.Customer{
		Name:     "Customer",
		Email:    "customer@ig.com",
		Wishlist: []*models.Product{product},
	}

	err := TestDB.Create(&product).Error
	assert.NoError(t, err)

	err = repo.Create(customer)
	assert.NoError(t, err)

	err = repo.RemoveProductFromWishlist(customer.ID.String(), product.ID)
	assert.NoError(t, err)

	fetched, err := repo.GetByID(customer.ID.String())
	assert.NoError(t, err)
	assert.Len(t, fetched.Wishlist, 0)
}
