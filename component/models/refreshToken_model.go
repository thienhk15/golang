package models

import (
	"main/enums"
	"time"
)

type RefreshToken struct {
	Id        int                      `json:"id,omitempty" db:"id"`
	Token     string                   `json:"token,omitempty" db:"token"`
	UserId    int                      `json:"user_id,omitempty" db:"user_id"`
	ExpiredAt time.Time                `json:"expired_at,omitempty" db:"expired_at"`
	Status    enums.RefreshTokenStatus `json:"status,omitempty" db:"status"`
	BaseModel
}
