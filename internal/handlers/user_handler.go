package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/luan-nguyen-huu/Adam/configs"
	"github.com/luan-nguyen-huu/Adam/internal/domain"
	user_dto "github.com/luan-nguyen-huu/Adam/internal/handlers/dto/user"
	"github.com/luan-nguyen-huu/Adam/internal/middlewares"
	"github.com/luan-nguyen-huu/Adam/internal/utils"
	"github.com/luan-nguyen-huu/Adam/internal/utils/jwt"
)

type UserHandler struct {
	userService domain.UserService
	jwtCfg      *configs.JWTConfig
}

func NewUserHandler(userService domain.UserService, jwtCfg *configs.JWTConfig) *UserHandler {
	return &UserHandler{
		userService: userService,
		jwtCfg:      jwtCfg,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req user_dto.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload format")
		return
	}

	if err := utils.ValidateStruct(&req); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.userService.Register(r.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		h.handleError(w, err)
		return
	}

	utils.SetAuthCookies(w, tokens.AccessToken, tokens.RefreshToken, h.jwtCfg)

	resp := user_dto.AuthResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
	utils.WriteSuccessResponse(w, http.StatusCreated, "User registered successfully", resp)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req user_dto.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload format")
		return
	}

	if err := utils.ValidateStruct(&req); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.userService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		h.handleError(w, err)
		return
	}

	utils.SetAuthCookies(w, tokens.AccessToken, tokens.RefreshToken, h.jwtCfg)

	resp := user_dto.AuthResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
	utils.WriteSuccessResponse(w, http.StatusOK, "User logged in successfully", resp)
}

func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middlewares.UserClaimsContextKey).(*jwt.UserClaims)
	if !ok || claims == nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, domain.ErrInvalidToken.Error())
		return
	}

	user, err := h.userService.GetMe(r.Context(), claims.UserID)
	if err != nil {
		h.handleError(w, err)
		return
	}

	resp := user_dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	utils.WriteSuccessResponse(w, http.StatusOK, "User profile retrieved successfully", resp)
}

func (h *UserHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middlewares.UserClaimsContextKey).(*jwt.UserClaims)
	if !ok || claims == nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, domain.ErrInvalidToken.Error())
		return
	}

	tokens, err := h.userService.RefreshToken(r.Context(), claims.UserID)
	if err != nil {
		h.handleError(w, err)
		return
	}

	utils.SetAuthCookies(w, tokens.AccessToken, tokens.RefreshToken, h.jwtCfg)

	resp := user_dto.AuthResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
	utils.WriteSuccessResponse(w, http.StatusOK, "Token refreshed successfully", resp)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	utils.ClearAuthCookies(w, h.jwtCfg)
	utils.WriteSuccessResponse(w, http.StatusOK, "Logged out successfully", nil)
}

func (h *UserHandler) handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidCredentials):
		utils.WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
	case errors.Is(err, domain.ErrUserAlreadyExists):
		utils.WriteErrorResponse(w, http.StatusConflict, err.Error())
	case errors.Is(err, domain.ErrUserNotFound):
		utils.WriteErrorResponse(w, http.StatusNotFound, err.Error())
	case errors.Is(err, domain.ErrInvalidToken), errors.Is(err, domain.ErrTokenExpired):
		utils.WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
	case errors.Is(err, domain.ErrBadRequest), errors.Is(err, domain.ErrValidationFailed):
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
	default:
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "An unexpected error occurred. Please try again later.")
	}
}