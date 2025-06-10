package jwt

import (
	"time"
)

const (
	AccessTokenDuration  = 2 * time.Hour
	RefreshTokenDuration = 720 * time.Hour
)

type JWTProvider struct {
	accessKey  string
	refreshKey string
}

type TokenUser interface {
	ID() string
	Email() string
}

func NewProvider(accessKey, refreshKey string) *JWTProvider {
	return &JWTProvider{accessKey, refreshKey}
}

func (j *JWTProvider) generateToken(user TokenUser, key string, duration time.Duration) (*string, *Claims, error) {
	claims := GenerateClaims(user.ID(), user.Email(), duration)
	token, err := GenerateUserToken(key, claims)
	if err != nil {
		return nil, nil, err
	}
	return &token, claims, nil
}

func (j *JWTProvider) GenerateAccessToken(user TokenUser) (*string, *Claims, error) {
	return j.generateToken(user, j.accessKey, AccessTokenDuration)
}

func (j *JWTProvider) GenerateRefreshToken(user TokenUser) (*string, *Claims, error) {
	return j.generateToken(user, j.refreshKey, RefreshTokenDuration)
}

func (j *JWTProvider) VerifyAccessToken(tokenStr string) (*Claims, error) {
	return Verify(j.accessKey, tokenStr)
}

func (j *JWTProvider) VerifyRefreshToken(tokenStr string) (*Claims, error) {
	return Verify(j.refreshKey, tokenStr)
}
