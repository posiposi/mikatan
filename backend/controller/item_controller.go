package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/posiposi/project/backend/usecase"
)

type IItemController interface {
	GetAllItems(c echo.Context) error
}

type itemController struct {
	iu usecase.IItemUsecase
}

func NewItemController(iu usecase.IItemUsecase) IItemController {
	return &itemController{iu}
}

func (ic *itemController) GetAllItems(c echo.Context) error {
	itemsRes, err := ic.iu.GetAllItems()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, itemsRes)
}
