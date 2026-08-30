package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/luan-nguyen-huu/Adam/internal/domain"
	"github.com/luan-nguyen-huu/Adam/internal/utils"
	"github.com/luan-nguyen-huu/Adam/internal/utils/jwt"
)

type contextKey string

const UserClaimsContextKey contextKey = "adam_user_claims"

// AuthMiddleware extracts and validates access token from either Authorization header or Cookie.
func AuthMiddleware(tokenMaker jwt.JWTMakerInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := extractToken(r, "access_token")
			if tokenString == "" {
				utils.WriteErrorResponse(w, http.StatusUnauthorized, domain.ErrMissingAuthToken.Error())
				return
			}

			claims, err := tokenMaker.VerifyAccessToken(tokenString)
			if err != nil {
				utils.WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
				return
			}

			ctx := context.WithValue(r.Context(), UserClaimsContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RefreshTokenMiddleware extracts and validates refresh token from either Cookie or Authorization header.
func RefreshTokenMiddleware(tokenMaker jwt.JWTMakerInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := extractToken(r, "refresh_token")
			if tokenString == "" {
				utils.WriteErrorResponse(w, http.StatusUnauthorized, domain.ErrMissingAuthToken.Error())
				return
			}

			claims, err := tokenMaker.VerifyRefreshToken(tokenString)
			if err != nil {
				utils.WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
				return
			}

			ctx := context.WithValue(r.Context(), UserClaimsContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func extractToken(r *http.Request, cookieName string) string {
	// 1. Try extracting from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") {
			return parts[1]
		}
	}

	// 2. Try extracting from Cookie
	if cookie, err := r.Cookie(cookieName); err == nil && cookie.Value != "" {
		return cookie.Value
	}

	return ""
}