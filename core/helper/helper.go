package helper

import (
	"bus_api/core/define"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"math/rand"
	"strings"
	"time"
)

func GeneratePassword(data string) (string, error) {
	h := hmac.New(sha256.New, []byte(define.Salt))
	_, err := h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil)), err
}

func CheckPassword(inputPassword, secretPassword string) (bool, error) {
	h := hmac.New(sha256.New, []byte(define.Salt))
	_, err := h.Write([]byte(inputPassword))
	left := hex.EncodeToString(h.Sum(nil))
	if err != nil {
		return false, err
	}
	if left == secretPassword {
		return true, err
	}
	return false, err
}

func GenerateToken(id int, identity, name string, second int, roles []string) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(second))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// UUID 生成uuid
func UUID() string {
	return uuid.New().String()
}

func GenerateCode(x ...int) string {
	n := 4
	if len(x) == 1 {
		n = x[0]
	}
	str := "0123456789"
	sb := strings.Builder{}
	// rand.NewSource(time.Now().UnixNano())
	// for i := 0; i < n; i++ {
	// 	sb.WriteByte(str[rand.Intn(len(str))])
	// }
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		sb.WriteByte(str[r.Intn(len(str))])
	}
	return sb.String()
}

// AnalyzeToken
// Token 解析
func AnalyzeToken(token string) (*define.UserClaim, error) {
	// 判断 token 是不是以 Bearer 开头，如果是则去掉
	if strings.HasPrefix(strings.ToLower(token), "bearer ") {
		token = token[7:]
	}
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("invalid token")
	}
	return uc, err
}
