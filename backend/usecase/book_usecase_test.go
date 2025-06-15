package usecase

import (
	"testing"

	"github.com/posiposi/project/backend/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) GetAllBooks(books *[]model.Book) error {
	args := m.Called(books)
	return args.Error(0)
}

func (m *MockBookRepository) GetBookByBookId(book *model.Book, bookId string) error {
	args := m.Called(book, bookId)
	return args.Error(0)
}

func (m *MockBookRepository) CreateBook(book *model.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookRepository) UpdateBook(book *model.Book, bookId string) error {
	args := m.Called(book, bookId)
	return args.Error(0)
}

func (m *MockBookRepository) DeleteBook(bookId string) error {
	args := m.Called(bookId)
	return args.Error(0)
}

func TestDeleteBook(t *testing.T) {
	mockRepo := new(MockBookRepository)
	usecase := NewBookUsecase(mockRepo)

	bookId := "1"
	mockRepo.On("DeleteBook", bookId).Return(nil)

	err := usecase.DeleteBook(bookId)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
