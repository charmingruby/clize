package application

import "github.com/charmingruby/clize/internal/validation"

const (
	awaiting = "awaiting"
	done     = "done"
)

func Status() map[string]string {
	return map[string]string{
		awaiting: "awaiting",
		done:     "done",
	}
}

func ValidateStatus(sts string) error {
	allStatus := Status()

	isStsValid := allStatus[sts] != ""

	if !isStsValid {
		return &validation.GenericValidationError{
			Message: "Invalid status",
		}
	}

	return nil
}
