package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const key = "*@pib7*#qik7e6oegi_pd#vyh2w!v#^bhs^k32+*cr2z&lu!nz"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(key))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(key), nil
	})
	if err != nil {
		return 0, errors.New("couldn't parse token")
	}
	isTokenValid := parsedToken.Valid
	if !isTokenValid {
		return 0, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	userId := int64(claims["userId"].(float64))
	return userId, nil
}
