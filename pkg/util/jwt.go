package util

import (
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
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
		jwtSecret = viper.GetString(`settings.jwt.secret`)
		timeout = viper.GetInt64(`settings.jwt.timeout`)
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
