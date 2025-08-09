package usecase

import (
	"context"
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

type MockPriceRepository struct {
	mock.Mock
}

func (m *MockPriceRepository) Create(ctx context.Context, price *domain.Price) error {
	args := m.Called(ctx, price)
	return args.Error(0)
}

func (m *MockPriceRepository) FindById(ctx context.Context, priceId string) (*domain.Price, error) {
	args := m.Called(ctx, priceId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Price), args.Error(1)
}

func (m *MockPriceRepository) FindByItemId(ctx context.Context, itemId string) ([]*domain.Price, error) {
	args := m.Called(ctx, itemId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Price), args.Error(1)
}

func (m *MockPriceRepository) FindCurrentByItemId(ctx context.Context, itemId string) (*domain.Price, error) {
	args := m.Called(ctx, itemId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Price), args.Error(1)
}

func (m *MockPriceRepository) UpdateByItemId(ctx context.Context, itemId string, price *domain.Price) error {
	args := m.Called(ctx, itemId, price)
	return args.Error(0)
}

func (m *MockPriceRepository) Delete(ctx context.Context, priceId string) error {
	args := m.Called(ctx, priceId)
	return args.Error(0)
}

func TestGetAllItems_ReturnsItems(t *testing.T) {
	mockItemRepo := new(MockItemRepository)
	mockPriceRepo := new(MockPriceRepository)
	uc := NewItemUsecase(mockItemRepo, mockPriceRepo)
	result := domain.Items{}
	mockItemRepo.On("GetAllItems").Return(result, nil)
	items, err := uc.GetAllItems()
	expected := []*domain.Item{}
	assert.NoError(t, err)
	assert.Equal(t, expected, items)
}

func TestCreateItem_Success(t *testing.T) {
	mockItemRepo := new(MockItemRepository)
	mockPriceRepo := new(MockPriceRepository)
	uc := NewItemUsecase(mockItemRepo, mockPriceRepo)

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

	mockItemRepo.On("CreateItem", mock.AnythingOfType("*domain.Item")).Return(domainItem, nil)
	result, err := uc.CreateItem(req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, domainItem.ItemId(), result.ItemId())
	assert.Equal(t, domainItem.ItemName(), result.ItemName())
	assert.Equal(t, domainItem.Stock(), result.Stock())
	assert.Equal(t, domainItem.Description(), result.Description())
	mockItemRepo.AssertExpectations(t)
}

func TestCreateItem_WithPrice_Success(t *testing.T) {
	mockItemRepo := new(MockItemRepository)
	mockPriceRepo := new(MockPriceRepository)
	uc := NewItemUsecase(mockItemRepo, mockPriceRepo)

	priceWithTax := 1100
	priceWithoutTax := 1000
	taxRate := 10.0
	currency := "JPY"

	req := request.CreateItemRequest{
		ItemName:        "Test Item",
		Stock:           true,
		Description:     "Test Description",
		UserId:          "f47ac10b-58cc-4372-a567-0e02b2c3d400",
		PriceWithTax:    &priceWithTax,
		PriceWithoutTax: &priceWithoutTax,
		TaxRate:         &taxRate,
		Currency:        &currency,
	}

	itemId, _ := domain.NewItemId("f47ac10b-58cc-4372-a567-0e02b2c3d401")
	userIdValue, _ := domain.NewUserId(req.UserId)
	itemName, _ := domain.NewItemName(req.ItemName)
	stock, _ := domain.NewStock(req.Stock)
	description, _ := domain.NewDescription(req.Description)
	domainItem, _ := domain.NewItem(itemId, *userIdValue, *itemName, *stock, *description)

	mockItemRepo.On("CreateItem", mock.AnythingOfType("*domain.Item")).Return(domainItem, nil)
	mockPriceRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Price")).Return(nil)

	result, err := uc.CreateItem(req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, domainItem.ItemId(), result.ItemId())
	mockItemRepo.AssertExpectations(t)
	mockPriceRepo.AssertExpectations(t)
}

func TestCreateItem_WithPrice_PriceCreationFails(t *testing.T) {
	mockItemRepo := new(MockItemRepository)
	mockPriceRepo := new(MockPriceRepository)
	uc := NewItemUsecase(mockItemRepo, mockPriceRepo)

	priceWithTax := 1100
	priceWithoutTax := 1000
	taxRate := 10.0
	currency := "JPY"

	req := request.CreateItemRequest{
		ItemName:        "Test Item",
		Stock:           true,
		Description:     "Test Description",
		UserId:          "f47ac10b-58cc-4372-a567-0e02b2c3d400",
		PriceWithTax:    &priceWithTax,
		PriceWithoutTax: &priceWithoutTax,
		TaxRate:         &taxRate,
		Currency:        &currency,
	}

	itemId, _ := domain.NewItemId("f47ac10b-58cc-4372-a567-0e02b2c3d401")
	userIdValue, _ := domain.NewUserId(req.UserId)
	itemName, _ := domain.NewItemName(req.ItemName)
	stock, _ := domain.NewStock(req.Stock)
	description, _ := domain.NewDescription(req.Description)
	domainItem, _ := domain.NewItem(itemId, *userIdValue, *itemName, *stock, *description)

	mockItemRepo.On("CreateItem", mock.AnythingOfType("*domain.Item")).Return(domainItem, nil)
	mockPriceRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Price")).Return(errors.New("price creation failed"))

	result, err := uc.CreateItem(req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "price creation failed", err.Error())
	mockItemRepo.AssertExpectations(t)
	mockPriceRepo.AssertExpectations(t)
}

func TestCreateItem_InvalidItemName(t *testing.T) {
	mockItemRepo := new(MockItemRepository)
	mockPriceRepo := new(MockPriceRepository)
	uc := NewItemUsecase(mockItemRepo, mockPriceRepo)

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
	mockItemRepo.AssertNotCalled(t, "CreateItem")
}

func TestCreateItem_InvalidUserId(t *testing.T) {
	mockItemRepo := new(MockItemRepository)
	mockPriceRepo := new(MockPriceRepository)
	uc := NewItemUsecase(mockItemRepo, mockPriceRepo)

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
	mockItemRepo.AssertNotCalled(t, "CreateItem")
}

func TestCreateItem_RepositoryError(t *testing.T) {
	mockItemRepo := new(MockItemRepository)
	mockPriceRepo := new(MockPriceRepository)
	uc := NewItemUsecase(mockItemRepo, mockPriceRepo)

	req := request.CreateItemRequest{
		ItemName:    "Test Item",
		Stock:       true,
		Description: "Test Description",
		UserId:      "f47ac10b-58cc-4372-a567-0e02b2c3d400",
	}
	mockItemRepo.On("CreateItem", mock.AnythingOfType("*domain.Item")).Return(nil, errors.New("database error"))
	result, err := uc.CreateItem(req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "database error", err.Error())
	mockItemRepo.AssertExpectations(t)
}

func TestGetItemByID_Success(t *testing.T) {
	mockItemRepo := new(MockItemRepository)
	mockPriceRepo := new(MockPriceRepository)
	uc := NewItemUsecase(mockItemRepo, mockPriceRepo)

	itemId, _ := domain.NewItemId("f47ac10b-58cc-4372-a567-0e02b2c3d401")
	userId, _ := domain.NewUserId("f47ac10b-58cc-4372-a567-0e02b2c3d400")
	itemName, _ := domain.NewItemName("Test Item")
	stock, _ := domain.NewStock(true)
	description, _ := domain.NewDescription("Test Description")
	domainItem, _ := domain.NewItem(itemId, *userId, *itemName, *stock, *description)

	itemIdValue, _ := domain.NewItemId("f47ac10b-58cc-4372-a567-0e02b2c3d401")
	mockItemRepo.On("GetItemByID", itemIdValue).Return(domainItem, nil)
	result, err := uc.GetItemByID("f47ac10b-58cc-4372-a567-0e02b2c3d401")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, domainItem, result)
	mockItemRepo.AssertExpectations(t)
}

func TestUpdateItem_Success(t *testing.T) {
	mockItemRepo := new(MockItemRepository)
	mockPriceRepo := new(MockPriceRepository)
	uc := NewItemUsecase(mockItemRepo, mockPriceRepo)

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

	mockItemRepo.On("GetItemByID", itemId).Return(existingItem, nil)
	mockItemRepo.On("UpdateItem", mock.AnythingOfType("*domain.Item")).Return(updatedItem, nil)
	result, err := uc.UpdateItem(req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockItemRepo.AssertExpectations(t)
}

func TestUpdateItem_WithPrice_Success(t *testing.T) {
	mockItemRepo := new(MockItemRepository)
	mockPriceRepo := new(MockPriceRepository)
	uc := NewItemUsecase(mockItemRepo, mockPriceRepo)

	oldPriceWithTax := 1100
	oldPriceWithoutTax := 1000

	newPriceWithTax := 2200
	newPriceWithoutTax := 2000
	newTaxRate := 10.0
	newCurrency := "JPY"

	req := request.UpdateItemRequest{
		ItemId:          "f47ac10b-58cc-4372-a567-0e02b2c3d401",
		ItemName:        "Updated Item",
		Stock:           false,
		Description:     "Updated Description",
		PriceWithTax:    &newPriceWithTax,
		PriceWithoutTax: &newPriceWithoutTax,
		TaxRate:         &newTaxRate,
		Currency:        &newCurrency,
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


	mockItemRepo.On("GetItemByID", itemId).Return(existingItem, nil)
	mockItemRepo.On("UpdateItem", mock.AnythingOfType("*domain.Item")).Return(updatedItem, nil)
	
	mockPriceRepo.On("UpdateByItemId", mock.Anything, req.ItemId, mock.MatchedBy(func(price *domain.Price) bool {
		return price.PriceWithTax() == newPriceWithTax &&
			price.PriceWithoutTax() == newPriceWithoutTax &&
			price.TaxRate() == newTaxRate &&
			price.Currency() == newCurrency
	})).Return(nil)

	result, err := uc.UpdateItem(req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockItemRepo.AssertExpectations(t)
	mockPriceRepo.AssertExpectations(t)

	assert.NotEqual(t, oldPriceWithTax, newPriceWithTax, "価格が更新されていること")
	assert.NotEqual(t, oldPriceWithoutTax, newPriceWithoutTax, "税抜価格が更新されていること")
}

func TestUpdateItem_WithPrice_PriceUpdateFails(t *testing.T) {
	mockItemRepo := new(MockItemRepository)
	mockPriceRepo := new(MockPriceRepository)
	uc := NewItemUsecase(mockItemRepo, mockPriceRepo)

	priceWithTax := 2200
	priceWithoutTax := 2000
	taxRate := 10.0
	currency := "JPY"

	req := request.UpdateItemRequest{
		ItemId:          "f47ac10b-58cc-4372-a567-0e02b2c3d401",
		ItemName:        "Updated Item",
		Stock:           false,
		Description:     "Updated Description",
		PriceWithTax:    &priceWithTax,
		PriceWithoutTax: &priceWithoutTax,
		TaxRate:         &taxRate,
		Currency:        &currency,
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

	mockItemRepo.On("GetItemByID", itemId).Return(existingItem, nil)
	mockItemRepo.On("UpdateItem", mock.AnythingOfType("*domain.Item")).Return(updatedItem, nil)
	mockPriceRepo.On("UpdateByItemId", mock.Anything, req.ItemId, mock.AnythingOfType("*domain.Price")).Return(errors.New("price update failed"))

	result, err := uc.UpdateItem(req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "price update failed", err.Error())
	mockItemRepo.AssertExpectations(t)
	mockPriceRepo.AssertExpectations(t)
}

func TestDeleteItem_Success(t *testing.T) {
	mockItemRepo := new(MockItemRepository)
	mockPriceRepo := new(MockPriceRepository)
	uc := NewItemUsecase(mockItemRepo, mockPriceRepo)

	itemIdValue, _ := domain.NewItemId("f47ac10b-58cc-4372-a567-0e02b2c3d401")
	mockItemRepo.On("DeleteItem", itemIdValue).Return(nil)
	err := uc.DeleteItem("f47ac10b-58cc-4372-a567-0e02b2c3d401")

	assert.NoError(t, err)
	mockItemRepo.AssertExpectations(t)
}
