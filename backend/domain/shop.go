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
	Name    string `json:"name"`
	Address string `json:"address"`
	Tel     string `json:"tel"`
	Image   string `json:"image"`
	TagIDs  []int  `json:"tag_ids"`
}
