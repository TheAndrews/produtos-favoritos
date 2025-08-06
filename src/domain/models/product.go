package models

type Product struct {
	ID          int32   `json:"id"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
}
