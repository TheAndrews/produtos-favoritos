package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"produtos-favoritos/src/api/forms"
	"produtos-favoritos/src/api/router"
	"produtos-favoritos/src/infrastructure/config"
	"produtos-favoritos/src/internals/exceptions"
	"produtos-favoritos/src/internals/mocks"
)

func setupWishlistTestRouter(t *testing.T) (*gin.Engine, *mocks.WishlistServicer) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Override config.API_KEY (since autoload might not work in tests)
	config.API_KEY = os.Getenv("API_KEY")
	gin.SetMode(gin.TestMode)
	r := gin.New()
	wishlistService := new(mocks.WishlistServicer)
	wishlistController := NewWishlistController(wishlistService)

	customerHandler := new(mocks.CustomerHandler)
	productHandler := new(mocks.ProductHandler)
	// Passing nil for productController and wishlistController for now, can add mocks if needed
	router.SetupRouter(r, customerHandler, productHandler, wishlistController)

	return r, wishlistService
}

// ----------------------
// WishlistProduct Tests
// ----------------------

func TestWishlistController_WishlistProduct_Success(t *testing.T) {
	r, mockService := setupWishlistTestRouter(t)

	form := forms.WishlistForm{ProductID: 123}
	body, _ := json.Marshal(form)

	mockService.On("WishlistProduct", int32(123), "1").Return(nil)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/customers/1/wishlist", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", config.API_KEY)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

func TestWishlistController_WishlistProduct_AlreadyWishlisted(t *testing.T) {
	r, mockService := setupWishlistTestRouter(t)

	form := forms.WishlistForm{ProductID: 123}
	body, _ := json.Marshal(form)

	mockService.On("WishlistProduct", int32(123), "1").Return(&exceptions.AlreadyWishlistedErr{Reason: "Already in wishlist"})

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/customers/1/wishlist", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", config.API_KEY)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusConflict, resp.Code)
	mockService.AssertExpectations(t)
}

func TestWishlistController_WishlistProduct_NotFound(t *testing.T) {
	r, mockService := setupWishlistTestRouter(t)

	form := forms.WishlistForm{ProductID: 123}
	body, _ := json.Marshal(form)

	mockService.On("WishlistProduct", int32(123), "1").Return(&exceptions.NotFoundEntityError{Reason: "Not found"})

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/customers/1/wishlist", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", config.API_KEY)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
	mockService.AssertExpectations(t)
}

func TestWishlistController_WishlistProduct_BadRequest(t *testing.T) {
	r, _ := setupWishlistTestRouter(t)

	// Invalid JSON
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/customers/1/wishlist", bytes.NewBuffer([]byte(`invalid`)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", config.API_KEY)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestWishlistController_WishlistProduct_InternalServerError(t *testing.T) {
	r, mockService := setupWishlistTestRouter(t)

	form := forms.WishlistForm{ProductID: 123}
	body, _ := json.Marshal(form)

	mockService.On("WishlistProduct", int32(123), "1").Return(errors.New("something went wrong"))

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/customers/1/wishlist", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", config.API_KEY)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	mockService.AssertExpectations(t)
}

// ----------------------
// RemoveFromWishlist Tests
// ----------------------

func TestWishlistController_RemoveFromWishlist_Success(t *testing.T) {
	r, mockService := setupWishlistTestRouter(t)

	mockService.On("RemoveProductFromWishlist", "1", int32(123)).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/customers/1/wishlist/123", nil)
	req.Header.Set("X-Api-Key", config.API_KEY)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

func TestWishlistController_RemoveFromWishlist_NotFound(t *testing.T) {
	r, mockService := setupWishlistTestRouter(t)

	mockService.On("RemoveProductFromWishlist", "1", int32(123)).Return(&exceptions.NotFoundEntityError{Reason: "not found"})

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/customers/1/wishlist/123", nil)
	req.Header.Set("X-Api-Key", config.API_KEY)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
	mockService.AssertExpectations(t)
}

func TestWishlistController_RemoveFromWishlist_InvalidID(t *testing.T) {
	r, _ := setupWishlistTestRouter(t)

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/customers/1/wishlist/abc", nil)
	req.Header.Set("X-Api-Key", config.API_KEY)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestWishlistController_RemoveFromWishlist_InternalServerError(t *testing.T) {
	r, mockService := setupWishlistTestRouter(t)

	mockService.On("RemoveProductFromWishlist", "1", int32(123)).Return(errors.New("db down"))

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/customers/1/wishlist/123", nil)
	req.Header.Set("X-Api-Key", config.API_KEY)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	mockService.AssertExpectations(t)
}
