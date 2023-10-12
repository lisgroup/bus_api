package define

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
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

// ObjectStorageType 对象存储类型
// 支持 minio\cos
var ObjectStorageType = os.Getenv("ObjectStorageType")

// TencentSecretKey 腾讯云对象存储
var TencentSecretKey = os.Getenv("TencentSecretKey")
var TencentSecretID = os.Getenv("TencentSecretID")
var CosBucket = "https://getcharzp-1256268070.cos.ap-chengdu.myqcloud.com"

// MinIOAccessKeyID MinIO 配置
var MinIOAccessKeyID = os.Getenv("MinIOAccessKeyID")
var MinIOAccessSecretKey = os.Getenv("MinIOAccessSecretKey")
var MinIOEndpoint = os.Getenv("MinIOEndpoint")
var MinIOBucket = os.Getenv("MinIOBucket")

// PageSize 分页的默认参数
var PageSize = 20

var Datetime = "2006-01-02 15:04:05"
