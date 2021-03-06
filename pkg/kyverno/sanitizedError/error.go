package sanitizedError

type customError struct {
	message string
}

func (c customError) Error() string {
	return c.message
}

func New(message string) error {
	return customError{message: message}
}

func IsErrorSanitized(err error) bool {
	if _, ok := err.(customError); !ok {
		return false
	}
	return true
}
