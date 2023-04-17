package exception

type UserNotFoundError struct {
	Error string
}

func NewUserNotFoundError(error string) UserNotFoundError {
	return UserNotFoundError{
		Error: error,
	}
}
