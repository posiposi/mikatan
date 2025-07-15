package usecase

import (
	"errors"
	"testing"

	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/model"
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

func (m *MockItemRepository) CreateItem(item *domain.Item) (*domain.Item, error) {
	args := m.Called(item)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Item), args.Error(1)
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

	req := model.Item{
		ItemName:    "Test Item",
		Stock:       true,
		Description: "Test Description",
	}
	userID := "f47ac10b-58cc-4372-a567-0e02b2c3d400"
	itemID, _ := domain.NewItemID("f47ac10b-58cc-4372-a567-0e02b2c3d401")
	userIDValue, _ := domain.NewUserID(userID)
	itemName, _ := domain.NewItemName(req.ItemName)
	stock, _ := domain.NewStock(req.Stock)
	description, _ := domain.NewDescription(req.Description)
	domainItem, _ := domain.NewItem(itemID, *userIDValue, *itemName, *stock, *description)

	mockRepo.On("CreateItem", mock.AnythingOfType("*domain.Item")).Return(domainItem, nil)
	result, err := uc.CreateItem(req, userID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, domainItem.ItemID(), result.ItemID())
	assert.Equal(t, domainItem.ItemName(), result.ItemName())
	assert.Equal(t, domainItem.Stock(), result.Stock())
	assert.Equal(t, domainItem.Description(), result.Description())
	mockRepo.AssertExpectations(t)
}

func TestCreateItem_InvalidItemName(t *testing.T) {
	mockRepo := new(MockItemRepository)
	uc := NewItemUsecase(mockRepo)

	req := model.Item{
		ItemName:    "",
		Stock:       true,
		Description: "Test Description",
	}
	userID := "f47ac10b-58cc-4372-a567-0e02b2c3d400"
	result, err := uc.CreateItem(req, userID)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "value count must be greater than 0")
	mockRepo.AssertNotCalled(t, "CreateItem")
}

func TestCreateItem_InvalidUserID(t *testing.T) {
	mockRepo := new(MockItemRepository)
	uc := NewItemUsecase(mockRepo)

	req := model.Item{
		ItemName:    "Test Item",
		Stock:       true,
		Description: "Test Description",
	}
	invalidUserID := "invalid-uuid"
	result, err := uc.CreateItem(req, invalidUserID)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "invalid UUID")
	mockRepo.AssertNotCalled(t, "CreateItem")
}

func TestCreateItem_RepositoryError(t *testing.T) {
	mockRepo := new(MockItemRepository)
	uc := NewItemUsecase(mockRepo)

	req := model.Item{
		ItemName:    "Test Item",
		Stock:       true,
		Description: "Test Description",
	}
	userID := "f47ac10b-58cc-4372-a567-0e02b2c3d400"
	mockRepo.On("CreateItem", mock.AnythingOfType("*domain.Item")).Return(nil, errors.New("database error"))
	result, err := uc.CreateItem(req, userID)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "database error", err.Error())
	mockRepo.AssertExpectations(t)
}
