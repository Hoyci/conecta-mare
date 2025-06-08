package jwt

import (
	"conecta-mare-server/pkg/uid"
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
