package repositories

import (
	"context"
	"database/sql"
	"main/component/models"

	"github.com/jmoiron/sqlx"
)

type CartRepo struct {
	db *sqlx.DB
}

func NewCartRepo(db *sqlx.DB) *CartRepo {
	return &CartRepo{
		db: db,
	}
}

func (r *CartRepo) GetAll(ctx context.Context) ([]models.Cart, error) {
	strSql := `SELECT cart_id, user_id, product_id, quantity, create_at FROM "cart"`

	var listCarts []models.Cart
	err := r.db.SelectContext(ctx, &listCarts, strSql)
	return listCarts, err
}

func (r *CartRepo) Insert(ctx context.Context, data models.Cart) (sql.Result, error) {
	strSql := `
		INSERT INTO "cart" (user_id, product_id, quantity, create_at)
		VALUES ( :user_id, :product_id, :quantity, :create_at)
	`

	result, err := r.db.NamedExecContext(ctx, strSql, data)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *CartRepo) Update(ctx context.Context, data models.Cart) (sql.Result, error) {
	strSql := `
		UPDATE "cart"
		SET user_id = :user_id, product_id = :product_id, quantity = :quantity, update_at = :update_at
		WHERE cart_id = :cart_id
	`

	result, err := r.db.NamedExecContext(ctx, strSql, data)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *CartRepo) Delete(ctx context.Context, cartID int) error {
	strSql := `DELETE FROM "cart" WHERE cart_id = $1`

	_, err := r.db.ExecContext(ctx, strSql, cartID)
	return err
}

func (r *CartRepo) GetByUserID(ctx context.Context, userID int) ([]models.Cart, error) {
	strSql := `SELECT cart_id, user_id, product_id, quantity, create_at FROM "cart" WHERE user_id = $1`

	var listCarts []models.Cart
	err := r.db.SelectContext(ctx, &listCarts, strSql, userID)
	return listCarts, err
}
