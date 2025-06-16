package session

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/pkg/exceptions"
	"context"
	"log/slog"
	"net/http"
)

type service struct {
	sessionRepo SessionsRepository
	logger      *slog.Logger
}

func NewService(sessionsRepository SessionsRepository, logger *slog.Logger) SessionsService {
	return &service{
		sessionRepo: sessionsRepository,
		logger:      logger,
	}
}

func (s service) CreateSession(ctx context.Context, input common.CreateSessionRequest) (*Session, error) {
	s.logger.InfoContext(ctx, "attempting to create user session", "user_id", input.UserID)
	sess, err := New(input.UserID, input.JTI)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to process session entity", "user_id", input.UserID, "err", err)
		return nil, exceptions.MakeApiErrorWithStatus(http.StatusUnprocessableEntity, err)

	}

	err = s.sessionRepo.Create(ctx, sess)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting create session", "err", err)
		return nil, exceptions.MakeGenericApiError()
	}

	return sess, nil
}

func (s service) DeactivateAllSessions(ctx context.Context, userID string) error {
	s.logger.InfoContext(ctx, "attempting to deactivate all user sessions", "user_id", userID)
	err := s.sessionRepo.DeactivateAll(ctx, userID)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to deactivate all user sessions", "user_id", userID, "err", err)
		return exceptions.MakeGenericApiError()
	}

	return nil
}

func (s service) GetActiveSessionByUserID(ctx context.Context, userID string) (*Session, error) {
	s.logger.InfoContext(ctx, "attempting to get user active session", "user_id", userID)
	sess, err := s.sessionRepo.GetActiveByUserID(ctx, userID)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attemting to get user session", "user_id", userID, "err", err)
		return nil, exceptions.MakeGenericApiError()
	}

	return sess, nil
}

func (s service) UpdateSession(ctx context.Context, session *Session) (*Session, error) {
	s.logger.InfoContext(ctx, "attempting to update session", "user_id", session.userID, "session", session.id)
	err := s.sessionRepo.Update(ctx, session)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while to update session", "user_id", session.userID, "err", err)
		return nil, exceptions.MakeGenericApiError()
	}

	return session, nil
}
