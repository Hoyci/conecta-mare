package users

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/server/middlewares"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/httphelpers"
	"conecta-mare-server/pkg/jwt"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/form/v4"
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

			// Private
			r.With(m.WithAuth).Patch("/logout", h.handleLogout)
		},
	)
}

var formDecoder = form.NewDecoder()

func (h userHandler) handleRegister(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	r.Body = http.MaxBytesReader(w, r.Body, 6<<20)

	if err := r.ParseMultipartForm(5 << 20); err != nil {
		apiErr := exceptions.MakeGenericApiError()
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	var formData common.RegisterUserRequest
	if err := formDecoder.Decode(&formData, r.PostForm); err != nil {
		apiErr := exceptions.MakeValidationError(err)
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	file, header, err := r.FormFile("avatar")
	if err != nil {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, exceptions.ErrAvatarEmpty)
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}
	defer file.Close()

	formData.Avatar = header

	if header.Size > 5<<20 {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, exceptions.ErrAvatarTooLarge)
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	if formData.Password != formData.ConfirmPassword {
		apiErr := exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, exceptions.ErrPasswordMatch)
		httphelpers.WriteJSON(w, apiErr.Code, apiErr)
		return
	}

	err = h.usersService.Register(ctx, formData)
	if err != nil {
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
