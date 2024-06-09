package mappers

import (
	"github.com/rizkyizh/go-fiber-boilerplate/app/dto"
	"github.com/rizkyizh/go-fiber-boilerplate/app/models"
)

func ToUserDTO(user models.User) dto.UserDTO {
	return dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
	}
}

func ToUser(userDTO dto.UserDTO) models.User {
	return models.User{
		Name:  userDTO.Name,
		Email: userDTO.Email,
		Age:   userDTO.Age,
	}
}

func UpdateUserDTO_ToUser(updateUserDTO dto.UpdateUserDTO) models.User {
	return models.User{
		Name:  updateUserDTO.Name,
		Email: updateUserDTO.Email,
		Age:   updateUserDTO.Age,
	}
}
