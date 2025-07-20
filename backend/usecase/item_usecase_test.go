package usecase

import (
	"errors"
	"testing"

	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/usecase/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockItemRepository struct {
	mock.Mock
}

func (m *MockItemRepository) GetAllItems() (domain.Items, error) {
	args := m.Called()
	return args.Get(0).(domain.Items), args.Error(1)
}

func (m *MockItemRepository) GetItemByID(itemId *domain.ItemId) (*domain.Item, error) {
	args := m.Called(itemId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Item), args.Error(1)
}

func (m *MockItemRepository) CreateItem(item *domain.Item) (*domain.Item, error) {
	args := m.Called(item)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Item), args.Error(1)
}

func (m *MockItemRepository) UpdateItem(item *domain.Item) (*domain.Item, error) {
	args := m.Called(item)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Item), args.Error(1)
}

func (m *MockItemRepository) DeleteItem(itemId *domain.ItemId) error {
	args := m.Called(itemId)
	return args.Error(0)
}

func TestGetAllItems_ReturnsItems(t *testing.T) {
	mockRepo := new(MockItemRepository)
	uc := NewItemUsecase(mockRepo)
	result := domain.Items{}
	mockRepo.On("GetAllItems").Return(result, nil)
	items, err := uc.GetAllItems()
	expected := []*domain.Item{}
	assert.NoError(t, err)
	assert.Equal(t, expected, items)
}

func TestCreateItem_Success(t *testing.T) {
	mockRepo := new(MockItemRepository)
	uc := NewItemUsecase(mockRepo)

	req := request.CreateItemRequest{
		ItemName:    "Test Item",
		Stock:       true,
		Description: "Test Description",
		UserId:      "f47ac10b-58cc-4372-a567-0e02b2c3d400",
	}
	itemId, _ := domain.NewItemId("f47ac10b-58cc-4372-a567-0e02b2c3d401")
	userIdValue, _ := domain.NewUserId(req.UserId)
	itemName, _ := domain.NewItemName(req.ItemName)
	stock, _ := domain.NewStock(req.Stock)
	description, _ := domain.NewDescription(req.Description)
	domainItem, _ := domain.NewItem(itemId, *userIdValue, *itemName, *stock, *description)

	mockRepo.On("CreateItem", mock.AnythingOfType("*domain.Item")).Return(domainItem, nil)
	result, err := uc.CreateItem(req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, domainItem.ItemId(), result.ItemId())
	assert.Equal(t, domainItem.ItemName(), result.ItemName())
	assert.Equal(t, domainItem.Stock(), result.Stock())
	assert.Equal(t, domainItem.Description(), result.Description())
	mockRepo.AssertExpectations(t)
}

func TestCreateItem_InvalidItemName(t *testing.T) {
	mockRepo := new(MockItemRepository)
	uc := NewItemUsecase(mockRepo)

	req := request.CreateItemRequest{
		ItemName:    "",
		Stock:       true,
		Description: "Test Description",
		UserId:      "f47ac10b-58cc-4372-a567-0e02b2c3d400",
	}
	result, err := uc.CreateItem(req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "value count must be greater than 0")
	mockRepo.AssertNotCalled(t, "CreateItem")
}

func TestCreateItem_InvalidUserId(t *testing.T) {
	mockRepo := new(MockItemRepository)
	uc := NewItemUsecase(mockRepo)

	req := request.CreateItemRequest{
		ItemName:    "Test Item",
		Stock:       true,
		Description: "Test Description",
		UserId:      "invalid-uuid",
	}
	result, err := uc.CreateItem(req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "invalid UUID")
	mockRepo.AssertNotCalled(t, "CreateItem")
}

func TestCreateItem_RepositoryError(t *testing.T) {
	mockRepo := new(MockItemRepository)
	uc := NewItemUsecase(mockRepo)

	req := request.CreateItemRequest{
		ItemName:    "Test Item",
		Stock:       true,
		Description: "Test Description",
		UserId:      "f47ac10b-58cc-4372-a567-0e02b2c3d400",
	}
	mockRepo.On("CreateItem", mock.AnythingOfType("*domain.Item")).Return(nil, errors.New("database error"))
	result, err := uc.CreateItem(req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "database error", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetItemByID_Success(t *testing.T) {
	mockRepo := new(MockItemRepository)
	uc := NewItemUsecase(mockRepo)

	itemId, _ := domain.NewItemId("f47ac10b-58cc-4372-a567-0e02b2c3d401")
	userId, _ := domain.NewUserId("f47ac10b-58cc-4372-a567-0e02b2c3d400")
	itemName, _ := domain.NewItemName("Test Item")
	stock, _ := domain.NewStock(true)
	description, _ := domain.NewDescription("Test Description")
	domainItem, _ := domain.NewItem(itemId, *userId, *itemName, *stock, *description)

	itemIdValue, _ := domain.NewItemId("f47ac10b-58cc-4372-a567-0e02b2c3d401")
	mockRepo.On("GetItemByID", itemIdValue).Return(domainItem, nil)
	result, err := uc.GetItemByID("f47ac10b-58cc-4372-a567-0e02b2c3d401")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, domainItem, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateItem_Success(t *testing.T) {
	mockRepo := new(MockItemRepository)
	uc := NewItemUsecase(mockRepo)

	req := request.UpdateItemRequest{
		ItemId:      "f47ac10b-58cc-4372-a567-0e02b2c3d401",
		ItemName:    "Updated Item",
		Stock:       false,
		Description: "Updated Description",
	}

	itemId, _ := domain.NewItemId(req.ItemId)
	userId, _ := domain.NewUserId("f47ac10b-58cc-4372-a567-0e02b2c3d400")
	existingItemName, _ := domain.NewItemName("Existing Item")
	existingStock, _ := domain.NewStock(true)
	existingDescription, _ := domain.NewDescription("Existing Description")
	existingItem, _ := domain.NewItem(itemId, *userId, *existingItemName, *existingStock, *existingDescription)

	updatedItemName, _ := domain.NewItemName(req.ItemName)
	updatedStock, _ := domain.NewStock(req.Stock)
	updatedDescription, _ := domain.NewDescription(req.Description)
	updatedItem, _ := domain.NewItem(itemId, *userId, *updatedItemName, *updatedStock, *updatedDescription)

	mockRepo.On("GetItemByID", itemId).Return(existingItem, nil)
	mockRepo.On("UpdateItem", mock.AnythingOfType("*domain.Item")).Return(updatedItem, nil)
	result, err := uc.UpdateItem(req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteItem_Success(t *testing.T) {
	mockRepo := new(MockItemRepository)
	uc := NewItemUsecase(mockRepo)

	itemIdValue, _ := domain.NewItemId("f47ac10b-58cc-4372-a567-0e02b2c3d401")
	mockRepo.On("DeleteItem", itemIdValue).Return(nil)
	err := uc.DeleteItem("f47ac10b-58cc-4372-a567-0e02b2c3d401")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
