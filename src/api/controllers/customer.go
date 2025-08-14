package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"produtos-favoritos/src/api/forms"
	handlers "produtos-favoritos/src/domain/interfaces/controllers"
	"produtos-favoritos/src/domain/interfaces/services"
	"produtos-favoritos/src/internals/exceptions"
)

type CustomerController struct {
	BaseController
	CustomerService services.CustomerServicer
}

func NewCustomerController(service services.CustomerServicer) handlers.CustomerHandler {
	return &CustomerController{CustomerService: service}
}

// GetCustomers godoc
// @Security     ApiKeyAuth
// @Summary      List customers
// @Description  Get all customers
// @Tags         customers
// @Produce      json
// @Success      200  {array}  models.Customer
// @Router       /api/v1/customers [get]
func (cc *CustomerController) List(c *gin.Context) {
	customers, err := cc.CustomerService.ListCustomers()
	if err != nil {
		cc.respondError(c, err)
		return
	}
	cc.respond(c, customers)
}

// GetCustomerById godoc
// @Security     ApiKeyAuth
// @Summary      Get Customer by Id
// @Description  Get a single customer by Id
// @Tags         customers
// @Produce      json
// @Param        id path string true "Customer ID"
// @Success      200  {}  models.Customer
// @Router       /api/v1/customers/{id} [get]
func (cc *CustomerController) GetByID(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID"})
		return
	}
	customer, err := cc.CustomerService.GetCustomerByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if customer == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "customer not found"})
		return
	}

	cc.respond(c, customer)

}

// CreateCustomer godoc
// @Security     ApiKeyAuth
// @Summary      Create a customer
// @Description  Create a customer
// @Tags         customers
// @Produce      json
// @Param        customer  body      forms.CustomerForm  true  "CustomerForm form"
// @Success      200  {}  models.Customer
// @Router       /api/v1/customers [post]
func (cc *CustomerController) Create(c *gin.Context) {
	var form forms.CustomerForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCustomer := form.ToModel()
	err := cc.CustomerService.CreateCustomer(newCustomer)
	if err != nil {
		var emailAlreadyRegisteredErr *exceptions.EmailAlreadyRegisteredErr
		if errors.As(err, &emailAlreadyRegisteredErr) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": emailAlreadyRegisteredErr.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newCustomer)
}

// DeleteCustomer godoc
// @Security     ApiKeyAuth
// @Summary      Delete a customer
// @Description  Removes a customer
// @Tags         customers
// @Produce      json
// @Param        id path string true "Customer ID"
// @Success      204  {}  models.Customer
// @Router       /api/v1/customers/{id} [delete]
func (cc *CustomerController) Delete(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID"})
		return
	}
	err := cc.CustomerService.DeleteCustomer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete customer"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// UpdateCustomer godoc
// @Security     ApiKeyAuth
// @Summary      Update a customer
// @Description  Update a customer
// @Tags         customers
// @Produce      json
// @Param        id path string true "Customer ID"
// @Param        customer  body      forms.CustomerForm  true  "CustomerForm form"
// @Success      200  {}  models.Customer
// @Router       /api/v1/customers/{id} [put]
func (cc *CustomerController) Update(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID"})
		return
	}
	var form forms.CustomerForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customerToUpdate := form.ToModel()

	updatedCustomer, err := cc.CustomerService.UpdateCustomer(id, customerToUpdate)
	if err != nil {
		var emailAlreadyRegisteredErr *exceptions.EmailAlreadyRegisteredErr
		if errors.As(err, &emailAlreadyRegisteredErr) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": emailAlreadyRegisteredErr.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update customer"})
		return
	}

	c.JSON(http.StatusOK, updatedCustomer)
}
