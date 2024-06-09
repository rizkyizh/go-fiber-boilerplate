package mappers

import (
	"github.com/rizkyizh/go-fiber-boilerplate/app/dto"
	"github.com/rizkyizh/go-fiber-boilerplate/app/models"
)

func UserModel_ToUserDTO(user *models.User) *dto.UserDTO {
	return &dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
	}
}

func UserDTO_ToUserModel(userDTO *dto.UserDTO) *models.User {
	return &models.User{
		Name:  userDTO.Name,
		Email: userDTO.Email,
		Age:   userDTO.Age,
	}
}

func UpdateUserDTO_ToUserModel(updateUserDTO *dto.UpdateUserDTO) *models.User {
	return &models.User{
		Name:  updateUserDTO.Name,
		Email: updateUserDTO.Email,
		Age:   updateUserDTO.Age,
	}
}

func UsersModel_ToUsersDTOs(users []*models.User) []*dto.UserDTO {
	dtos := make([]*dto.UserDTO, len(users))
	for i, user := range users {
		dtos[i] = &dto.UserDTO{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Age:   user.Age,
		}
	}
	return dtos
}

func CreateUserDTO_ToUserModel(createUserDTO *dto.CreateUserDTO) *models.User {
	return &models.User{
		Name:  createUserDTO.Name,
		Email: createUserDTO.Email,
		Age:   createUserDTO.Age,
	}
}
