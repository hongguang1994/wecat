package util

import (
	"sync"
	"time"
	"wecat/pkg/setting"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username string
	Password string
	jwt.StandardClaims
}

var (
	jwtSecret string
	timeout   int64
	once      sync.Once
)

func GenerateToken(username, password string) (string, error) {
	once.Do(func() {
		jwtSecret = setting.JwtSetting.Secret
		timeout = setting.JwtSetting.Timeout
	})

	nowTime := time.Now()

	expireTime := nowTime.Add(time.Duration(timeout) * time.Second)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
