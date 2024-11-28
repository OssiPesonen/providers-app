package models

type Provider struct {
	Id             int    `db:"id,omitempty"`
	Name           string `db:"name"`
	City           string `db:"city"`
	Region         string `db:"region"`
	LineOfBusiness string `db:"line_of_business"`
	Keywords       string `db:"keywords"`
	UserId         int    `db:"user_id"`
}
