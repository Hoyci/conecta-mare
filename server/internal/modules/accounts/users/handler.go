package users

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/server/middlewares"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/httphelpers"
	"conecta-mare-server/pkg/jwt"
	"conecta-mare-server/pkg/security"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
)

var (
	instance *userHandler
	Once     sync.Once
)

func NewHandler(usersService UsersService, accessKey string) *userHandler {
	Once.Do(
		func() {
			instance = &userHandler{
				usersService: usersService,
				accessKey:    accessKey,
			}
		},
	)

	return instance
}

func (h userHandler) RegisterRoutes(r *chi.Mux) {
	m := middlewares.NewWithAuth(h.accessKey)
	r.Route(
		"/api/v1/users", func(r chi.Router) {
			// Public
			r.Post("/register", h.handleRegister)
			r.Post("/login", h.handleLogin)
			r.Get("/professionals", h.handleGetProfessionals)

			// Private
			r.With(m.WithAuth).Patch("/logout", h.handleLogout)
			r.With(m.WithAuth).Get("/", h.handleGetSigned)
		},
	)
}

func (h userHandler) handleRegister(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req common.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, exceptions.ErrInvalidJSON)
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	if req.Password != req.ConfirmPassword {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, exceptions.ErrPasswordMatch)
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	err := security.ValidatePassword(req.Password)
	if err != nil {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, err)
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	if err := h.usersService.Register(ctx, req); err != nil {
		var jsonErr map[string]any
		if json.Unmarshal([]byte(err.Error()), &jsonErr) == nil {
			httphelpers.WriteJSON(w, http.StatusBadRequest, jsonErr)
		} else {
			errorResponse := map[string]any{
				"errors": map[string]string{
					"message": err.Error(),
				},
			}
			httphelpers.WriteJSON(w, http.StatusBadRequest, errorResponse)
		}
		return
	}

	httphelpers.WriteSuccess(w, http.StatusCreated)
}

func (h userHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body common.LoginUserRequest
	err := httphelpers.ReadRequestBody(w, r, &body)
	if err != nil {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, err)
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	response, loginErr := h.usersService.Login(ctx, common.LoginUserRequest{Email: body.Email, Password: body.Password})
	if loginErr != nil {
		httphelpers.WriteJSON(w, loginErr.Code, loginErr.Err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    *response.RefreshToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		Path:     "/",
		MaxAge:   int(jwt.RefreshTokenDuration.Seconds()),
	})

	httphelpers.WriteJSON(w, http.StatusOK, response)
}

func (h userHandler) handleLogout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := h.usersService.Logout(ctx)
	if err != nil {
		httphelpers.WriteJSON(w, err.Code, err.Err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	httphelpers.WriteSuccess(w, http.StatusOK)
}

func (h userHandler) handleGetSigned(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := h.usersService.GetSigned(ctx)
	if err != nil {
		httphelpers.WriteJSON(w, err.Code, err.Error())
		return
	}

	if user == nil {
		httphelpers.WriteJSON(w, http.StatusUnauthorized, exceptions.ErrTokenExpired)
		return
	}

	httphelpers.WriteJSON(w, http.StatusOK, &common.UserResponse{
		User: &common.User{
			ID:    user.ID(),
			Email: user.Email(),
			Role:  user.Role(),
		},
	})
}

func (h userHandler) handleGetProfessionals(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	professionals, err := h.usersService.GetProfessionals(ctx)
	if err != nil {
		httphelpers.WriteJSON(w, err.Code, err.Error())
		return
	}
	if professionals == nil {
		httphelpers.WriteJSON(w, http.StatusNoContent, map[string]string{"message": "any professional found"})
		return
	}

	httphelpers.WriteJSON(w, http.StatusOK, map[string][]*common.ProfessionalResponse{"professionals": professionals})
}
