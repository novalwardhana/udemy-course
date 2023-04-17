package web

type UserUpdateRequest struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
}
