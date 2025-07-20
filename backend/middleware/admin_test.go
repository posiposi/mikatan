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

func (m *MockUserRepository) GetUserByID(userID string) (*domain.User, error) {
	args := m.Called(userID)
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

	c.Set("user_id", "admin-user-id")

	userId, _ := domain.NewUserId("admin-user-id")
	email, _ := domain.NewEmail("admin@example.com")
	password, _ := domain.NewPassword("password123")
	role, _ := domain.NewRole("ADMINISTRATOR")
	adminUser, _ := domain.NewUserWithRole(userId, "Admin User", email, password, role)

	mockRepo := new(MockUserRepository)
	mockRepo.On("GetUserByID", "admin-user-id").Return(adminUser, nil)

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

	c.Set("user_id", "regular-user-id")

	userId, _ := domain.NewUserId("regular-user-id")
	email, _ := domain.NewEmail("user@example.com")
	password, _ := domain.NewPassword("password123")
	role, _ := domain.NewRole("USER")
	regularUser, _ := domain.NewUserWithRole(userId, "Regular User", email, password, role)

	mockRepo := new(MockUserRepository)
	mockRepo.On("GetUserByID", "regular-user-id").Return(regularUser, nil)

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
