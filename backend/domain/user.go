package domain

type User struct {
	ID          int    `json:"id"`
	FirebaseUID string `json:"firebase_uid"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Image       string `json:"image"`
}

type Users []User
