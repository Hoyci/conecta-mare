package exceptions

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	ErrClientCannotContainSubcat = errors.New("client cannot have subcategory")
	ErrProfessinalWithoutSubcat  = errors.New("professional cannot have empty subcategory")
	ErrRoleEmpty                 = errors.New("role cannot be empty")
	ErrNameEmpty                 = errors.New("name cannot be empty")
	ErrEmailEmpty                = errors.New("email cannot be empty")
	ErrEmailInvalid              = errors.New("email is invalid")
	ErrPasswordEmpty             = errors.New("password cannot be empty")
	ErrPasswordMatch             = errors.New("password and confirm password doesnt match")
	ErrInvalidLoginAttempt       = errors.New("email or password is invalid")
	ErrUserNotFound              = errors.New("user was not found")
	ErrUserDisabled              = errors.New("user not active")
	ErrEmailTaken                = errors.New("email already exists")
	ErrInvalidSigningMethod      = errors.New("invalid token signing method")
	ErrUnauthorized              = errors.New("unauthorized attempt to access resource")
	ErrInternalServerError       = errors.New("unexpected error occurred")
	ErrInvalidRequestBody        = errors.New("request body is invalid")
	ErrInvalidLimitOrOffsetValue = errors.New("limit or offset value is not a valid integer")
	ErrInvalidTokenHeader        = errors.New("authorization token is malformed")
	ErrInvalidRole               = errors.New("role is invalid")
	ErrCannotFollowSelf          = errors.New("cannot follow or unfollow self")
	ErrNilInput                  = errors.New("cannot pass nil value")
	ErrAvatarEmpty               = errors.New("avatar cannot be empty")
	ErrAvatarTooLarge            = errors.New("avatar needs to be max 5mb")
	ErrActiveSessionNotFound     = errors.New("active session not found")
	ErrAccesTokenNotFound        = errors.New("access token not found")
	ErrTokenExpired              = errors.New("token expired")
	ErrCategoriesNotFound        = errors.New("any category found")
	ErrSubcategoriesNotFound     = errors.New("any subcategory found")
	ErrInvalidJSON               = errors.New("invalid json")
	ErrUserIDRequired            = errors.New("user_id is required")
)

func IsValidSqlErr(err error) bool {
	return err != nil && err != sql.ErrNoRows
}

type (
	// ApiErrorMap represents a map of application errors that can occur while processing requests.
	ApiErrorMap[T string | []string] map[string]T

	// ApiError represents the serialized error response to be propagated back to clients.
	// We include is solely for test assertions and inspection of errors for debugging, no need to report them back to users.
	ApiError[T string | []string] struct {
		Code   int            `json:"-"`
		Err    error          `json:"-"`
		Errors ApiErrorMap[T] `json:"errors"`
	}
)

func ErrorWithContext(message string, err error) error {
	return fmt.Errorf("%s: %w", message, err)
}

// Error converts an ApiError type into a valid error value.
func (err ApiError[T]) Error() string {
	serialized, _ := json.Marshal(err)
	return string(serialized)
}

// MakeValidationError converts the validator struct errors into a valid API error response type.
func MakeValidationError(validationErrors error) *ApiError[[]string] {
	apiErrors := make(map[string][]string)

	// Sift through the validation errors and append them to the error response
	for _, validationErr := range validationErrors.(validator.ValidationErrors) {
		structField := strings.ToLower(validationErr.StructField())

		// We'll only append errors to the response map if they have not already been added
		if _, exists := apiErrors[structField]; !exists {
			apiErrors[structField] = []string{MakeFriendlyError(validationErr)}
		}
	}

	return &ApiError[[]string]{
		Code:   http.StatusUnprocessableEntity,
		Errors: apiErrors,
	}
}

// MakeFriendlyError converts the field level validation messages into a user-friendly string.
func MakeFriendlyError(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fieldError.StructField())
	case "email":
		return fmt.Sprintf("%s is invalid", fieldError.Value())
	}

	return fieldError.Error()
}

func makeApiErrorMapWithMessage(err error) ApiErrorMap[string] {
	return ApiErrorMap[string]{
		"message": err.Error(),
	}
}

func MakeGenericApiError() *ApiError[string] {
	return &ApiError[string]{
		Code:   http.StatusInternalServerError,
		Errors: makeApiErrorMapWithMessage(ErrInternalServerError),
	}
}

func MakeApiError(err error) *ApiError[string] {
	return &ApiError[string]{
		Code:   http.StatusInternalServerError,
		Errors: makeApiErrorMapWithMessage(err),
	}
}

func MakeApiErrorWithFallback(currentError, fallback error) *ApiError[string] {
	var err error
	{
		if fallback != nil {
			err = fallback
		} else {
			err = currentError
		}
	}

	return &ApiError[string]{
		Code:   http.StatusInternalServerError,
		Errors: makeApiErrorMapWithMessage(err),
	}
}

func MakeApiErrorWithStatus(code int, err error) *ApiError[string] {
	return &ApiError[string]{
		Err:    err,
		Code:   code,
		Errors: makeApiErrorMapWithMessage(err),
	}
}
