package users

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/httphelpers"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/form/v4"
)

var (
	instance *userHandler
	Once     sync.Once
)

func NewHandler(usersService UsersService) *userHandler {
	Once.Do(
		func() {
			instance = &userHandler{
				usersService: usersService,
			}
		},
	)

	return instance
}

func (h userHandler) RegisterRoutes(r *chi.Mux) {
	r.Route(
		"/api/v1/users", func(r chi.Router) {
			// Public
			r.Post("/register", h.handleRegister)
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

	httphelpers.WriteJSON(w, http.StatusCreated, common.RegisterUserResponse{Message: "success"})
}
