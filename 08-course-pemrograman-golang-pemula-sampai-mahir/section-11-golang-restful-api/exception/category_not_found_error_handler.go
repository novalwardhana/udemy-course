package exception

type CategoryNotFoundError struct {
	Error string
}

func NewCategoryNotFoundError(error string) CategoryNotFoundError {
	return CategoryNotFoundError{Error: error}
}
