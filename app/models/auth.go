package models

// LoginDTO defined the /login payload
type LoginDTO struct {
	Username string `json:"username" validate:"required,min=5,max=16"`
	Password string `json:"password" validate:"password"`
}

// SignupDTO defined the /login payload
type SignupDTO struct {
	FirstName string `json:"firstName" validate:"required,min=1"`
	LastName  string `json:"lastName" validate:"required,min=1"`
	Username  string `json:"username" validate:"required,min=5,max=16"`
	Password  string `json:"password" validate:"password"`
}
