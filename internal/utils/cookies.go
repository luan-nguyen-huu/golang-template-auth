package utils

import (
	"net/http"
	"time"

	"github.com/luan-nguyen-huu/Adam/configs"
)

func SetAuthCookies(w http.ResponseWriter, accessToken, refreshToken string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   configs.Cfg.JWT.Secure,
		Expires:  time.Now().Add(configs.Cfg.JWT.AccessTokenExpire),
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   configs.Cfg.JWT.Secure,
		Expires:  time.Now().Add(configs.Cfg.JWT.RefreshTokenExpire),
		SameSite: http.SameSiteLaxMode,
	})
}
