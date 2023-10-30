package define

import (
	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenExpire            = 3600
	RefreshTokenExpire     = 7200 * 24 * 7
	CodeLength             = 4
	GeeTestBypassStatusKey = "gt_server_bypass_status" // bypass状态存入redis时使用的key值
	ByPassUrl              = "http://bypass.geetest.com/v1/bypass_status.php"
	GeeTestCycleTime       = 100
)

var (
	AppUrl       string
	Salt         string
	JwtKey       string
	UserName     string
	MailPassword string
	GeeTestId    string
	GeeTestKey   string
)

// var JwtKey = "go-zero-pan-key"

type UserClaim struct {
	Id       int      `json:"id"`
	Identity string   `json:"identity"`
	Name     string   `json:"name"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

// PageSize 分页的默认参数
var PageSize = 20

var Datetime = "2006-01-02 15:04:05"
