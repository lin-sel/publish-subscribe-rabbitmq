package error

// ValidationError contain Value
type ValidationError struct {
	HTTPError string `json:"error"`
}

func (str ValidationError) Error() string {
	return str.HTTPError
}

// NewValidationError return new object of ValidationError
func NewValidationError(err string) error {
	return &ValidationError{
		HTTPError: err,
	}
}
