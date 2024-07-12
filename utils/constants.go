package utils

type AuthType int
type TelegramFormatType string

const (
	AuthModeNeyu     = "neyu"
	AuthModeKeycloak = "keycloak"
)

// The collection format types for time
const (
	DateTimeFormat20060102T150405 = "2006-01-02T15:04:05"
	DateTimeFormat20060102150405  = "2006-01-02 15:04:05"
	DateTimeFormat20060102        = "2006-01-02"
	DateTimeFormat200601          = "2006-01"
)

// The collection keys context
const (
	ContextUserAuthorized = "ctx_user_authorized"
	ContextUserToken      = "ctx_user_token"
	ContextClientId       = "ctx_client_id"
)

// The collection roles
const (
	RoleManager  AuthType = 16
	RoleDirector AuthType = 24
)

// The collection external host
const (
	HostTelegram string = "https://api.telegram.org"
)

// The collection mode message telegram
const (
	ModeMarkdown   TelegramFormatType = "Markdown"
	ModeMarkdownV2 TelegramFormatType = "MarkdownV2"
	ModeHTML       TelegramFormatType = "HTML"
	ModeNone       TelegramFormatType = "None"
)

var (
	ModeText map[TelegramFormatType]bool = map[TelegramFormatType]bool{
		ModeHTML:     true,
		ModeMarkdown: true,
		ModeNone:     true,
	}
)

// permission
type Permission string

const (
	UserPermission  Permission = "USER"
	AdminPermission Permission = "ADMIN"
)
