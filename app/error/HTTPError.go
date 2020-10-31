package error

// HTTPError contain Value
type HTTPError struct {
	HTTPError  string `json:"error"`
	HTTPStatus int    `json:"code"`
}

func (str HTTPError) Error() string {
	return str.HTTPError
}

// NewHTTPError return new object of HTTPError
func NewHTTPError(err string, status int) error {
	return &HTTPError{
		HTTPError:  err,
		HTTPStatus: status,
	}
}
