package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db     *sql.DB
	mock   sqlmock.Sqlmock
	gormDB *gorm.DB
	repo   *bookRepository
)

func TestMain(m *testing.M) {
	var err error
	db, mock, err = sqlmock.New()
	if err != nil {
		fmt.Printf("an error '%s' was not expected when opening a stub database connection", err)
		os.Exit(1)
	}
	defer db.Close()

	dsn := "sqlmock_db_0"
	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to open gorm DB: %v", err)
		os.Exit(1)
	}
	repo = &bookRepository{db: gormDB}

	code := m.Run()
	os.Exit(code)
}

func TestDeleteBook(t *testing.T) {
	t.Run("successfully delete book", func(t *testing.T) {
		bookId := "1"
		mock.ExpectExec("DELETE FROM `books` WHERE id = ?").
			WithArgs(bookId).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.DeleteBook(bookId)
		if err == nil {
			t.Errorf("Error was not expected while deleting")

		}
	})

	t.Run("book not found", func(t *testing.T) {
		bookId := "2"
		mock.ExpectExec("DELETE FROM `books` WHERE id = ?").
			WithArgs(bookId).
			WillReturnResult(sqlmock.NewErrorResult(errors.New("record not found"))).
			WillReturnError(errors.New("record not found"))

		err := repo.DeleteBook(bookId)
		if err == nil {
			t.Error("An error should have occurred.")
		}
	})

	t.Run("database error", func(t *testing.T) {
		bookId := "3"
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `books` WHERE id = ?").
			WithArgs(bookId).
			WillReturnError(errors.New("db error"))
		mock.ExpectCommit()

		err := repo.DeleteBook(bookId)
		if err == nil {
			t.Errorf("Error was not expected while deleting")

		}
	})
}
