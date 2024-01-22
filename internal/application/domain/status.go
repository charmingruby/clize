package domain

import "github.com/charmingruby/clize/pkg/errors"

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
		return &errors.GenericValidationError{
			Message: "Invalid status",
		}
	}

	return nil
}
