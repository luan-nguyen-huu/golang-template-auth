package handlers

import (
	"encoding/json"
	"github.com/luan-nguyen-huu/Adam/internal/entities"
	"github.com/luan-nguyen-huu/Adam/internal/middlewares"
	user_resp "github.com/luan-nguyen-huu/Adam/internal/handlers/dto/user"
	"github.com/luan-nguyen-huu/Adam/internal/utils"
	"github.com/luan-nguyen-huu/Adam/internal/utils/jwt"
	"github.com/luan-nguyen-huu/Adam/internal/exceptions"
	exceptions_auth "github.com/luan-nguyen-huu/Adam/internal/exceptions/auth"
	"net/http"
)

type UserHandler struct {
	userService entities.UserServiceInterface
}

func NewUserHandler(userService entities.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req user_resp.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, exceptions.ErrInvalidRequest)
		return
	}

	accessToken, refreshToken, err := h.userService.RegisterUser(req.Name, req.Password, req.Email);
	if  err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, exceptions.FormatErrorMessage(exceptions.ErrFailedToCreateTemplate, "user"))
		return
	}
	rep := user_resp.RegisterUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	utils.WriteSuccessResponse(w, http.StatusCreated, "User registered successfully", rep)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req user_resp.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, exceptions.ErrInvalidRequest)
		return
	}

	accessToken, refreshToken, err := h.userService.LoginUser(req.Email, req.Password);
	if  err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, exceptions_auth.IncorrectEmailOrPassword)
		return
	}
	utils.SetAuthCookies(w, accessToken, refreshToken)

	rep := user_resp.LoginUserResponse{}
	utils.WriteSuccessResponse(w, http.StatusOK, "User logged in successfully", rep)
}

func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middlewares.UserClaimsContextKey).(*jwt.UserClaims)
	if !ok {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, exceptions_auth.ErrInvalidAuthToken)
		return
	}

	user, err := h.userService.GetMe(claims.UserID)
	if err != nil {
		utils.WriteErrorResponse(
			w,
			http.StatusInternalServerError,
			exceptions.FormatErrorMessage(exceptions.ErrNotFoundTemplate, "user"),
		)
		return
	}

	resp := user_resp.GetMeResponse{
		Name:  user.Name,
		Email: user.Email,
	}
	utils.WriteSuccessResponse(w, http.StatusOK, "User fetched successfully", resp)
}

func (h *UserHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middlewares.UserClaimsContextKey).(*jwt.UserClaims)
	if !ok {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, exceptions_auth.ErrInvalidAuthToken)
		return
	}
	access_token, refresh_token, err := h.userService.RefreshToken(claims.UserID)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, exceptions.FormatErrorMessage(exceptions.ErrFailedToCreateTemplate, "token"))
		return
	}
	utils.SetAuthCookies(w, access_token, refresh_token)

	utils.WriteSuccessResponse(w, http.StatusOK, "Token refreshed successfully", nil)
}