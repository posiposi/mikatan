package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/posiposi/project/backend/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserByEmail(user *model.User, email string) error {
	args := m.Called(user, email)
	if args.Get(0) != nil {
		*user = args.Get(0).(model.User)
	}
	return args.Error(1)
}

func (m *MockUserRepository) CreateUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) SaveUserToken(userID, token string) error {
	args := m.Called(userID, token)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUserToken(userID string) error {
	args := m.Called(userID)
	return args.Error(0)
}

func TestUserUsecase_SignUpWithAutoLogin(t *testing.T) {
	t.Run("should auto login after successful signup", func(t *testing.T) {
		// Arrange
		mockRepo := new(MockUserRepository)
		usecase := NewUserUsecase(mockRepo)

		signupUser := model.User{
			Email:    "test@example.com",
			Password: "password123",
		}

		// Mock expectations
		mockRepo.On("CreateUser", mock.AnythingOfType("*model.User")).Return(nil)
		mockRepo.On("SaveUserToken", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

		// Act
		response, token, err := usecase.SignUpWithAutoLogin(signupUser)

		// Assert
		require.NoError(t, err)
		assert.NotEmpty(t, response.ID)
		assert.Equal(t, signupUser.Email, response.Email)
		assert.NotEmpty(t, token)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if token save fails during signup", func(t *testing.T) {
		// Arrange
		mockRepo := new(MockUserRepository)
		usecase := NewUserUsecase(mockRepo)

		signupUser := model.User{
			Email:    "test@example.com",
			Password: "password123",
		}

		// Mock expectations
		mockRepo.On("CreateUser", mock.AnythingOfType("*model.User")).Return(nil)
		mockRepo.On("SaveUserToken", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(assert.AnError)

		// Act
		_, _, err := usecase.SignUpWithAutoLogin(signupUser)

		// Assert
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserUsecase_Logout(t *testing.T) {
	t.Run("should delete user token successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(MockUserRepository)
		usecase := NewUserUsecase(mockRepo)

		userID := uuid.New().String()

		// Mock expectations
		mockRepo.On("DeleteUserToken", userID).Return(nil)

		// Act
		err := usecase.Logout(userID)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if token deletion fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(MockUserRepository)
		usecase := NewUserUsecase(mockRepo)

		userID := uuid.New().String()

		// Mock expectations
		mockRepo.On("DeleteUserToken", userID).Return(assert.AnError)

		// Act
		err := usecase.Logout(userID)

		// Assert
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserUsecase_Login_WithTokenSave(t *testing.T) {
	t.Run("should save token to database on successful login", func(t *testing.T) {
		// Arrange
		mockRepo := new(MockUserRepository)
		usecase := NewUserUsecase(mockRepo)

		userID := uuid.New().String()
		email := "test@example.com"
		password := "password123"
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

		storedUser := model.User{
			ID:       userID,
			Email:    email,
			Password: string(hashedPassword),
		}

		loginUser := model.User{
			Email:    email,
			Password: password,
		}

		// Mock expectations
		mockRepo.On("GetUserByEmail", mock.AnythingOfType("*model.User"), email).Return(storedUser, nil)
		mockRepo.On("SaveUserToken", userID, mock.AnythingOfType("string")).Return(nil)

		// Act
		token, err := usecase.Login(loginUser)

		// Assert
		require.NoError(t, err)
		assert.NotEmpty(t, token)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if token save fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(MockUserRepository)
		usecase := NewUserUsecase(mockRepo)

		userID := uuid.New().String()
		email := "test@example.com"
		password := "password123"
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

		storedUser := model.User{
			ID:       userID,
			Email:    email,
			Password: string(hashedPassword),
		}

		loginUser := model.User{
			Email:    email,
			Password: password,
		}

		// Mock expectations
		mockRepo.On("GetUserByEmail", mock.AnythingOfType("*model.User"), email).Return(storedUser, nil)
		mockRepo.On("SaveUserToken", userID, mock.AnythingOfType("string")).Return(assert.AnError)

		// Act
		_, err := usecase.Login(loginUser)

		// Assert
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}