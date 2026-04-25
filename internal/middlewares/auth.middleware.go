package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/luan-nguyen-huu/Adam/internal/exceptions/auth"
	"github.com/luan-nguyen-huu/Adam/internal/utils"
	"github.com/luan-nguyen-huu/Adam/internal/utils/jwt"
)

type contextKey string

const UserClaimsContextKey contextKey = "adam_user_claims"

func AuthMiddlewareByHeader(tokenMaker jwt.JWTMakerInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.WriteErrorResponse(w, http.StatusUnauthorized, auth.ErrMissingAuthToken)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				utils.WriteErrorResponse(w, http.StatusUnauthorized, auth.ErrInvalidAuthHeader)
				return
			}
			
			tokenString := parts[1]
			claims, err := tokenMaker.VerifyAccessToken(tokenString)
			if err != nil {
				utils.WriteErrorResponse(w, http.StatusUnauthorized, auth.ErrInvalidAuthToken)
				return
			}

			ctx := context.WithValue(r.Context(), UserClaimsContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AuthMiddlewareByCookie(tokenMaker jwt.JWTMakerInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("access_token")
			if err != nil {
				utils.WriteErrorResponse(w, http.StatusUnauthorized, auth.ErrMissingAuthToken)
				return
			}

			tokenString := cookie.Value
			claims, err := tokenMaker.VerifyAccessToken(tokenString)
			if err != nil {
				utils.WriteErrorResponse(w, http.StatusUnauthorized, auth.ErrInvalidAuthToken)
				return
			}

			ctx := context.WithValue(r.Context(), UserClaimsContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RefreshTokenMiddleware(tokenMaker jwt.JWTMakerInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("refresh_token")
			if err != nil {
				utils.WriteErrorResponse(w, http.StatusUnauthorized, auth.ErrMissingAuthToken)
				return
			}

			tokenString := cookie.Value
			claims, err := tokenMaker.VerifyRefreshToken(tokenString)
			if err != nil {
				utils.WriteErrorResponse(w, http.StatusUnauthorized, auth.ErrInvalidAuthToken)
				return
			}

			ctx := context.WithValue(r.Context(), UserClaimsContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}