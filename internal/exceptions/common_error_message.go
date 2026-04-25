package exceptions

import "fmt"

// ======================
// Request / Validation
// ======================

const (
	ErrInvalidRequest = "The request data is invalid. Please check your input and try again."
	ErrMissingRequiredField = "Some required information is missing. Please complete all required fields."
	ErrParameterMissing = "A required parameter is missing from the request."
)

// ======================
// System / Server
// ======================

const (
	ErrInternalServer = "Something went wrong on our side. Please try again later."
)

// ======================
// Entity-based templates
// ======================

const (
	ErrFailedToCreateTemplate = "We couldn't create the %s. Please try again."
	ErrFailedToUpdateTemplate = "We couldn't update the %s. Please try again."
	ErrFailedToDeleteTemplate = "We couldn't delete the %s. Please try again."
	ErrNotFoundTemplate       = "The %s you are looking for could not be found."
)

// ======================
// Formatter
// ======================

func FormatErrorMessage(template string, args ...interface{}) string {
	return fmt.Sprintf(template, args...)
}
