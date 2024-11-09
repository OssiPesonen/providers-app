package repositories

import (
	"log"

	"github.com/ossipesonen/go-traffic-lights/internal/app/core/models"
	"github.com/ossipesonen/go-traffic-lights/internal/app/core/services"
	"github.com/ossipesonen/go-traffic-lights/pkg/database"
)

type UserRepository struct {
	db     database.Database
	logger *log.Logger
	dbName string
}

// Ensure we implement interface
var _ services.UserRepository = &UserRepository{}

func NewUserRepository(db database.Database, logger *log.Logger) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: logger,
		dbName: "users",
	}
}

func (r *UserRepository) Add(user *models.User) error {
	sqlStatement := `INSERT INTO users (email, username, password, salt) VALUES ($1, $2, $3, $4);`
	_, err := r.db.Handle().Exec(sqlStatement, user.Email, user.Username, user.Password, user.Salt)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Find(email string) (*models.User, error) {
	var userId int
	var userEmail string
	var username string
	var userPassword string
	var userPasswordSalt string

	err := r.db.Handle().
		QueryRow("select id, username, email, password, salt from users where email = $1", email).
		Scan(&userId, &username, &userEmail, &userPassword, &userPasswordSalt)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:       userId,
		Username: username,
		Email:    userEmail,
		Password: userPassword,
		Salt:     userPasswordSalt,
	}, nil
}

func (r *UserRepository) Read(id int) (*models.User, error) {
	var userId int
	var userEmail string
	var username string
	var userPassword string
	var userPasswordSalt string

	err := r.db.Handle().
		QueryRow("select id, username, email, password, salt from users where id = $1", id).
		Scan(&userId, &username, &userEmail, &userPassword, &userPasswordSalt)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:       id,
		Username: username,
		Email:    userEmail,
		Password: userPassword,
		Salt:     userPasswordSalt,
	}, nil
}

func (r *UserRepository) SaveRefreshToken(refreshTokenEntry *models.RefreshTokenEntry) error {
	sqlStatement := `INSERT INTO sessions (token, user_id, expires) VALUES ($1, $2, $3);`
	_, err := r.db.Handle().Exec(sqlStatement, refreshTokenEntry.RefreshToken, refreshTokenEntry.UserId, refreshTokenEntry.Expires)

	if err != nil {
		return err
	}

	return nil
}

// Revoke all refresh tokens by purging them from the sessions table
func (r *UserRepository) RevokeRefreshToken(userId int) error {
	sqlStatement := `DELETE FROM sessions WHERE user_id = $1;`
	_, err := r.db.Handle().Exec(sqlStatement, userId)

	if err != nil {
		return err
	}

	return nil
}
