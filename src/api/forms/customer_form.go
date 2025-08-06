package forms

import "produtos-favoritos/src/domain/models"

type CustomerForm struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (f *CustomerForm) ToModel() *models.Customer {
	return &models.Customer{
		Name:  f.Name,
		Email: f.Email,
	}
}
