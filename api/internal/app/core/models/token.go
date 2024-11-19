package models

import "time"

type RefreshTokenEntry struct {
	RefreshToken string    `db:"token"`
	Expires      time.Time `db:"expires"`
	UserId       int       `db:"user_id"`
}
