package middlewares

import (
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/httphelpers"
	"conecta-mare-server/pkg/jwt"
	"context"
	"net/http"
	"strings"
)

type AuthKey struct{}

type middleware struct {
	accessKey string
}

func NewWithAuth(accessKey string) *middleware {
	return &middleware{
		accessKey: accessKey,
	}
}

func (m *middleware) WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")

		if len(accessToken) == 0 {
			apiError := exceptions.MakeApiErrorWithStatus(http.StatusUnauthorized, exceptions.ErrAccesTokenNotFound)
			httphelpers.WriteJSON(w, apiError.Code, apiError)
			return
		}

		claims, err := jwt.Verify(m.accessKey, accessToken)
		if err != nil {
			if strings.Contains(err.Error(), "token has expired") {
				apiError := exceptions.MakeApiErrorWithStatus(http.StatusUnauthorized, exceptions.ErrTokenExpired)
				httphelpers.WriteJSON(w, apiError.Code, apiError)
				return
			}
			apiError := exceptions.MakeApiErrorWithStatus(http.StatusUnauthorized, exceptions.ErrInvalidTokenHeader)
			httphelpers.WriteJSON(w, apiError.Code, apiError)
			return
		}

		ctx := context.WithValue(r.Context(), AuthKey{}, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
