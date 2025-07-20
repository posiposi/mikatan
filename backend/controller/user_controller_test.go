package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/usecase/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) SignUp(req request.SignUpRequest) (*domain.User, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserUsecase) Login(req request.LogInRequest) (string, *domain.User, error) {
	args := m.Called(req)
	if args.Get(1) == nil {
		return args.String(0), nil, args.Error(2)
	}
	return args.String(0), args.Get(1).(*domain.User), args.Error(2)
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