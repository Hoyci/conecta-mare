package common

import (
	"time"
)

type (
	Session struct {
		ID         string    `json:"id"`
		UserID     string    `json:"user_id"`
		JTI        string    `json:"jti"`
		Active     bool      `json:"active"`
		CreatedAt  time.Time `json:"created_at"`
		UupdatedAt time.Time `json:"updated_at"`
		ExpiresAt  time.Time `json:"expires_at"`
	}

	CreateSessionRequest struct {
		UserID string
		JTI    string
	}
)
