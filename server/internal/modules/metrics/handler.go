package metrics

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/server/middlewares"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/httphelpers"
	"conecta-mare-server/pkg/jwt"
	"encoding/json"
	"log/slog"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
)

var (
	instance *metricsHandler
	Once     sync.Once
)

func NewHandler(accessKey string, metricsService MetricsService, logger *slog.Logger) *metricsHandler {
	Once.Do(
		func() {
			instance = &metricsHandler{
				accessKey:      accessKey,
				metricsService: metricsService,
				logger:         logger,
			}
		},
	)

	return instance
}

func (h metricsHandler) RegisterRoutes(r *chi.Mux) {
	m := middlewares.NewWithAuth(h.accessKey)
	r.Route(
		"/api/v1/metrics", func(r chi.Router) {
			// Private
			r.With(m.WithAuth).Post("/event", h.handleTrackEvent)
		},
	)
}

func (h metricsHandler) handleTrackEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	h.logger.InfoContext(ctx, "attempting to create metric")

	c, ok := ctx.Value(middlewares.AuthKey{}).(*jwt.Claims)
	if !ok {
		h.logger.ErrorContext(ctx, "error while attempting to get auth values from context")
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusUnauthorized, exceptions.ErrUnauthorized)
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	h.logger.InfoContext(ctx, "destructured data from context", "user_id", c.UserID)

	var metric common.MetricRequest
	if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, exceptions.ErrInvalidJSON)
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	ok = h.metricsService.TrackEvent(ctx, metric.Type, c.UserID)

	if !ok {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusInternalServerError, exceptions.ErrInternalServerError)
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	httphelpers.WriteJSON(w, http.StatusCreated, map[string]string{"message": "metric created"})
}
