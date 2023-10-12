package define

import (
	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenExpire        = 3600
	RefreshTokenExpire = 7200 * 24 * 7
	CodeLength         = 4
)

var (
	Salt         string
	JwtKey       string
	UserName     string
	MailPassword string
)

// var JwtKey = "go-zero-pan-key"

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.RegisteredClaims
}

// PageSize 分页的默认参数
var PageSize = 20

var Datetime = "2006-01-02 15:04:05"
