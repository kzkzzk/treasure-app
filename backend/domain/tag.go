package domain

type Tag struct {
	ID   int    `json:"id"`
	Name string `form:"name" json:"name"`
}

type Tags []Tag
