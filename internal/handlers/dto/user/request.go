package user

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}