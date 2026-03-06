package dto

type UserDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"  validate:"required,min=3,max=32"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age"   validate:"required"`
	Role  string `json:"role"`
}

type CreateUserDTO struct {
	Name     string `json:"name"     validate:"required,min=3,max=32"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Age      int    `json:"age"      validate:"required"`
	Role     string `json:"role"     validate:"omitempty,oneof=admin user"`
}

type UpdateUserDTO struct {
	Name  string `json:"name"  validate:"omitempty,min=3,max=32"`
	Email string `json:"email" validate:"omitempty,email"`
	Age   int    `json:"age"   validate:"omitempty"`
}
