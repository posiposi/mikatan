package repository

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/posiposi/project/backend/model"
	"gorm.io/gorm"
)

type IBookRepository interface {
	GetAllBooks(books *[]model.Book) error
	GetBookByBookId(book *model.Book, bookId string) error
	CreateBook(book *model.Book) error
	UpdateBook(book *model.Book, bookId string) error
	DeleteBook(bookId string) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &bookRepository{db}
}

func (br *bookRepository) GetAllBooks(books *[]model.Book) error {
	if err := br.db.Select("books.*, readings.progress_page, readings.review").
		Order("created_at").
		Joins("left join readings on readings.book_id = books.id").
		Find(books).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookRepository) GetBookByBookId(book *model.Book, bookId string) error {
	if err := br.db.First(book, bookId).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookRepository) CreateBook(book *model.Book) error {
	br.db.Transaction(func(tx *gorm.DB) error {
		book.ID = uuid.NewString()
		if err := br.db.Omit("ProgressPage", "Review", "ProgressPercentage").Create(book).Error; err != nil {
			return err
		}
		return nil
	})
	return nil
}

func (br *bookRepository) UpdateBook(book *model.Book, bookId string) error {
	result := br.db.Model(&book).Where("id = ?", bookId).Updates(map[string]interface{}{
		"title":        book.Title,
		"genre":        book.Genre,
		"total_page":   book.TotalPage,
		"price":        book.Price,
		"author":       book.Author,
		"publisher":    book.Publisher,
		"published_at": book.PublishedAt,
	})
	// userIDはuser振り分けを実装するまでは暫定的に固定値とする
	// reviewについても感想登録機能実装までは""固定値とする
	reading_result := br.db.Save(&model.Reading{
		BookId:       bookId,
		UserId:       "user_1",
		ProgressPage: book.ProgressPage,
		Review:       "",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})

	if result.Error != nil {
		return result.Error
	}
	if reading_result.Error != nil {
		return reading_result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object is not exist")
	}
	return nil
}

func (br *bookRepository) DeleteBook(bookId string) error {
	result := br.db.Where("id = ?", bookId).Delete(&model.Book{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object is not exist")
	}
	return nil
}
