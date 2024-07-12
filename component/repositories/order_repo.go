package repositories

import (
	"context"
	"database/sql"
	"main/component/models"

	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) GetAll(ctx context.Context) ([]models.Order, error) {
	strSql := `SELECT order_id, user_id, total_price FROM "order"`

	var listOrders []models.Order
	err := r.db.SelectContext(ctx, &listOrders, strSql)
	return listOrders, err
}

func (r *OrderRepo) Insert(ctx context.Context, data models.Order) (sql.Result, error) {
	strSql := `
		INSERT INTO "order" (user_id, total_price, create_at)
		VALUES ( :user_id, :total_price, :create_at)
	`

	result, err := r.db.NamedExecContext(ctx, strSql, data)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *OrderRepo) Update(ctx context.Context, data models.Order) (sql.Result, error) {
	strSql := `
		UPDATE "order"
		SET user_id = :user_id, total_price = :total_price, update_at = :update_at
		WHERE order_id = :order_id
	`

	result, err := r.db.NamedExecContext(ctx, strSql, data)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *OrderRepo) Cancel(ctx context.Context, orderID int) (sql.Result, error) {
	strSql := `DELETE FROM "order" WHERE order_id = $1`

	result, err := r.db.ExecContext(ctx, strSql, orderID)
	if err != nil {
		return nil, err
	}
	return result, err
}

// get order by user id
func (r *OrderRepo) GetByUserID(ctx context.Context, userID int) ([]models.Order, error) {
	strSql := `SELECT order_id, user_id, total_price FROM "order" WHERE user_id = $1`

	var listOrders []models.Order
	err := r.db.SelectContext(ctx, &listOrders, strSql, userID)
	return listOrders, err
}

// get order by order id
func (r *OrderRepo) GetByID(ctx context.Context, orderID int) (models.Order, error) {
	strSql := `SELECT order_id, user_id, total_price FROM "order" WHERE order_id = $1`

	var order models.Order
	err := r.db.GetContext(ctx, &order, strSql, orderID)
	return order, err
}

// get by shop id
func (r *OrderRepo) GetByShopID(ctx context.Context, shopID int) ([]models.Order, error) {
	strSql := `SELECT order_id, user_id, total_price FROM "order" WHERE user_id = $1`

	var listOrders []models.Order
	err := r.db.SelectContext(ctx, &listOrders, strSql, shopID)
	return listOrders, err
}
