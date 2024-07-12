package models

import (
	"main/enums"
)

type User struct {
	Id       int              `json:"user_id,omitempty" db:"user_id"`
	Email    string           `json:"email,omitempty" db:"email" binding:"required"`
	Password string           `json:"password,omitempty" db:"password"`
	Phone    string           `json:"phone,omitempty" db:"phone"`
	FullName string           `json:"full_name,omitempty" db:"full_name"`
	Avatar   string           `json:"avatar,omitempty" db:"avatar"`
	Role     enums.UserRole   `json:"role,omitempty" db:"role"`
	Status   enums.UserStatus `json:"status,omitempty" db:"status"`
	BaseModel
}

type UserInsertBatchReq struct {
	Data []User `json:"data"`
}

type Request struct {
	Emails []string `json:"emails" binding:"required"`
}
