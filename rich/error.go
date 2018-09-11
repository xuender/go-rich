package rich

// HTTPError 错误
type HTTPError struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message"`
}

func newHTTPError(err error) *HTTPError {
	return &HTTPError{
		Message: err.Error(),
	}
}
