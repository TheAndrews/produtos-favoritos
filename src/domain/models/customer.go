package models

type Customer struct {
	BaseModel
	Name     string     `json:"name"`
	Email    string     `json:"email" gorm:"uniqueIndex"`
	Wishlist []*Product `gorm:"many2many:wishlists"`
}
