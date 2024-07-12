package enums

type ShopStatus string

const (
	ActiveShop   ShopStatus = "ACTIVE"
	InActiveShop ShopStatus = "INACTIVE"
	BanedShop    ShopStatus = "BANED"
	DeletedShop  ShopStatus = "DELETED"
)
