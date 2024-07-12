package repositories

import (
	"context"
	"database/sql"
	"main/component/models"
	"math/rand"

	"github.com/jmoiron/sqlx"
)

type ShopRepo struct {
	db *sqlx.DB
}

func NewShopRepo(db *sqlx.DB) *ShopRepo {
	return &ShopRepo{
		db: db,
	}
}

func (r *ShopRepo) GetAll(ctx context.Context) ([]models.Shop, error) {
	strSql := `SELECT shop_id, email, password, COALESCE(phone, '') as phone, full_name, avatar, status, create_at FROM "shop"`

	var listShops []models.Shop
	err := r.db.SelectContext(ctx, &listShops, strSql)
	return listShops, err
}

// get by id
func (r *ShopRepo) GetById(ctx context.Context, shopID int) (models.Shop, error) {
	strSql := `SELECT shop_id, email, password, COALESCE(phone, '') as phone, full_name, avatar, status, create_at FROM "shop" WHERE shop_id = $1`

	var shop models.Shop
	err := r.db.GetContext(ctx, &shop, strSql, shopID)
	return shop, err
}

func (r *ShopRepo) Insert(ctx context.Context, data models.Shop) (sql.Result, error) {
	strSql := `
		INSERT INTO "shop" (email, password, phone, full_name, avatar, status, create_at)
		VALUES ( :email, :password, :phone, :full_name, :avatar, :status, :create_at)
	`

	result, err := r.db.NamedExecContext(ctx, strSql, data)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *ShopRepo) GetByEmailAndPassword(ctx context.Context, email, password string) (models.Shop, error) {
	strSql := `SELECT shop_id, email, password, COALESCE(phone, '') as phone, full_name, avatar, status, create_at FROM "shop" WHERE email = $1 AND password = $2`

	var shop models.Shop
	err := r.db.GetContext(ctx, &shop, strSql, email, password)
	return shop, err
}

func (r *ShopRepo) EmailExists(ctx context.Context, email string) (bool, error) {
	strSql := `SELECT COUNT(*) FROM "shop" WHERE email = $1`

	var count int
	err := r.db.GetContext(ctx, &count, strSql, email)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *ShopRepo) GeneratePassword(ctx context.Context) (string, error) {
	const (
		tokenLength = 16
		pattern     = "[a-zA-Z0-9]" // Định nghĩa mẫu ký tự
	)

	// Tạo một chuỗi ngẫu nhiên
	password := ""
	for i := 0; i < tokenLength; i++ {
		password += string(pattern[rand.Intn(len(pattern))])
	}

	return password, nil
}

func (r *ShopRepo) ExistingEmails(ctx context.Context, emails []string) ([]string, error) {
	strSql := `SELECT email FROM "shop" WHERE email IN (?)`

	var existingEmails []string
	err := r.db.SelectContext(ctx, &existingEmails, strSql, emails)
	return existingEmails, err
}

func (r *ShopRepo) GetByEmail(ctx context.Context, email string) (models.Shop, error) {
	strSql := `SELECT shop_id, email, password, COALESCE(phone, '') as phone, full_name, avatar, status, create_at FROM "shop" WHERE email = $1`

	var shop models.Shop
	err := r.db.GetContext(ctx, &shop, strSql, email)
	return shop, err
}

func (r *ShopRepo) GetByID(ctx context.Context, shopID int) (models.Shop, error) {
	strSql := `SELECT shop_id, email, password, COALESCE(phone, '') as phone, full_name, avatar, status, create_at FROM "shop" WHERE shop_id = $1`

	var shop models.Shop
	err := r.db.GetContext(ctx, &shop, strSql, shopID)
	return shop, err
}

func (r *ShopRepo) Update(ctx context.Context, data models.Shop) (sql.Result, error) {
	strSql := `
		UPDATE "shop"
		SET email = :email, password = :password, phone = :phone, full_name = :full_name, avatar = :avatar, status = :status, update_at = :update_at
		WHERE shop_id = :shop_id
	`

	result, err := r.db.NamedExecContext(ctx, strSql, data)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *ShopRepo) Delete(ctx context.Context, shopID int) error {
	strSql := `DELETE FROM "shop" WHERE shop_id = $1`

	_, err := r.db.ExecContext(ctx, strSql, shopID)

	return err
}
