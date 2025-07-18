package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/posiposi/project/backend/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) SignUp(user model.User) (model.UserResponse, error) {
	args := m.Called(user)
	return args.Get(0).(model.UserResponse), args.Error(1)
}

func (m *MockUserUsecase) Login(user model.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func TestCheckAuth_Success(t *testing.T) {
	e := echo.New()
	mockUsecase := new(MockUserUsecase)
	controller := NewUserController(mockUsecase)

	req := httptest.NewRequest(http.MethodGet, "/v1/auth/check", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := controller.CheckAuth(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `"authenticated":true`)
}