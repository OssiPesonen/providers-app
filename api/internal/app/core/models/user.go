package models

import "time"

type User struct {
	Id       int    `db:"id,omitempty"`
	Username string `db:"username"`
	Email    string `db:"email"`
	// Hashed
	Password string `db:"password"`
	Salt     string `db:"salt"`
}

type RefreshTokenEntry struct {
	RefreshToken string    `db:"token"`
	Expires      time.Time `db:"expires"`
	UserId       int       `db:"user_id"`
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
