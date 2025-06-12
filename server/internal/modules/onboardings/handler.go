package onboardings

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/server/middlewares"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/httphelpers"
	"conecta-mare-server/pkg/jwt"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
)

var (
	instance *onboardingHandler
	Once     sync.Once
)

func NewHandler(service OnboardingsService, accessKey string) *onboardingHandler {
	Once.Do(
		func() {
			instance = &onboardingHandler{
				service:   service,
				accessKey: accessKey,
			}
		},
	)

	return instance
}

func (h *onboardingHandler) RegisterRoutes(r *chi.Mux) {
	m := middlewares.NewWithAuth(h.accessKey)

	r.Route("/api/v1/onboarding", func(r chi.Router) {
		r.With(m.WithAuth).Post("/", h.handleCompleteOnboarding)
	})
}

func (h *onboardingHandler) handleCompleteOnboarding(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	c, ok := ctx.Value(middlewares.AuthKey{}).(*jwt.Claims)
	if !ok {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusUnauthorized, exceptions.ErrUnauthorized)
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
	}

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, exceptions.ErrInvalidRequestBody)
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	profileData := r.FormValue("profile")
	if profileData == "" {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, fmt.Errorf("profile data is required"))
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	var req common.OnboardingRequest
	if err := json.Unmarshal([]byte(profileData), &req); err != nil {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, exceptions.ErrInvalidJSON)
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	req.UserID = c.UserID

	_, _, err := r.FormFile("profile_image")
	if err != nil {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, fmt.Errorf("profile_image is required"))
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	if err := h.service.CompleteOnboarding(ctx, &req, r); err != nil {
		var apiErr *exceptions.ApiError[string]
		if castedErr, ok := err.(*exceptions.ApiError[string]); ok {
			apiErr = castedErr
		} else {
			apiErr = exceptions.MakeApiError(err)
		}
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	httphelpers.WriteSuccess(w, http.StatusCreated)
}
