package dto

type UserDTO struct {
	ID    uint   `json:"id,omitempty"`
	Name  string `json:"name"         validate:"required,min=3,max=32"`
	Email string `json:"email"        validate:"required,email"`
	Age   int    `json:"age"          validate:"required"`
}

type UpdateUserDTO struct {
	Name  string `json:"name,omitempty"  validate:"omitempty,min=3,max=32"`
	Email string `json:"email,omitempty" validate:"omitempty,email"`
	Age   int    `json:"age,omitempty"   validate:"omitempty"`
}
