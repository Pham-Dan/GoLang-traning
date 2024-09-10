package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(userId int64, secret string, hour int) (string, error) {
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Hour * time.Duration(hour)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(secret))
	return jwtToken, err
}

func GetUserIdFromJwtToken(token string, secret string) (int64, error) {
	hmacSecret := []byte(secret)
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		return int64(claims["id"].(float64)), nil
	}
	return 0, err
}