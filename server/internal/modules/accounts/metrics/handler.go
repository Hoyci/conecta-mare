package metrics

import (
	"conecta-mare-server/internal/server/middlewares"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/httphelpers"
	"conecta-mare-server/pkg/jwt"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
)

var (
	instance *metricsHandler
	Once     sync.Once
)

func NewHandler(metricsService MetricsService, accessKey string) *metricsHandler {
	Once.Do(
		func() {
			instance = &metricsHandler{
				metricsService: metricsService,
				accessKey:      accessKey,
			}
		},
	)

	return instance
}

func (h metricsHandler) RegisterRoutes(r *chi.Mux) {
	m := middlewares.NewWithAuth(h.accessKey)
	r.Route(
		"/api/v1/metrics/user-profile-views", func(r chi.Router) {
			// Private
			r.With(m.WithAuth).Get("/", h.handleGetUserProfileViews)
		},
	)
}

func (h metricsHandler) handleGetUserProfileViews(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	c, ok := ctx.Value(middlewares.AuthKey{}).(*jwt.Claims)
	if !ok {
		fmt.Println("error while attempting to get auth values from context")
		httphelpers.WriteJSON(w, http.StatusUnauthorized, exceptions.ErrUnauthorized)
	}

	query := r.URL.Query()
	startDateStr := query.Get("startDate")
	endDateStr := query.Get("endDate")

	userProfileViews, err := h.metricsService.GetUserProfileViews(ctx, c.UserID, startDateStr, endDateStr)
	if err != nil {
		httphelpers.WriteJSON(w, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		return
	}

	httphelpers.WriteJSON(w, http.StatusOK, map[string]any{"metrics": userProfileViews})
}
