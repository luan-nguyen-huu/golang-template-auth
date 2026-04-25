package user

type RegisterUserResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

type LoginUserResponse struct {
}

type GetMeResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}