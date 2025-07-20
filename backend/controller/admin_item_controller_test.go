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
	"github.com/posiposi/project/backend/usecase/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockItemUsecase struct {
	mock.Mock
}

func (m *MockItemUsecase) GetAllItems() ([]*domain.Item, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Item), args.Error(1)
}

func (m *MockItemUsecase) GetItemByID(itemId string) (*domain.Item, error) {
	args := m.Called(itemId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Item), args.Error(1)
}

func (m *MockItemUsecase) CreateItem(req request.CreateItemRequest) (*domain.Item, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Item), args.Error(1)
}

func (m *MockItemUsecase) UpdateItem(req request.UpdateItemRequest) (*domain.Item, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Item), args.Error(1)
}

func (m *MockItemUsecase) DeleteItem(itemId string) error {
	args := m.Called(itemId)
	return args.Error(0)
}

func TestAdminItemController_GetAllItems(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/admin/items", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUsecase := new(MockItemUsecase)
	controller := NewAdminItemController(mockUsecase)

	var err error
	itemId, err := domain.NewItemId("f47ac10b-58cc-4372-a567-0e02b2c3d479")
	assert.NoError(t, err)
	userId, err := domain.NewUserId("f47ac10b-58cc-4372-a567-0e02b2c3d480")
	assert.NoError(t, err)
	itemName, err := domain.NewItemName("Test Item")
	assert.NoError(t, err)
	stock, err := domain.NewStock(true)
	assert.NoError(t, err)
	description, err := domain.NewDescription("Test Description")
	assert.NoError(t, err)
	item, err := domain.NewItem(itemId, *userId, *itemName, *stock, *description)
	assert.NoError(t, err)
	assert.NotNil(t, item)

	mockUsecase.On("GetAllItems").Return([]*domain.Item{item}, nil)

	err = controller.GetAllItems(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestAdminItemController_GetItemByID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/admin/items/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockUsecase := new(MockItemUsecase)
	controller := NewAdminItemController(mockUsecase)

	var err error
	itemId, err := domain.NewItemId("f47ac10b-58cc-4372-a567-0e02b2c3d481")
	assert.NoError(t, err)
	userId, err := domain.NewUserId("f47ac10b-58cc-4372-a567-0e02b2c3d482")
	assert.NoError(t, err)
	itemName, err := domain.NewItemName("Test Item")
	assert.NoError(t, err)
	stock, err := domain.NewStock(true)
	assert.NoError(t, err)
	description, err := domain.NewDescription("Test Description")
	assert.NoError(t, err)
	item, err := domain.NewItem(itemId, *userId, *itemName, *stock, *description)
	assert.NoError(t, err)
	assert.NotNil(t, item)

	mockUsecase.On("GetItemByID", "1").Return(item, nil)

	err = controller.GetItemByID(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestAdminItemController_GetItemByID_InvalidID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/admin/items/invalid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	mockUsecase := new(MockItemUsecase)
	controller := NewAdminItemController(mockUsecase)

	mockUsecase.On("GetItemByID", "invalid").Return(nil, errors.New("item not found"))

	err := controller.GetItemByID(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestAdminItemController_CreateItem(t *testing.T) {
	e := echo.New()
	e.Validator = &MockValidator{}
	
	requestBody := map[string]interface{}{
		"item_name":    "Test Item",
		"stock":        true,
		"description": "Test Description",
	}
	jsonBody, _ := json.Marshal(requestBody)
	
	req := httptest.NewRequest(http.MethodPost, "/admin/items", bytes.NewBuffer(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user_id", "f47ac10b-58cc-4372-a567-0e02b2c3d483")

	mockUsecase := new(MockItemUsecase)
	controller := NewAdminItemController(mockUsecase)

	var err error
	itemId, err := domain.NewItemId("f47ac10b-58cc-4372-a567-0e02b2c3d484")
	assert.NoError(t, err)
	userId, err := domain.NewUserId("f47ac10b-58cc-4372-a567-0e02b2c3d485")
	assert.NoError(t, err)
	itemName, err := domain.NewItemName("Test Item")
	assert.NoError(t, err)
	stock, err := domain.NewStock(true)
	assert.NoError(t, err)
	description, err := domain.NewDescription("Test Description")
	assert.NoError(t, err)
	item, err := domain.NewItem(itemId, *userId, *itemName, *stock, *description)
	assert.NoError(t, err)
	assert.NotNil(t, item)

	mockUsecase.On("CreateItem", mock.AnythingOfType("request.CreateItemRequest")).Return(item, nil)

	err = controller.CreateItem(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestAdminItemController_UpdateItem(t *testing.T) {
	e := echo.New()
	e.Validator = &MockValidator{}
	
	requestBody := map[string]interface{}{
		"item_name":    "Updated Item",
		"stock":        false,
		"description": "Updated Description",
	}
	jsonBody, _ := json.Marshal(requestBody)
	
	req := httptest.NewRequest(http.MethodPut, "/admin/items/1", bytes.NewBuffer(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockUsecase := new(MockItemUsecase)
	controller := NewAdminItemController(mockUsecase)

	var err error
	itemId, err := domain.NewItemId("f47ac10b-58cc-4372-a567-0e02b2c3d486")
	assert.NoError(t, err)
	userId, err := domain.NewUserId("f47ac10b-58cc-4372-a567-0e02b2c3d487")
	assert.NoError(t, err)
	itemName, err := domain.NewItemName("Updated Item")
	assert.NoError(t, err)
	stock, err := domain.NewStock(false)
	assert.NoError(t, err)
	description, err := domain.NewDescription("Updated Description")
	assert.NoError(t, err)
	item, err := domain.NewItem(itemId, *userId, *itemName, *stock, *description)
	assert.NoError(t, err)
	assert.NotNil(t, item)

	mockUsecase.On("UpdateItem", mock.AnythingOfType("request.UpdateItemRequest")).Return(item, nil)

	err = controller.UpdateItem(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestAdminItemController_DeleteItem(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/admin/items/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockUsecase := new(MockItemUsecase)
	controller := NewAdminItemController(mockUsecase)

	mockUsecase.On("DeleteItem", "1").Return(nil)

	err := controller.DeleteItem(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestAdminItemController_DeleteItem_Error(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/admin/items/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockUsecase := new(MockItemUsecase)
	controller := NewAdminItemController(mockUsecase)

	mockUsecase.On("DeleteItem", "1").Return(errors.New("delete failed"))

	err := controller.DeleteItem(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUsecase.AssertExpectations(t)
}