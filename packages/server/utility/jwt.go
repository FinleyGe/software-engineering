package utility

import (
	"github.com/golang-jwt/jwt/v5"
	. "se/config"
	"time"
)

type Data struct {
	ID   uint64 `json:"id"`
	Role string `json:"role"`
}
type JwtData struct {
	Data Data
	jwt.RegisteredClaims
}

func GenerateToken(data Data) (string, error) {
	claims := JwtData{
		data,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(Config.Jwt.Expires))),
			Issuer:    Config.Jwt.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Config.Jwt.Secret))
}

func ParseToken(tokenString string) (*Data, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtData{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.Jwt.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtData); ok && token.Valid {
		return &claims.Data, nil
	}
	return nil, nil
}
