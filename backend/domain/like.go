package domain

type Like struct {
	UserID int `json:"user_id"`
	ShopID int `form:"shop_id" json:"shop_id"`
}
