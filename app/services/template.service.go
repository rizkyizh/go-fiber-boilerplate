package services

import (
	"errors"
	"strconv"

	"github.com/rizkyizh/go-fiber-boilerplate/app/dto"
	"github.com/rizkyizh/go-fiber-boilerplate/app/mappers"
	"github.com/rizkyizh/go-fiber-boilerplate/app/repositories"
	"github.com/rizkyizh/go-fiber-boilerplate/utils"
)

type UserService interface {
	CreateUser(dto *dto.UserDTO) error
	GetAllUsers(query utils.QueryParams) ([]*dto.UserDTO, utils.Meta, error)
	GetUserById(userId string) (*dto.UserDTO, error)
	UpdateUser(userId string, dto *dto.UpdateUserDTO) (*dto.UserDTO, error)
	DeleteUser(userId string) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateUser(dto *dto.UserDTO) error {
	user := mappers.ToUser(dto)

	err := s.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) GetAllUsers(query utils.QueryParams) ([]*dto.UserDTO, utils.Meta, error) {
	page, PerPage := utils.GetPaginationParams(query.Page, query.PerPage)

	users, totalItems, err := s.userRepository.GetUsers(page, PerPage)
	if err != nil {
		return nil, utils.Meta{}, err
	}

	usersDTOs := mappers.UsersToDTOs(users)

	meta := utils.MetaPagination(
		page, PerPage, int(totalItems),
	)

	return usersDTOs, meta, err
}

func (s *userService) GetUserById(id string) (*dto.UserDTO, error) {
	userId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return nil, nil
	}
	user, err := s.userRepository.GetUser(uint(userId))
	if err != nil {
		return nil, err
	}

	userDTO := mappers.ToUserDTO(user)

	return userDTO, err
}

func (s *userService) UpdateUser(id string, dto *dto.UpdateUserDTO) (*dto.UserDTO, error) {
	userId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return nil, nil
	}

	user, err := s.userRepository.GetUser(uint(userId))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("User not found")
	}

	userModel := mappers.UpdateUserDTO_ToUser(dto)

	err = s.userRepository.UpdateUser(uint(userId), userModel)
	if err != nil {
		return nil, err
	}

	updateduser, err := s.userRepository.GetUser(uint(userId))
	if err != nil {
		return nil, err
	}

	userDTO := mappers.ToUserDTO(updateduser)

	return userDTO, nil
}

func (s *userService) DeleteUser(id string) error {
	userId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return err
	}

	user, err := s.userRepository.GetUser(uint(userId))
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	err = s.userRepository.DeleteUser(uint(userId))
	if err != nil {
		return err
	}

	return nil
}
