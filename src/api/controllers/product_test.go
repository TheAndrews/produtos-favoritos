package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"produtos-favoritos/src/api/router"
	"produtos-favoritos/src/domain/models"
	"produtos-favoritos/src/infrastructure/config"
	"produtos-favoritos/src/internals/mocks"
)

// Setup test router with mock product controller
func setupProductTestRouter(t *testing.T) (*gin.Engine, *mocks.ProductServicer) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Override config.API_KEY (since autoload might not work in tests)
	config.API_KEY = os.Getenv("API_KEY")
	gin.SetMode(gin.TestMode)
	r := gin.New()

	mockProductService := mocks.NewProductServicer(t) // Adjust if your mock package name differs
	productController := NewProductController(mockProductService)

	customerHandler := new(mocks.CustomerHandler)
	wishlistHandler := new(mocks.WishlistHandler)
	// Passing nil for productController and wishlistController for now, can add mocks if needed
	router.SetupRouter(r, customerHandler, productController, wishlistHandler)

	return r, mockProductService
}

// Sample mock data
var mockProducts = []models.Product{
	{
		Title:       "Laptop",
		Description: "Powerful machine",
	},
}

func TestProductController_List_Success(t *testing.T) {
	r, mockService := setupProductTestRouter(t)

	mockService.On("GetProducts").Return(mockProducts, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/products/", nil)
	req.Header.Set("X-Api-Key", os.Getenv("API_KEY"))
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var products []models.Product
	err := json.Unmarshal(resp.Body.Bytes(), &products)
	assert.NoError(t, err)
	assert.Len(t, products, 1)
	assert.Equal(t, "Laptop", products[0].Title)
	mockService.AssertExpectations(t)
}

func TestProductController_List_Error(t *testing.T) {
	r, mockService := setupProductTestRouter(t)

	mockService.On("GetProducts").Return(nil, assert.AnError)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/products/", nil)
	req.Header.Set("X-Api-Key", os.Getenv("API_KEY"))
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	mockService.AssertExpectations(t)
}
