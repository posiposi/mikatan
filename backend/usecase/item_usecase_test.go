package usecase

import (
	"testing"

	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/dto"
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

func TestGetAllItems_ReturnsItems(t *testing.T) {
	mockRepo := new(MockItemRepository)
	uc := NewItemUsecase(mockRepo)
	result := domain.Items{}
	mockRepo.On("GetAllItems").Return(result, nil)
	items, err := uc.GetAllItems()
	expected := []dto.ItemResponse{}
	assert.NoError(t, err)
	assert.Equal(t, expected, items)
}
