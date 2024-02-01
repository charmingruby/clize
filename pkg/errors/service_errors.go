package errors

type UniqueValueViolationError struct {
	Entity  string `json:"entity"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (v *UniqueValueViolationError) Error() string {
	return v.Message
}

type ResourceNotFoundError struct {
	Entity  string `json:"entity"`
	Message string `json:"message"`
}

func (nf *ResourceNotFoundError) Error() string {
	return nf.Message
}
