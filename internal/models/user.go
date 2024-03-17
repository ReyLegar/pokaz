package models

type User struct {
	UserID       int
	Username     string
	PasswordHash string
	Role         string
}
