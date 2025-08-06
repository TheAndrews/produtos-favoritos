package forms

type WishlistForm struct {
	ProductID int32 `json:"productId" binding:"required,gte=1"`
}
