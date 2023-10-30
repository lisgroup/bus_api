package middleware

import (
	"bus_api/core/helper"
	"context"
	"net/http"
	"strconv"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Unauthorized"))
			return
		}
		uc, err := helper.AnalyzeToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "id", uc.Id)
		ctx = context.WithValue(ctx, "identity", uc.Identity)
		ctx = context.WithValue(ctx, "name", uc.Name)
		r.Header.Set("user_id", strconv.Itoa(uc.Id))
		// r.Header.Set("user_identity", uc.Identity)
		// r.Header.Set("user_name", uc.Name)
		next(w, r.WithContext(ctx))
	}
}
