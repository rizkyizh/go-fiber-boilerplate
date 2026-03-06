package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserID uint   `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID uint, email, role, secret, expiry string) (string, error) {
	duration, err := time.ParseDuration(expiry)
	if err != nil {
		duration = 15 * time.Minute
	}

	claims := TokenClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GenerateRefreshToken(userID uint, email, role, secret, expiry string) (string, error) {
	duration, err := time.ParseDuration(expiry)
	if err != nil {
		duration = 168 * time.Hour
	}

	claims := TokenClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenStr, secret string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
