package domain

type Shop struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Tel     string `json:"tel"`
	Image   string `json:"image"`
}

type Shops []Shop

type RequestCreateShop struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
	Tel     string `form:"tel" json:"tel"`
	// Image   *multipart.FileHeader `form:"image" json:"image"`
	TagIDs []int `form:"tag_ids" json:"tag_ids"`
}

type ShopDetail struct {
	Shop
	Tags []Tag `json:"tags"`
}
