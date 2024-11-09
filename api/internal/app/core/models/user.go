package models

import "time"

type User struct {
	Id       int
	Username string
	Email    string
	// Hashed
	Password string
	Salt     string
}

type RefreshTokenEntry struct {
	RefreshToken string
	Expires      time.Time
	UserId       int
}

type UserCredentials struct {
	Email    string
	Password string
}

type UserInfo struct {
	Email    string
	Password string
	Username string
}
