package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"produtos-favoritos/src/api/forms"
	handlers "produtos-favoritos/src/domain/interfaces/controllers"
	servicers "produtos-favoritos/src/domain/interfaces/services"
	"produtos-favoritos/src/internals/exceptions"

	"github.com/gin-gonic/gin"
)

type WishlistController struct {
	WishlistService servicers.WishlistServicer
}

func NewWishlistController(wishlistService servicers.WishlistServicer) handlers.WishlistHandler {
	return &WishlistController{WishlistService: wishlistService}
}

// WishlistProduct godoc
// @Security     ApiKeyAuth
// @Summary      Add Product To Wishlist
// @Description  Given a customer and a product add the product to the customer wishlist
// @Tags         wishlist
// @Produce      json
// @Success      200
// @Param        id path string true "Customer ID"
// @Param        wishlist  body      forms.WishlistForm  true  "WishlistForm form"
// @Router       /api/v1/customers/{id}/wishlist [post]
func (wc *WishlistController) WishlistProduct(c *gin.Context) {
	customerID := c.Param("id")

	var form forms.WishlistForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := wc.WishlistService.WishlistProduct(form.ProductID, customerID)
	if err != nil {
		var alreadyWishlistedErr *exceptions.AlreadyWishlistedErr
		if errors.As(err, &alreadyWishlistedErr) {
			c.JSON(http.StatusConflict, gin.H{"error": alreadyWishlistedErr.Error()})
			return
		}
		var notFoundEntityError *exceptions.NotFoundEntityError
		if errors.As(err, &notFoundEntityError) {
			c.JSON(http.StatusNotFound, gin.H{"error": notFoundEntityError.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to wishlist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added to wishlist"})
}

// RemoveFromWishlist godoc
// @Security     ApiKeyAuth
// @Summary      Remove Product From Wishlist
// @Description  Given a customer and a product remove the product from the wishlist
// @Tags         wishlist
// @Produce      json
// @Success      200
// @Param        id path string true "Customer ID"
// @Param        product_id path string true "Product ID"
// @Router       /api/v1/customers/{id}/wishlist/{product_id} [delete]
func (wc *WishlistController) RemoveFromWishlist(c *gin.Context) {
	customerID := c.Param("id")
	productIDParam := c.Param("product_id")

	productID, err := strconv.Atoi(productIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = wc.WishlistService.RemoveProductFromWishlist(customerID, int32(productID))
	if err != nil {
		var notFoundErr *exceptions.NotFoundEntityError
		if errors.As(err, &notFoundErr) {
			c.JSON(http.StatusNotFound, gin.H{"error": notFoundErr.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product removed from wishlist"})
}
