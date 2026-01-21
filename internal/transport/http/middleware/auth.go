package middleware

import (
	"context"
	"net/http"
	"strings"
	"tasker_go/internal/auth"
	"tasker_go/internal/transport/http/responder"
)

type ctxKey string

const UserIDKey ctxKey = "userID"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			responder.RespondWithError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		token := strings.TrimPrefix(h, "Bearer ")
		userID, err := auth.ParseJWT(token)
		if err != nil {
			responder.RespondWithError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
