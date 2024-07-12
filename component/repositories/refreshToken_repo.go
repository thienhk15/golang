package repositories

import (
	"context"
	"database/sql"
	"main/component/models"

	"github.com/jmoiron/sqlx"
)

type RefreshTokenRepo struct {
	db *sqlx.DB
}

func NewRefreshTokenRepo(db *sqlx.DB) *RefreshTokenRepo {
	return &RefreshTokenRepo{
		db: db,
	}
}

func (r *RefreshTokenRepo) GetAll(ctx context.Context) ([]models.RefreshToken, error) {
	strSql := `SELECT token, user_id, expired_at FROM "refresh_token"`

	var listRefreshTokens []models.RefreshToken
	err := r.db.SelectContext(ctx, &listRefreshTokens, strSql)
	return listRefreshTokens, err
}

func (r *RefreshTokenRepo) Insert(ctx context.Context, data models.RefreshToken) (sql.Result, error) {
	strSql := `
		INSERT INTO "refresh_token" (token, user_id, expired_at)
		VALUES ( :token, :user_id, :expired_at)
	`

	result, err := r.db.NamedExecContext(ctx, strSql, data)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *RefreshTokenRepo) Update(ctx context.Context, data models.RefreshToken) (sql.Result, error) {
	strSql := `
		UPDATE "refresh_token"
		SET token = :token, user_id = :user_id, expired_at = :expired_at
		WHERE token = :token
	`

	result, err := r.db.NamedExecContext(ctx, strSql, data)
	if err != nil {
		return nil, err
	}
	return result, err
}

// get by token
func (r *RefreshTokenRepo) GetByToken(ctx context.Context, token string) (models.RefreshToken, error) {
	strSql := `SELECT token, user_id, expired_at FROM "refresh_token" WHERE token = $1`

	var refreshToken models.RefreshToken
	err := r.db.GetContext(ctx, &refreshToken, strSql, token)
	return refreshToken, err
}
