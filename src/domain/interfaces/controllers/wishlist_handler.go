package controllers

import "github.com/gin-gonic/gin"

type Wishlisthandler interface {
	WishlistProduct(c *gin.Context)
	RemoveFromWishlist(c *gin.Context)
}
