package domain

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Image    string
}

type Users []User
