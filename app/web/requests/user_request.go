package requests

type UserCreateRequest struct {
    FullName string `json:"fullname" validate:"required,min=3,max=50"` 
    Email string `json:"email" validate:"required,email"`
    Username string `json:"username" validate:"required,min=6,max=100"`
    Password string `json:"password" validate:"required,min=6,max=100"`
}

type UserLoginRequest struct {
    Username string `json:"username" validate:"required,min=6,max=100"`
    Password string `json:"password" validate:"required,min=8,max=100"`
}