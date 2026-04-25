package auth

const (
	ErrFailedToLoginTemplate = "failed to login because: %s"
	IncorrectEmailOrPassword  = "Incorrect email or password. Please try again."
)

const (
	ErrMissingAuthToken   = "missing authorization token"
	ErrInvalidAuthHeader  = "invalid authorization header format"
	ErrInvalidAuthToken   = "invalid or expired authorization token"
)