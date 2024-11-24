package users

import (
	"log"

	"github.com/ossipesonen/providers-app/internal/app/core/models"
	"github.com/ossipesonen/providers-app/pkg/database"
)

type UserRepository struct {
	db            database.Database
	logger        *log.Logger
	usersTable    string
	sessionsTable string
}

// Ensure we implement interface
var _ IUserRepository = &UserRepository{}

func NewUserRepository(db database.Database, logger *log.Logger) *UserRepository {
	return &UserRepository{
		db:            db,
		logger:        logger,
		usersTable:    "users",
		sessionsTable: "sessions",
	}
}

func (r *UserRepository) Add(user *models.User) error {
	_, err := r.db.Handle().SQL().InsertInto(r.usersTable).Values(models.User{
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

	q := r.db.Handle().SQL().Select("id", "username", "email", "password", "salt").From(r.usersTable).Where("email = ?", email)
	if err := q.One(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Read(id int) (*models.User, error) {
	var user = &models.User{}

	q := r.db.Handle().SQL().Select("id", "username", "email", "password", "salt").From(r.usersTable).Where("id = ?", id)
	if err := q.One(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) SaveRefreshToken(refreshTokenEntry *models.RefreshTokenEntry) error {
	_, err := r.db.Handle().SQL().InsertInto(r.sessionsTable).Values(refreshTokenEntry).Exec()
	if err != nil {
		return err
	}

	return nil
}

// Revoke all refresh tokens by purging them from the sessions table
func (r *UserRepository) RevokeRefreshToken(token string) error {
	q := r.db.Handle().SQL().DeleteFrom(r.sessionsTable).Where("token = ?", token)

	_, err := q.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) RevokeAllRefreshTokens(userId int) error {
	q := r.db.Handle().SQL().DeleteFrom(r.sessionsTable).Where("user_id = ?", userId)

	_, err := q.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetRefreshToken(token string) (*models.RefreshTokenEntry, error) {
	var refreshToken = &models.RefreshTokenEntry{}

	q := r.db.Handle().SQL().Select("token", "expires", "user_id").From(r.sessionsTable).Where("token = ?", token)
	if err := q.One(&refreshToken); err != nil {
		return nil, err
	}

	return refreshToken, nil
}
