package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/posiposi/project/backend/model"
	"github.com/posiposi/project/backend/usecase"
)

type IBookController interface {
	GetAllBooks(c echo.Context) error
	GetBookByBookId(c echo.Context) error
	CreateBook(c echo.Context) error
	UpdateBook(c echo.Context) error
	DeleteBook(c echo.Context) error
	RecommendBooks(c echo.Context) error
}

type bookController struct {
	bu usecase.IBookUsecase
}

type chatResponse struct {
	message string
	tokens  int
}

func NewBookController(bu usecase.IBookUsecase) IBookController {
	return &bookController{bu}
}

func (bc *bookController) GetAllBooks(c echo.Context) error {
	booksRes, err := bc.bu.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, booksRes)
}

func (bc *bookController) GetBookByBookId(c echo.Context) error {
	bookId := c.Param("bookId")
	bookRes, err := bc.bu.GetBookByBookId(bookId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, bookRes)
}

func (bc *bookController) CreateBook(c echo.Context) error {
	book := model.Book{}
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	bookRes, err := bc.bu.CreateBook(book)
	if err != nil {
		// TODO より詳細なエラーステータスコードを返す
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, bookRes)
}

func (bc *bookController) UpdateBook(c echo.Context) error {
	bookId := c.Param("bookId")
	book := model.Book{}
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	bookRes, err := bc.bu.UpdateBook(book, bookId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, bookRes)
}

func (bc *bookController) DeleteBook(c echo.Context) error {
	bookId := c.Param("bookId")
	err := bc.bu.DeleteBook(bookId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, model.MessageResponse{Message: "書籍を削除しました。"})
}

func (bc *bookController) RecommendBooks(c echo.Context) error {
	chatResp, err := bc.bu.FetchRecommendBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, chatResp)
}
