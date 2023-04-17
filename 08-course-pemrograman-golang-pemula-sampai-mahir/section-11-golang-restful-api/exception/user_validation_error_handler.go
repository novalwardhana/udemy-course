package exception

type UserValidationError struct {
	Error string
}

func NewUserValidationError(error string) UserValidationError {
	return UserValidationError{
		Error: error,
	}
}
