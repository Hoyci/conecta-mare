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
		"/users", func(r chi.Router) {
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
		exceptions.MakeGenericApiError()
		return
	}

	var formData common.RegisterUserRequest
	if err := formDecoder.Decode(&formData, r.PostForm); err != nil {
		exceptions.MakeValidationError(err)
		return
	}

	file, header, err := r.FormFile("avatar")
	if err != nil {
		exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, exceptions.ErrAvatarEmpty)
		return
	}
	defer file.Close()

	if header.Size > 5<<20 {
		exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, exceptions.ErrAvatarTooLarge)
		return
	}

	if formData.Password != formData.ConfirmPassword {
		exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, exceptions.ErrPasswordMatch)
	}

	res, err := h.usersService.Register(ctx, formData)
	if err != nil {
		exceptions.MakeApiError(err)
		return
	}

	httphelpers.WriteJSON(w, http.StatusOK, res)
}
