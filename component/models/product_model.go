package models

type Product struct {
	Id     int     `json:"product_id,omitempty" db:"product_id"`
	Name   string  `json:"name,omitempty" db:"name" binding:"required"`
	Price  float64 `json:"price,omitempty" db:"price" binding:"required"`
	Stock  int     `json:"stock,omitempty" db:"stock" binding:"required"`
	ShopId int     `json:"shop_id,omitempty" db:"shop_id" binding:"required"`
	BaseModel
}
