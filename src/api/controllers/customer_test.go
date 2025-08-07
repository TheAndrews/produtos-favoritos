package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"produtos-favoritos/src/api/forms"
	"produtos-favoritos/src/api/router"
	"produtos-favoritos/src/domain/models"
	"produtos-favoritos/src/infrastructure/config"
	"produtos-favoritos/src/internals/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTestRouter(t *testing.T) (*gin.Engine, *mocks.CustomerServicer) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Override config.API_KEY (since autoload might not work in tests)
	config.API_KEY = os.Getenv("API_KEY")
	gin.SetMode(gin.TestMode)
	r := gin.New()

	mockCustomerService := mocks.NewCustomerServicer(t) // Adjust if your mock package name differs
	mockCustomerController := NewCustomerController(mockCustomerService)

	productHandler := new(mocks.ProductHandler)
	wishlistHandler := new(mocks.WishlistHandler)
	// Passing nil for productController and wishlistController for now, can add mocks if needed
	router.SetupRouter(r, mockCustomerController, productHandler, wishlistHandler)

	return r, mockCustomerService
}

var mockCustomers = []models.Customer{
	{
		Name:  "John Doe",
		Email: "john@example.com",
	},
}

func TestCustomerController_List_Success(t *testing.T) {
	r, mockService := setupTestRouter(t)

	mockService.On("ListCustomers").Return(mockCustomers, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/customers/", nil)
	req.Header.Set("X-Api-Key", config.API_KEY)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	// Debug output
	fmt.Printf("Status: %d\n", resp.Code)
	fmt.Printf("Body: %s\n", resp.Body.String())

	assert.Equal(t, http.StatusOK, resp.Code)

	var respCustomers []models.Customer
	err := json.Unmarshal(resp.Body.Bytes(), &respCustomers)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	assert.Len(t, respCustomers, 1)
	mockService.AssertExpectations(t)
}

func TestCustomerController_List_Error(t *testing.T) {
	r, mockService := setupTestRouter(t)

	mockService.On("ListCustomers").Return(nil, errors.New("error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/customers/", nil)
	req.Header.Set("X-Api-Key", config.API_KEY)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	mockService.AssertExpectations(t)
}

func TestCustomerController_GetByID_Success(t *testing.T) {
	r, mockService := setupTestRouter(t)

	mockService.On("GetCustomerByID", "1").Return(&mockCustomers[0], nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/customers/1", nil)
	req.Header.Set("X-Api-Key", config.API_KEY)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var customer models.Customer
	err := json.Unmarshal(resp.Body.Bytes(), &customer)
	assert.NoError(t, err)
	assert.Equal(t, "00000000-0000-0000-0000-000000000000", customer.ID.String())
	mockService.AssertExpectations(t)
}

func TestCustomerController_GetByID_NotFound(t *testing.T) {
	r, mockService := setupTestRouter(t)

	mockService.On("GetCustomerByID", "1").Return(nil, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/customers/1", nil)
	req.Header.Set("X-Api-Key", config.API_KEY)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
	mockService.AssertExpectations(t)
}

func TestCustomerController_GetByID_Error(t *testing.T) {
	r, mockService := setupTestRouter(t)

	mockService.On("GetCustomerByID", "1").Return(nil, errors.New("db error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/customers/1", nil)
	req.Header.Set("X-Api-Key", config.API_KEY)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	mockService.AssertExpectations(t)
}

func TestCustomerController_Create_Success(t *testing.T) {
	r, mockService := setupTestRouter(t)

	form := forms.CustomerForm{Name: "New Customer", Email: "new@example.com"}
	body, _ := json.Marshal(form)

	mockService.On("CreateCustomer", mock.AnythingOfType("*models.Customer")).Return(nil)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/customers/", bytes.NewBuffer(body))
	req.Header.Set("X-Api-Key", config.API_KEY)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	mockService.AssertExpectations(t)
}

func TestCustomerController_Create_BadRequest(t *testing.T) {
	r, _ := setupTestRouter(t)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/customers/", bytes.NewBuffer([]byte(`invalid-json`)))
	req.Header.Set("X-Api-Key", config.API_KEY)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestCustomerController_Create_Error(t *testing.T) {
	r, mockService := setupTestRouter(t)

	form := forms.CustomerForm{Name: "New Customer", Email: "new@example.com"}
	body, _ := json.Marshal(form)

	mockService.On("CreateCustomer", mock.AnythingOfType("*models.Customer")).Return(errors.New("db error"))

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/customers/", bytes.NewBuffer(body))
	req.Header.Set("X-Api-Key", config.API_KEY)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	mockService.AssertExpectations(t)
}

func TestCustomerController_Delete_Success(t *testing.T) {
	r, mockService := setupTestRouter(t)

	mockService.On("DeleteCustomer", "1").Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/customers/1", nil)
	req.Header.Set("X-Api-Key", config.API_KEY)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
	mockService.AssertExpectations(t)
}

func TestCustomerController_Delete_Error(t *testing.T) {
	r, mockService := setupTestRouter(t)

	mockService.On("DeleteCustomer", "1").Return(errors.New("db error"))

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/customers/1", nil)
	req.Header.Set("X-Api-Key", config.API_KEY)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	mockService.AssertExpectations(t)
}

func TestCustomerController_Update_Success(t *testing.T) {
	r, mockService := setupTestRouter(t)

	form := forms.CustomerForm{Name: "Updated Name", Email: "updated@example.com"}
	body, _ := json.Marshal(form)

	mockService.On("UpdateCustomer", "1", mock.AnythingOfType("*models.Customer")).Return(&mockCustomers[0], nil)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/customers/1", bytes.NewBuffer(body))
	req.Header.Set("X-Api-Key", config.API_KEY)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

func TestCustomerController_Update_BadRequest(t *testing.T) {
	r, _ := setupTestRouter(t)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/customers/1", bytes.NewBuffer([]byte(`invalid-json`)))
	req.Header.Set("X-Api-Key", config.API_KEY)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestCustomerController_Update_Error(t *testing.T) {
	r, mockService := setupTestRouter(t)

	form := forms.CustomerForm{Name: "Updated Name", Email: "updated@example.com"}
	body, _ := json.Marshal(form)

	mockService.On("UpdateCustomer", "1", mock.AnythingOfType("*models.Customer")).Return(nil, errors.New("update error"))

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/customers/1", bytes.NewBuffer(body))
	req.Header.Set("X-Api-Key", config.API_KEY)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	mockService.AssertExpectations(t)
}
