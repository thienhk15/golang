package models

import (
	"main/enums"
)

type Shop struct {
	Id       int              `json:"shop_id,omitempty" db:"shop_id"`
	Email    string           `json:"email,omitempty" db:"email" binding:"required"`
	Password string           `json:"password,omitempty" db:"password"`
	Phone    string           `json:"phone,omitempty" db:"phone"`
	Name     string           `json:"full_name,omitempty" db:"full_name"`
	Avatar   string           `json:"avatar,omitempty" db:"avatar"`
	Status   enums.ShopStatus `json:"status,omitempty" db:"status"`
	BaseModel
}
