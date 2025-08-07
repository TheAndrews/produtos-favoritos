package controllers

import "github.com/gin-gonic/gin"

type WishlistHandler interface {
	WishlistProduct(c *gin.Context)
	RemoveFromWishlist(c *gin.Context)
}
