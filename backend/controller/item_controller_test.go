package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/model"
	"github.com/posiposi/project/backend/presenter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockItemUsecase struct {
	mock.Mock
}

type MockValidator struct {
	shouldFail bool
}

func (m *MockValidator) Validate(i interface{}) error {
	if m.shouldFail {
		return errors.New("validation error")
	}
	return nil
}

func (m *MockItemUsecase) GetAllItems() ([]*domain.Item, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Item), args.Error(1)
}

func (m *MockItemUsecase) CreateItem(item model.Item, userId string) (*domain.Item, error) {
	args := m.Called(item, userId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Item), args.Error(1)
}

func TestCreateItem_Success(t *testing.T) {
	e := echo.New()
	e.Validator = &MockValidator{}
	mockUsecase := new(MockItemUsecase)
	controller := NewItemController(mockUsecase)

	reqBody := model.Item{
		ItemName:    "Test Item",
		Stock:       true,
		Description: "Test Description",
	}
	jsonBody, _ := json.Marshal(reqBody)

	userId, _ := domain.NewUserId("f47ac10b-58cc-4372-a567-0e02b2c3d400")
	itemName, _ := domain.NewItemName("Test Item")
	stock, _ := domain.NewStock(true)
	description, _ := domain.NewDescription("Test Description")
	expectedDomainItem, _ := domain.NewItem(nil, *userId, *itemName, *stock, *description)

	mockUsecase.On("CreateItem", reqBody, "f47ac10b-58cc-4372-a567-0e02b2c3d400").Return(expectedDomainItem, nil)
	req := httptest.NewRequest(http.MethodPost, "/v1/items", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user_id", "f47ac10b-58cc-4372-a567-0e02b2c3d400")

	err := controller.CreateItem(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var response presenter.ItemResponseJSON
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedDomainItem.ItemId(), response.ItemId)
	assert.Equal(t, expectedDomainItem.ItemName(), response.ItemName)
	assert.Equal(t, expectedDomainItem.Stock(), response.Stock)
	assert.Equal(t, expectedDomainItem.Description(), response.Description)
	mockUsecase.AssertExpectations(t)
}

func TestCreateItem_InvalidJSON(t *testing.T) {
	e := echo.New()
	mockUsecase := new(MockItemUsecase)
	controller := NewItemController(mockUsecase)

	req := httptest.NewRequest(http.MethodPost, "/v1/items", bytes.NewReader([]byte("invalid json")))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user_id", "f47ac10b-58cc-4372-a567-0e02b2c3d400")
	err := controller.CreateItem(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockUsecase.AssertNotCalled(t, "CreateItem")
}

func TestCreateItem_ValidationError(t *testing.T) {
	e := echo.New()
	e.Validator = &MockValidator{shouldFail: true}
	mockUsecase := new(MockItemUsecase)
	controller := NewItemController(mockUsecase)

	reqBody := model.Item{
		ItemName:    "",
		Stock:       true,
		Description: "Test Description",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/v1/items", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user_id", "f47ac10b-58cc-4372-a567-0e02b2c3d400")

	err := controller.CreateItem(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockUsecase.AssertNotCalled(t, "CreateItem")
}

func TestCreateItem_MissingUserId(t *testing.T) {
	e := echo.New()
	e.Validator = &MockValidator{}
	mockUsecase := new(MockItemUsecase)
	controller := NewItemController(mockUsecase)

	reqBody := model.Item{
		ItemName:    "Test Item",
		Stock:       true,
		Description: "Test Description",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/v1/items", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := controller.CreateItem(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	mockUsecase.AssertNotCalled(t, "CreateItem")
}

func TestCreateItem_UsecaseError(t *testing.T) {
	e := echo.New()
	e.Validator = &MockValidator{}
	mockUsecase := new(MockItemUsecase)
	controller := NewItemController(mockUsecase)

	reqBody := model.Item{
		ItemName:    "Test Item",
		Stock:       true,
		Description: "Test Description",
	}
	jsonBody, _ := json.Marshal(reqBody)
	mockUsecase.On("CreateItem", reqBody, "f47ac10b-58cc-4372-a567-0e02b2c3d400").Return(nil, errors.New("usecase error"))
	req := httptest.NewRequest(http.MethodPost, "/v1/items", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user_id", "f47ac10b-58cc-4372-a567-0e02b2c3d400")
	err := controller.CreateItem(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "usecase error")
	mockUsecase.AssertExpectations(t)
}
