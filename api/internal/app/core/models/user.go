package models

type User struct {
	Id       int    `db:"id,omitempty"`
	Username string `db:"username"`
	Email    string `db:"email"`
	// Hashed
	Password string `db:"password"`
	Salt     string `db:"salt"`
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
