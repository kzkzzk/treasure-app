package domain

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Tags []Tag
