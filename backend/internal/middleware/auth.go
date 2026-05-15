package middleware

import (
	"backend/internal/entity"
	"backend/internal/transport"
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func JWTAuth(jwtSecret []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				transport.WriteAppError(w, entity.ErrorUnauthorized("missing authorization header"))
				return
			}
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				transport.WriteAppError(w, entity.ErrorUnauthorized("invalid authorization format"))
				return
			}
			tokenString := parts[1]
			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, entity.ErrorUnauthorized("unexpected signing method")
				}
				return jwtSecret, nil
			})
			if err != nil || !token.Valid {
				transport.WriteAppError(w, entity.ErrorUnauthorized("invalid or expired token"))
				return
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				transport.WriteAppError(w, entity.ErrorUnauthorized("invalid token claims"))
				return
			}
			userID, ok := claims["sub"].(string)
			if !ok {
				transport.WriteAppError(w, entity.ErrorUnauthorized("invalid token subject"))
				return
			}
			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserID(ctx context.Context) string {
	if v, ok := ctx.Value(UserIDKey).(string); ok {
		return v
	}
	return ""
}
