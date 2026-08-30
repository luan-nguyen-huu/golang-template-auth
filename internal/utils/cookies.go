package utils

import (
	"net/http"
	"time"

	"github.com/luan-nguyen-huu/Adam/configs"
)

// SetAuthCookies sets access_token and refresh_token cookies on the HTTP response.
func SetAuthCookies(w http.ResponseWriter, accessToken, refreshToken string, jwtCfg *configs.JWTConfig) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   jwtCfg.Secure,
		Expires:  time.Now().Add(jwtCfg.AccessTokenExpire),
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   jwtCfg.Secure,
		Expires:  time.Now().Add(jwtCfg.RefreshTokenExpire),
		SameSite: http.SameSiteLaxMode,
	})
}

// ClearAuthCookies removes access_token and refresh_token cookies.
func ClearAuthCookies(w http.ResponseWriter, jwtCfg *configs.JWTConfig) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   jwtCfg.Secure,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   jwtCfg.Secure,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
	})
}
