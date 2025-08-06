package controllers

import (
	"net/http"
	handlers "produtos-favoritos/src/domain/interfaces/controllers"
	servicers "produtos-favoritos/src/domain/interfaces/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService servicers.ProductServicer
}

func NewProductController(productService servicers.ProductServicer) handlers.ProductHandler {
	return &ProductController{ProductService: productService}
}

// GetProducts godoc
// @Security     ApiKeyAuth
// @Summary      List products
// @Description  Get all products
// @Tags         products
// @Produce      json
// @Success      200  {array}  models.Product
// @Router       /api/v1/products [get]
func (pc *ProductController) List(c *gin.Context) {
	customers, err := pc.ProductService.GetProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}
