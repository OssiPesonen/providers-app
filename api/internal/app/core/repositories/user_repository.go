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
	_, err := r.db.Handle().SQL().InsertInto("user").Values(models.User{
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
		Salt:     user.Salt,
	}).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Find(email string) (*models.User, error) {
	var user = &models.User{}

	q := r.db.Handle().SQL().Select("id", "username", "email", "password", "salt").From("users").Where("email = ?", email)
	if err := q.One(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Read(id int) (*models.User, error) {
	var user = &models.User{}

	q := r.db.Handle().SQL().Select("id", "username", "email", "password", "salt").From("users").Where("id = ?", id)
	if err := q.One(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) SaveRefreshToken(refreshTokenEntry *models.RefreshTokenEntry) error {
	_, err := r.db.Handle().SQL().InsertInto("sessions").Values(refreshTokenEntry).Exec()
	if err != nil {
		return err
	}

	return nil
}

// Revoke all refresh tokens by purging them from the sessions table
func (r *UserRepository) RevokeRefreshToken(token string, userId int) error {
	q := r.db.Handle().SQL().DeleteFrom("sessions").Where("token = ?", token).And("user_id = ?", userId)

	_, err := q.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) RevokeAllRefreshTokens(userId int) error {
	q := r.db.Handle().SQL().DeleteFrom("sessions").Where("user_id = ?", userId)

	_, err := q.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetRefreshToken(token string, userId int) (*models.RefreshTokenEntry, error) {
	var refreshToken = &models.RefreshTokenEntry{}

	q := r.db.Handle().SQL().Select("token", "expires", "user_id").From("sessions").Where("token = ?", token).And("user_id = ?", userId)
	if err := q.One(&refreshToken); err != nil {
		return nil, err
	}

	return refreshToken, nil
}
