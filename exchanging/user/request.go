package user

type SignupRequest struct {
	Email    string `json:"email" validate:"required , email" `
	Password string `json:"password" validate:"required , min=6"`
	UserName string `json:"user_name" validate:"required"`
	Role     string `json:"role" validate:"eq=Student"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" `
	Password string `json:"password" validate:"required,min=6"`
}

type LogoutRequest struct {
	Id int `json:"id"`
}
