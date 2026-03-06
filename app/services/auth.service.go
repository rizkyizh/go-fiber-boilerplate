package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/rizkyizh/go-fiber-boilerplate/app/dto"
	"github.com/rizkyizh/go-fiber-boilerplate/app/mappers"
	"github.com/rizkyizh/go-fiber-boilerplate/app/repositories"
	"github.com/rizkyizh/go-fiber-boilerplate/config"
	"github.com/rizkyizh/go-fiber-boilerplate/utils"
)

type AuthService interface {
	Register(dto *dto.RegisterDTO) error
	Login(dto *dto.LoginDTO) (*dto.TokenResponseDTO, error)
	RefreshToken(refreshToken string) (*dto.TokenResponseDTO, error)
	Logout(userID uint) error
}

type authService struct {
	authRepo repositories.AuthRepository
	userRepo repositories.UserRepository
}

func NewAuthService(authRepo repositories.AuthRepository, userRepo repositories.UserRepository) AuthService {
	return &authService{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (s *authService) Register(registerDTO *dto.RegisterDTO) error {
	existing, _ := s.authRepo.GetUserByEmail(registerDTO.Email)
	if existing != nil {
		return errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(registerDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := mappers.RegisterDTO_ToUserModel(registerDTO)
	user.Password = string(hashed)

	return s.userRepo.CreateUser(user)
}

func (s *authService) Login(loginDTO *dto.LoginDTO) (*dto.TokenResponseDTO, error) {
	user, err := s.authRepo.GetUserByEmail(loginDTO.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	cfg := config.AppConfig

	accessToken, err := utils.GenerateAccessToken(user.ID, user.Email, user.Role, cfg.JWT_SECRET, cfg.JWT_ACCESS_EXPIRY)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Email, user.Role, cfg.JWT_REFRESH_SECRET, cfg.JWT_REFRESH_EXPIRY)
	if err != nil {
		return nil, err
	}

	if err := s.authRepo.UpdateRefreshToken(user.ID, refreshToken); err != nil {
		return nil, err
	}

	return &dto.TokenResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) RefreshToken(refreshToken string) (*dto.TokenResponseDTO, error) {
	cfg := config.AppConfig

	claims, err := utils.ValidateToken(refreshToken, cfg.JWT_REFRESH_SECRET)
	if err != nil {
		return nil, errors.New("invalid or expired refresh token")
	}

	user, err := s.authRepo.GetUserByEmail(claims.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.RefreshToken != refreshToken {
		return nil, errors.New("refresh token mismatch")
	}

	newAccessToken, err := utils.GenerateAccessToken(user.ID, user.Email, user.Role, cfg.JWT_SECRET, cfg.JWT_ACCESS_EXPIRY)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := utils.GenerateRefreshToken(user.ID, user.Email, user.Role, cfg.JWT_REFRESH_SECRET, cfg.JWT_REFRESH_EXPIRY)
	if err != nil {
		return nil, err
	}

	if err := s.authRepo.UpdateRefreshToken(user.ID, newRefreshToken); err != nil {
		return nil, err
	}

	return &dto.TokenResponseDTO{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *authService) Logout(userID uint) error {
	return s.authRepo.ClearRefreshToken(userID)
}

