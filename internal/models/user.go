package models

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	CreatedAt      string `json:"created_at"`
}
