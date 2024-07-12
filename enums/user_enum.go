package enums

type UserStatus string
type UserRole string

const (
	ActiveUser   UserStatus = "ACTIVE"
	InActiveUser UserStatus = "INACTIVE"
	BanedUser    UserStatus = "BANED"
	DeletedUser  UserStatus = "DELETED"
)

const (
	Customer UserRole = "CUSTOMER"
	Seller   UserRole = "ADMIN"
)
