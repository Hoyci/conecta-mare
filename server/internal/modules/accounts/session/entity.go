package session

import (
	"conecta-mare-server/internal/database/models"
	"conecta-mare-server/pkg/uid"
	"fmt"
	"time"
)

const (
	ttl = time.Hour * 24 * 30
)

type Session struct {
	id        string
	userID    string
	jti       string
	active    bool
	createdAt time.Time
	updatedAt time.Time
	expiresAt time.Time
}

func New(userID, JTI string) (*Session, error) {
	if userID == "" || JTI == "" {
		return nil, fmt.Errorf("userID and JTI are required")
	}

	return &Session{
		id:        uid.New("sess"),
		userID:    userID,
		jti:       JTI,
		active:    true,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		expiresAt: time.Now().Add(ttl),
	}, nil
}

func NewFromModel(m models.Session) *Session {
	return &Session{
		id:        m.ID,
		userID:    m.UserID,
		jti:       m.JTI,
		active:    m.Active,
		createdAt: m.CreatedAt,
		updatedAt: m.UpdatedAt,
		expiresAt: m.ExpiresAt,
	}
}

func (s *Session) ToModel() models.Session {
	return models.Session{
		ID:        s.id,
		UserID:    s.userID,
		JTI:       s.jti,
		Active:    s.active,
		CreatedAt: s.createdAt,
		UpdatedAt: s.updatedAt,
		ExpiresAt: s.expiresAt,
	}
}

func (s *Session) IsExpired() bool {
	return s.expiresAt.Before(time.Now())
}

func (s *Session) ChangeJTI(JTI string) {
	s.jti = JTI
	s.updatedAt = time.Now()
}

func (s *Session) Activate() {
	s.active = true
	s.updatedAt = time.Now()
}

func (s *Session) Deactivate() {
	s.active = false
	s.updatedAt = time.Now()
}
