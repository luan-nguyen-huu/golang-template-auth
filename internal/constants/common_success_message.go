package constants

import "fmt"

// ======================
// Success message templates
// ======================

const (
	SuccessCreatedTemplate   = "%s has been created successfully."
	SuccessRetrievedTemplate = "%s has been retrieved successfully."
	SuccessUpdatedTemplate   = "%s has been updated successfully."
	SuccessDeletedTemplate   = "%s has been deleted successfully."
)

// ======================
// Formatter
// ======================

func FormatSuccessMessage(template string, args ...interface{}) string {
	return fmt.Sprintf(template, args...)
}
