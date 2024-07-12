package repositories

import (
	"context"
	"database/sql"
	"main/component/models"

	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) GetAll(ctx context.Context) ([]models.Product, error) {
	strSql := `SELECT product_id, name, price, quantity, description, image FROM "product"`

	var listProducts []models.Product
	err := r.db.SelectContext(ctx, &listProducts, strSql)
	return listProducts, err
}

func (r *ProductRepo) Insert(ctx context.Context, data models.Product) (sql.Result, error) {
	strSql := `
		INSERT INTO "product" (name, price, quantity, description, image, create_at)
		VALUES ( :name, :price, :quantity, :description, :image, :create_at)
	`

	result, err := r.db.NamedExecContext(ctx, strSql, data)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *ProductRepo) Update(ctx context.Context, data models.Product) (sql.Result, error) {
	strSql := `
		UPDATE "product"
		SET name = :name, price = :price, quantity = :quantity, description = :description, image = :image, update_at = :update_at
		WHERE product_id = :product_id
	`

	result, err := r.db.NamedExecContext(ctx, strSql, data)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *ProductRepo) Delete(ctx context.Context, productID int) error {
	strSql := `DELETE FROM "product" WHERE product_id = $1`

	_, err := r.db.ExecContext(ctx, strSql, productID)

	return err
}

func (r *ProductRepo) GetByID(ctx context.Context, productID int) (models.Product, error) {
	strSql := `SELECT product_id, name, price, quantity, description, image FROM "product" WHERE product_id = $1`

	var product models.Product
	err := r.db.GetContext(ctx, &product, strSql, productID)
	return product, err
}

func (r *ProductRepo) GetByName(ctx context.Context, name string) (models.Product, error) {
	strSql := `SELECT product_id, name, price, quantity, description, image FROM "product" WHERE name = $1`

	var product models.Product
	err := r.db.GetContext(ctx, &product, strSql, name)
	return product, err
}
