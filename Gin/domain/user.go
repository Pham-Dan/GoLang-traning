package domain

type CreateUserRequest struct {
	Name     string `form:"name" validate:"required"`
	Email    string `form:"email" validate:"required"`
	Password string `form:"password" validate:"required"`
}
