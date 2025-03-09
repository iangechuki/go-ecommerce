package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iangechuki/go-ecommerce/internal/store"
)

func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("unauthorized"), "token is missing")
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("auth header is malformed"), "unauthorized")
			return
		}
		token := parts[1]
		jwtToken, err := app.accessAuthenticator.ValidateToken(token)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err, "unauthorized")
			return
		}
		claims := jwtToken.Claims.(jwt.MapClaims)
		userID, err := strconv.ParseInt(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err, "unauthorized")
			return
		}
		user, err := app.getUser(r.Context(), userID)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err, "unauthorized")
			return
		}
		ctx := context.WithValue(r.Context(), userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo3LCJ0eXBlIjoiYWNjZXNzIiwiaXNzIjoiZ28tZWNvbW1lcmNlIiwiYXVkIjpbImdvLWVjb21tZXJjZSJdLCJleHAiOjE3NDE0NjM3NzJ9.EZ3ZSwMwn2CJKZkTm6m_3KHumn6zbTtYFmz8uc9_VKo
func (app *application) getUser(ctx context.Context, userID int64) (*store.User, error) {
	user, err := app.store.Users.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
