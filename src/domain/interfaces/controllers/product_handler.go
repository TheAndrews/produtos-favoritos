package controllers

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	List(c *gin.Context)
}
