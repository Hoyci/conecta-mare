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
	instance *onboardingsHandler
	Once     sync.Once
)

func NewHandler(
	onboardingsService OnboardingsService,
	accessKey string,
) *onboardingsHandler {
	Once.Do(
		func() {
			instance = &onboardingsHandler{
				onboardingsService: onboardingsService,
				accessKey:          accessKey,
			}
		},
	)

	return instance
}

func (h *onboardingsHandler) RegisterRoutes(r *chi.Mux) {
	m := middlewares.NewWithAuth(h.accessKey)

	r.Route("/api/v1", func(r chi.Router) {
		r.With(m.WithAuth).Post("/onboarding", h.handleCompleteOnboarding)
	})
}

func (h *onboardingsHandler) handleCompleteOnboarding(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	c, ok := ctx.Value(middlewares.AuthKey{}).(*jwt.Claims)
	if !ok {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusUnauthorized, exceptions.ErrUnauthorized)
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
	}

	// if err := r.ParseMultipartForm(32 << 20); err != nil {
	// 	apiErr := exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, exceptions.ErrInvalidRequestBody)
	// 	httphelpers.WriteJSON(w, apiErr.Code, apiErr)
	// 	return
	// }

	profileData := r.FormValue("body")
	if profileData == "" {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, fmt.Errorf("body data is required"))
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	var req common.OnboardingRequest
	if err := json.Unmarshal([]byte(profileData), &req); err != nil {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, exceptions.ErrInvalidJSON)
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	if req.SubcategoryID == "" {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, fmt.Errorf("category_id is required"))
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

	if err := h.onboardingsService.MakeOnboarding(ctx, r, &req); err != nil {
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
