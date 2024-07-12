package models

type ProducOrder struct {
	ProductId int `json:"product_id,omitempty" db:"product_id" binding:"required"`
	Quantity  int `json:"quantity,omitempty" db:"quantity" binding:"required"`
	Price     int `json:"price,omitempty" db:"price" binding:"required"`
}

type Order struct {
	Id         int `json:"order_id,omitempty" db:"order_id"`
	UserId     int `json:"user_id,omitempty" db:"user_id" binding:"required"`
	TotalPrice int `json:"total_price,omitempty" db:"total_price"`
	BaseModel
}

type OrderDetails struct {
	OrderId    int `json:"order_id,omitempty" db:"order_id"`
	UserId     int `json:"user_id,omitempty" db:"user_id" binding:"required"`
	ProductId  int `json:"product_id,omitempty" db:"product_id" binding:"required"`
	TotalPrice int `json:"total_price,omitempty" db:"total_price"`
}

type OrderInfo struct {
	OrderId    int           `json:"order_id,omitempty" db:"order_id"`
	UserId     int           `json:"user_id,omitempty" db:"user_id"`
	TotalPrice int           `json:"total_price,omitempty" db:"total_price"`
	Products   []ProducOrder `json:"product_order,omitempty" db:"product_order"`
}
