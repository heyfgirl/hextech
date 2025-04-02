package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrUnknown          = errors.New("未知错误")
	ErrTokenExpired     = errors.New("token已过期")
	ErrTokenNotValidYet = errors.New("token尚未激活")
	ErrTokenMalformed   = errors.New("这不是一个token")
	ErrSignatureInvalid = errors.New("无效签名")
	ErrTokenInvalid     = errors.New("无法处理此token")
)

type JWT struct {
	SigningKey  []byte
	ExpiresTime int64
	Issuer      string
	BufferTime  int64
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID       uint
	Username string
}

func NewJWT(signingKey string, expiresTime int64, bufferTime int64, issuer string) *JWT {
	return &JWT{
		SigningKey:  []byte(signingKey),
		ExpiresTime: expiresTime,
		BufferTime:  bufferTime,
		Issuer:      issuer,
	}
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, ErrTokenExpired
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, ErrTokenMalformed
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return nil, ErrSignatureInvalid
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, ErrTokenNotValidYet
		default:
			return nil, ErrTokenInvalid
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, ErrTokenInvalid
	} else {
		return nil, ErrTokenInvalid
	}
}

// CreateClaims 创建Claims
func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: j.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),                                      // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(j.ExpiresTime))), // 过期时间 7天  配置文件
			Issuer:    j.Issuer,                                                                       // 签名的发行者
		},
	}
	return claims
}
