package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/posiposi/project/backend/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserById(userId *domain.UserId) (*domain.User, error) {
	args := m.Called(userId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestAdminMiddleware_WithAdminUser_ShouldProceed(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/admin/items", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	adminUserID := "f47ac10b-58cc-4372-a567-0e02b2c3d500"
	c.Set("user_id", adminUserID)

	userId, _ := domain.NewUserId(adminUserID)
	email, _ := domain.NewEmail("admin@example.com")
	password, _ := domain.NewPassword("password123")
	role, _ := domain.NewRole("ADMINISTRATOR")
	adminUser, _ := domain.NewUserWithRole(userId, "Admin User", email, password, role)

	mockRepo := new(MockUserRepository)
	mockRepo.On("GetUserById", userId).Return(adminUser, nil)

	called := false
	nextHandler := func(c echo.Context) error {
		called = true
		return c.String(http.StatusOK, "success")
	}

	middleware := AdminMiddleware(mockRepo)
	err := middleware(nextHandler)(c)

	assert.NoError(t, err)
	assert.True(t, called)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestAdminMiddleware_WithNonAdminUser_ShouldReturnForbidden(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/admin/items", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	regularUserID := "f47ac10b-58cc-4372-a567-0e02b2c3d501"
	c.Set("user_id", regularUserID)

	userId, _ := domain.NewUserId(regularUserID)
	email, _ := domain.NewEmail("user@example.com")
	password, _ := domain.NewPassword("password123")
	role, _ := domain.NewRole("USER")
	regularUser, _ := domain.NewUserWithRole(userId, "Regular User", email, password, role)

	mockRepo := new(MockUserRepository)
	mockRepo.On("GetUserById", userId).Return(regularUser, nil)

	called := false
	nextHandler := func(c echo.Context) error {
		called = true
		return c.String(http.StatusOK, "success")
	}

	middleware := AdminMiddleware(mockRepo)
	err := middleware(nextHandler)(c)

	assert.NoError(t, err)
	assert.False(t, called)
	assert.Equal(t, http.StatusForbidden, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestAdminMiddleware_WithoutUserID_ShouldReturnUnauthorized(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/admin/items", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo := new(MockUserRepository)

	called := false
	nextHandler := func(c echo.Context) error {
		called = true
		return c.String(http.StatusOK, "success")
	}

	middleware := AdminMiddleware(mockRepo)
	err := middleware(nextHandler)(c)

	assert.NoError(t, err)
	assert.False(t, called)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	mockRepo.AssertExpectations(t)
}
