package enums

type RefreshTokenStatus string

const (
	ActiveRefreshToken  RefreshTokenStatus = "ACTIVE"
	BanedRefeshToken    RefreshTokenStatus = "BANED"
	DeletedRefreshToken RefreshTokenStatus = "DELETED"
)
