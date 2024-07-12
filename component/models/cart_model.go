package models

type ProductCart struct {
	ProductId int `json:"product_id,omitempty" db:"product_id" binding:"required"`
	Quantity  int `json:"quantity,omitempty" db:"quantity" binding:"required"`
	Price     int `json:"price,omitempty" db:"price" binding:"required"`
}

type Cart struct {
	Id          int           `json:"cart_id,omitempty" db:"cart_id"`
	UserId      int           `json:"user_id,omitempty" db:"user_id" binding:"required"`
	ProductCart []ProductCart `json:"product_cart,omitempty" db:"product_cart" binding:"required"`
}
