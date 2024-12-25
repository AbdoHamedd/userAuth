package user

type UserSignupResponse struct {
	Name  string `json:"name"`
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserLoginResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
	Role  string `json:"role"`
}
