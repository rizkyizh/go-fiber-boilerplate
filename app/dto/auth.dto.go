package dto

type RegisterDTO struct {
	Name     string `json:"name"     validate:"required,min=3,max=100"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Age      int    `json:"age"      validate:"required"`
}

type LoginDTO struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type TokenResponseDTO struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
