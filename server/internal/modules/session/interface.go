package session

import (
	"conecta-mare-server/internal/common"
	"context"
	"log/slog"
)

type (
	SessionsRepository interface {
		Create(ctx context.Context, session *Session) error
		Update(ctx context.Context, session *Session) error
		GetAllByUserID(ctx context.Context, userID string) ([]*Session, error)
		GetActiveByUserID(ctx context.Context, userID string) (*Session, error)
		GetByJTI(ctx context.Context, JTI string) (*Session, error)
		DeactivateAll(ctx context.Context, userID string) error
	}

	SessionsService interface {
		CreateSession(ctx context.Context, input common.CreateSessionRequest) (*Session, error)
		DeactivateAllSessions(ctx context.Context, userID string) error
		GetActiveSessionByUserID(ctx context.Context, userID string) (*Session, error)
		UpdateSession(ctx context.Context, session *Session) (*Session, error)
	}

	sessionService struct {
		reposityro SessionsRepository
		logger     *slog.Logger
	}
)
