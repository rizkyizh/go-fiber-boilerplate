package services_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"

	"github.com/rizkyizh/go-fiber-boilerplate/app/dto"
	"github.com/rizkyizh/go-fiber-boilerplate/app/models"
	"github.com/rizkyizh/go-fiber-boilerplate/app/services"
	"github.com/rizkyizh/go-fiber-boilerplate/config"
)

func init() {
	config.AppConfig = config.Config{
		JWT_SECRET:         "test-access-secret",
		JWT_REFRESH_SECRET: "test-refresh-secret",
		JWT_ACCESS_EXPIRY:  "15m",
		JWT_REFRESH_EXPIRY: "168h",
	}
}

// mockAuthRepository is a manual mock of repositories.AuthRepository.
type mockAuthRepository struct {
	user           *models.User
	getUserErr     error
	updateTokenErr error
}

func (m *mockAuthRepository) GetUserByEmail(email string) (*models.User, error) {
	if m.getUserErr != nil {
		return nil, m.getUserErr
	}
	if m.user != nil && m.user.Email == email {
		return m.user, nil
	}
	return nil, errors.New("user not found")
}

func (m *mockAuthRepository) UpdateRefreshToken(userID uint, refreshToken string) error {
	if m.updateTokenErr != nil {
		return m.updateTokenErr
	}
	if m.user != nil {
		m.user.RefreshToken = refreshToken
	}
	return nil
}

func (m *mockAuthRepository) ClearRefreshToken(userID uint) error {
	if m.user != nil {
		m.user.RefreshToken = ""
	}
	return nil
}

// mockUserRepositoryForAuth implements repositories.UserRepository for auth tests.
type mockUserRepositoryForAuth struct {
	createErr error
	created   []*models.User
}

func (m *mockUserRepositoryForAuth) CreateUser(user *models.User) error {
	if m.createErr != nil {
		return m.createErr
	}
	m.created = append(m.created, user)
	return nil
}

func (m *mockUserRepositoryForAuth) GetUsers(page, perPage int) ([]*models.User, int64, error) {
	return nil, 0, nil
}

func (m *mockUserRepositoryForAuth) GetUser(userID uint) (*models.User, error) {
	return nil, nil
}

func (m *mockUserRepositoryForAuth) UpdateUser(userID uint, user *models.User) error {
	return nil
}

func (m *mockUserRepositoryForAuth) DeleteUser(userID uint) error {
	return nil
}

func newTestAuthService(authRepo *mockAuthRepository, userRepo *mockUserRepositoryForAuth) services.AuthService {
	return services.NewAuthService(authRepo, userRepo)
}

func TestAuthRegister_Success(t *testing.T) {
	authRepo := &mockAuthRepository{getUserErr: errors.New("not found")}
	userRepo := &mockUserRepositoryForAuth{}
	svc := newTestAuthService(authRepo, userRepo)

	err := svc.Register(&dto.RegisterDTO{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "Password1!",
		Age:      25,
	})

	assert.NoError(t, err)
	assert.Len(t, userRepo.created, 1)
	assert.Equal(t, models.RoleUser, userRepo.created[0].Role)
}

func TestAuthRegister_DuplicateEmail(t *testing.T) {
	existing := &models.User{ID: 1, Email: "alice@example.com"}
	authRepo := &mockAuthRepository{user: existing}
	userRepo := &mockUserRepositoryForAuth{}
	svc := newTestAuthService(authRepo, userRepo)

	err := svc.Register(&dto.RegisterDTO{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "Password1!",
		Age:      25,
	})

	assert.EqualError(t, err, "email already registered")
}

func TestAuthLogin_Success(t *testing.T) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte("Password1!"), bcrypt.DefaultCost)
	existing := &models.User{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: string(hashed),
		Role:     models.RoleUser,
	}
	authRepo := &mockAuthRepository{user: existing}
	userRepo := &mockUserRepositoryForAuth{}
	svc := newTestAuthService(authRepo, userRepo)

	tokens, err := svc.Login(&dto.LoginDTO{
		Email:    "alice@example.com",
		Password: "Password1!",
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, tokens.AccessToken)
	assert.NotEmpty(t, tokens.RefreshToken)
}

func TestAuthLogin_WrongPassword(t *testing.T) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte("Password1!"), bcrypt.DefaultCost)
	existing := &models.User{
		ID:       1,
		Email:    "alice@example.com",
		Password: string(hashed),
		Role:     models.RoleUser,
	}
	authRepo := &mockAuthRepository{user: existing}
	userRepo := &mockUserRepositoryForAuth{}
	svc := newTestAuthService(authRepo, userRepo)

	_, err := svc.Login(&dto.LoginDTO{
		Email:    "alice@example.com",
		Password: "WrongPassword",
	})

	assert.EqualError(t, err, "invalid email or password")
}

func TestAuthLogin_UserNotFound(t *testing.T) {
	authRepo := &mockAuthRepository{getUserErr: errors.New("not found")}
	userRepo := &mockUserRepositoryForAuth{}
	svc := newTestAuthService(authRepo, userRepo)

	_, err := svc.Login(&dto.LoginDTO{
		Email:    "nobody@example.com",
		Password: "Password1!",
	})

	assert.EqualError(t, err, "invalid email or password")
}

func TestAuthLogout_Success(t *testing.T) {
	existing := &models.User{ID: 1, RefreshToken: "sometoken"}
	authRepo := &mockAuthRepository{user: existing}
	userRepo := &mockUserRepositoryForAuth{}
	svc := newTestAuthService(authRepo, userRepo)

	err := svc.Logout(1)

	assert.NoError(t, err)
	assert.Empty(t, existing.RefreshToken)
}
