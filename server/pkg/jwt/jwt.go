package jwt

import (
	"conecta-mare-server/pkg/uid"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateUserToken(secretKey string, claims *Claims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return jwtToken.SignedString([]byte(secretKey))
}

func GenerateClaims(id, email string, duration time.Duration) *Claims {
	jti := uid.New("jti")

	return &Claims{
		UserID: id,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			Issuer:    "conecta-mare-server",
			ID:        jti,
		},
	}
}

func Verify(secretKey string, value string) (*Claims, error) {
	if strings.TrimSpace(value) == "" {
		return nil, fmt.Errorf("invalid token")
	}

	value = strings.TrimPrefix(value, "Bearer ")

	keyFunc := func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return []byte(secretKey), nil
	}

	token, err := jwt.ParseWithClaims(value, &Claims{}, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
