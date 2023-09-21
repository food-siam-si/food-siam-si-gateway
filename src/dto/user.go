package dto

type UserType string

const (
	Customer UserType = "Customer"
	Owner    UserType = "Owner"
)

type UserToken struct {
	Id   uint32   `json:"id"`
	Type UserType `json:"type"`
	Name string   `json:"username"`
}

type RestUserData struct {
	Id       uint32   `json:"userId"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	UserType UserType `json:"userType"`
}

type FetchUserResponse struct {
	Data    RestUserData `json:"data"`
	Message string       `json:"message"`
}

type CreateUserRequest struct {
	Email    string   `json:"email" validate:"required,email"`
	Username string   `json:"username" validate:"required"`
	Password string   `json:"password" validate:"required,password"`
	UserType UserType `json:"userType" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type GetCurrentUserResponse struct {
	Role     string `json:"role"`
	Id       string `json:"id"`
	Username string `json:"username"`
}
