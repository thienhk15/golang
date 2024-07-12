package repositories

import (
	"context"
	"database/sql"
	"errors"
	"main/component/models"
	"math/rand"
	"regexp"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetAll(ctx context.Context) ([]models.User, error) {
	strSql := `SELECT user_id, email, password, COALESCE(phone, '') as phone, full_name, avater, role, status, create_at  FROM "user"`

	var listUsers []models.User
	err := r.db.SelectContext(ctx, &listUsers, strSql)
	return listUsers, err
}

func (r *UserRepo) GetById(ctx context.Context, userId int) (models.User, error) {
	strSql := `SELECT user_id, email, password, COALESCE(phone, '') as phone, full_name, avater, role, status, create_at FROM "user" WHERE user_id = $1`

	var user models.User
	err := r.db.GetContext(ctx, &user, strSql, userId)
	return user, err
}

func (r *UserRepo) Insert(ctx context.Context, data models.User) (sql.Result, error) {
	strSql := `
		INSERT INTO "user" (email, password, phone, full_name, avatar, role, status, create_at)
		VALUES ( :email, :password, :phone, :full_name, :avatar, :role, :status, :create_at)
	`

	result, err := r.db.NamedExecContext(ctx, strSql, data)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *UserRepo) Update(ctx context.Context, data models.User) (sql.Result, error) {
	strSql := `
		UPDATE "user"
		SET email = :email, password = :password, phone = :phone, full_name = :full_name, avatar = :avatar, role = :role, status = :status, create_at = :create_at
		WHERE user_id = :user_id
	`
	result, err := r.db.NamedExecContext(ctx, strSql, data)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (models.User, error) {
	strSql := `SELECT user_id, email, password, COALESCE(phone, '') as phone, full_name, avater, role, status, create_at FROM "user" WHERE email = $1`

	var user models.User
	err := r.db.GetContext(ctx, &user, strSql, email)
	return user, err
}

func (r *UserRepo) GetByEmailAndPassword(ctx context.Context, email, password string) (models.User, error) {
	strSql := `SELECT user_id, email, password, COALESCE(phone, '') as phone, full_name, avater, role, status, create_at FROM "user" WHERE email = $1 AND password = $2`

	var user models.User
	err := r.db.GetContext(ctx, &user, strSql, email, password)
	return user, err
}

func (r *UserRepo) EmailExists(ctx context.Context, email string) (bool, error) {
	strSql := `SELECT COUNT(*) FROM "user" WHERE email = $1`

	var count int
	err := r.db.GetContext(ctx, &count, strSql, email)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *UserRepo) GeneratePassword(ctx context.Context) (string, error) {
	const (
		tokenLength = 16
		pattern     = "[a-zA-Z0-9]" // Định nghĩa mẫu ký tự
	)

	regex := regexp.MustCompile(pattern)

	tokenBuilder := make([]byte, tokenLength)
	for i := range tokenBuilder {
		index := rand.Intn(len(pattern) - 1)
		match := regex.FindStringSubmatch(pattern[index:])
		if len(match) == 0 {
			return "", errors.New("invalid match")
		}
		tokenBuilder[i] = match[0][0]
	}

	return string(tokenBuilder), nil
}
