package dto

type UserToken struct {
	Role     string `json:"role"`
	UserId   string `json:"userId"`
	Username string `json:"username"`
}

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type GetCurrentUserResponse struct {
	Role     string `json:"role"`
	Id       string `json:"id"`
	Username string `json:"username"`
}
