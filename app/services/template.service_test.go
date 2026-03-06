package services_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rizkyizh/go-fiber-boilerplate/app/dto"
	"github.com/rizkyizh/go-fiber-boilerplate/app/models"
	"github.com/rizkyizh/go-fiber-boilerplate/app/services"
	"github.com/rizkyizh/go-fiber-boilerplate/utils"
)

// mockUserRepository is a manual mock of repositories.UserRepository.
type mockUserRepository struct {
	users     []*models.User
	createErr error
	getErr    error
	updateErr error
	deleteErr error
}

func (m *mockUserRepository) CreateUser(user *models.User) error {
	if m.createErr != nil {
		return m.createErr
	}
	user.ID = uint(len(m.users) + 1)
	m.users = append(m.users, user)
	return nil
}

func (m *mockUserRepository) GetUsers(page, perPage int) ([]*models.User, int64, error) {
	if m.getErr != nil {
		return nil, 0, m.getErr
	}
	return m.users, int64(len(m.users)), nil
}

func (m *mockUserRepository) GetUser(userID uint) (*models.User, error) {
	if m.getErr != nil {
		return nil, m.getErr
	}
	for _, u := range m.users {
		if u.ID == userID {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m *mockUserRepository) UpdateUser(userID uint, user *models.User) error {
	if m.updateErr != nil {
		return m.updateErr
	}
	for _, u := range m.users {
		if u.ID == userID {
			if user.Name != "" {
				u.Name = user.Name
			}
			if user.Email != "" {
				u.Email = user.Email
			}
			return nil
		}
	}
	return errors.New("user not found")
}

func (m *mockUserRepository) DeleteUser(userID uint) error {
	if m.deleteErr != nil {
		return m.deleteErr
	}
	for i, u := range m.users {
		if u.ID == userID {
			m.users = append(m.users[:i], m.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func newUserService(repo *mockUserRepository) services.UserService {
	return services.NewUserService(repo)
}

func TestCreateUser_Success(t *testing.T) {
	repo := &mockUserRepository{}
	svc := newUserService(repo)

	err := svc.CreateUser(&dto.CreateUserDTO{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "Password1!",
		Age:      25,
	})

	assert.NoError(t, err)
	assert.Len(t, repo.users, 1)
	assert.Equal(t, "Alice", repo.users[0].Name)
}

func TestCreateUser_RepositoryError(t *testing.T) {
	repo := &mockUserRepository{createErr: errors.New("db error")}
	svc := newUserService(repo)

	err := svc.CreateUser(&dto.CreateUserDTO{
		Name:     "Bob",
		Email:    "bob@example.com",
		Password: "Password1!",
		Age:      30,
	})

	assert.Error(t, err)
}

func TestGetAllUsers_Success(t *testing.T) {
	repo := &mockUserRepository{
		users: []*models.User{
			{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
			{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30},
		},
	}
	svc := newUserService(repo)

	users, meta, err := svc.GetAllUsers(utils.QueryParams{Page: "1", PerPage: "10"})

	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, 1, meta.CurrentPage)
	assert.Equal(t, 2, meta.TotalData)
}

func TestGetUserById_Success(t *testing.T) {
	repo := &mockUserRepository{
		users: []*models.User{
			{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
		},
	}
	svc := newUserService(repo)

	user, err := svc.GetUserById("1")

	assert.NoError(t, err)
	assert.Equal(t, "Alice", user.Name)
}

func TestGetUserById_NotFound(t *testing.T) {
	repo := &mockUserRepository{}
	svc := newUserService(repo)

	_, err := svc.GetUserById("99")

	assert.Error(t, err)
}

func TestGetUserById_InvalidID(t *testing.T) {
	repo := &mockUserRepository{}
	svc := newUserService(repo)

	_, err := svc.GetUserById("not-a-number")

	assert.Error(t, err)
}

func TestUpdateUser_Success(t *testing.T) {
	repo := &mockUserRepository{
		users: []*models.User{
			{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
		},
	}
	svc := newUserService(repo)

	updated, err := svc.UpdateUser("1", &dto.UpdateUserDTO{Name: "Alice Updated"})

	assert.NoError(t, err)
	assert.Equal(t, "Alice Updated", updated.Name)
}

func TestDeleteUser_Success(t *testing.T) {
	repo := &mockUserRepository{
		users: []*models.User{
			{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
		},
	}
	svc := newUserService(repo)

	err := svc.DeleteUser("1")

	assert.NoError(t, err)
	assert.Len(t, repo.users, 0)
}

func TestDeleteUser_NotFound(t *testing.T) {
	repo := &mockUserRepository{}
	svc := newUserService(repo)

	err := svc.DeleteUser("99")

	assert.Error(t, err)
}
