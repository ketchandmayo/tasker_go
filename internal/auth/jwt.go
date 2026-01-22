package auth

import (
	"errors"
	"tasker_go/internal/config/jwt_config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = jwt_config.Secret

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ParseJWT(tokenStr string) (uint, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		jwt.MapClaims{},
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return secret, nil
		},
	)

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id missing or invalid")
	}

	return uint(userIDFloat), nil
}
