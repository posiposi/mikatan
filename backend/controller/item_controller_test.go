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
	"github.com/posiposi/project/backend/presenter"
	"github.com/posiposi/project/backend/usecase/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockItemUsecaseForUserController struct {
	mock.Mock
}

func (m *MockItemUsecaseForUserController) GetAllItems() ([]*domain.Item, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Item), args.Error(1)
}

func (m *MockItemUsecaseForUserController) GetItemByID(itemId string) (*domain.Item, error) {
	args := m.Called(itemId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Item), args.Error(1)
}

func (m *MockItemUsecaseForUserController) CreateItem(req request.CreateItemRequest) (*domain.Item, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Item), args.Error(1)
}

func (m *MockItemUsecaseForUserController) UpdateItem(req request.UpdateItemRequest) (*domain.Item, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Item), args.Error(1)
}

func (m *MockItemUsecaseForUserController) DeleteItem(itemId string) error {
	args := m.Called(itemId)
	return args.Error(0)
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

func TestCreateItem_Success(t *testing.T) {
	e := echo.New()
	e.Validator = &MockValidator{}
	mockUsecase := new(MockItemUsecaseForUserController)
	controller := NewItemController(mockUsecase)

	reqBody := map[string]interface{}{
		"item_name":   "Test Item",
		"stock":       true,
		"description": "Test Description",
	}
	jsonBody, _ := json.Marshal(reqBody)

	userId, _ := domain.NewUserId("f47ac10b-58cc-4372-a567-0e02b2c3d400")
	itemName, _ := domain.NewItemName("Test Item")
	stock, _ := domain.NewStock(true)
	description, _ := domain.NewDescription("Test Description")
	itemId, _ := domain.NewItemId("f47ac10b-58cc-4372-a567-0e02b2c3d401")
	expectedDomainItem, _ := domain.NewItem(itemId, *userId, *itemName, *stock, *description)

	expectedReq := request.CreateItemRequest{
		ItemName:    "Test Item",
		Stock:       true,
		Description: "Test Description",
		UserId:      "f47ac10b-58cc-4372-a567-0e02b2c3d400",
	}
	mockUsecase.On("CreateItem", expectedReq).Return(expectedDomainItem, nil)
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
	mockUsecase := new(MockItemUsecaseForUserController)
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
	mockUsecase := new(MockItemUsecaseForUserController)
	controller := NewItemController(mockUsecase)

	reqBody := map[string]interface{}{
		"item_name":   "",
		"stock":       true,
		"description": "Test Description",
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
	mockUsecase := new(MockItemUsecaseForUserController)
	controller := NewItemController(mockUsecase)

	reqBody := map[string]interface{}{
		"item_name":   "Test Item",
		"stock":       true,
		"description": "Test Description",
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
	mockUsecase := new(MockItemUsecaseForUserController)
	controller := NewItemController(mockUsecase)

	reqBody := map[string]interface{}{
		"item_name":   "Test Item",
		"stock":       true,
		"description": "Test Description",
	}
	jsonBody, _ := json.Marshal(reqBody)
	expectedReq := request.CreateItemRequest{
		ItemName:    "Test Item",
		Stock:       true,
		Description: "Test Description",
		UserId:      "f47ac10b-58cc-4372-a567-0e02b2c3d400",
	}
	mockUsecase.On("CreateItem", expectedReq).Return(nil, errors.New("usecase error"))
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
