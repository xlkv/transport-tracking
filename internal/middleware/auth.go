package middleware

import (
	"context"
	"net/http"
	"strings"

	"tracking.xlkv.com/internal/response"
	"tracking.xlkv.com/internal/service"
)

type contextKey string

const UserID contextKey = "user_id"

func Auth(service *service.AuthService) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				response.Error(w, http.StatusUnauthorized, "token is empty")
				return
			}

			parts := strings.Split(token, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				response.Error(w, http.StatusBadRequest, "invalid token format")
				return
			}

			claims, err := service.ValidateAccessToken(r.Context(), parts[1])
			if err != nil {
				response.Error(w, http.StatusUnauthorized, "invalid token")
				return
			}

			ctx := context.WithValue(r.Context(), UserID, claims.DriverID)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
