// Package controller handles HTTP request/response processing and input validation.
package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/posiposi/project/backend/model"
	"github.com/posiposi/project/backend/usecase"
)

type IItemController interface {
	GetAllItems(c echo.Context) error
	CreateItem(c echo.Context) error
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

func (ic *itemController) CreateItem(c echo.Context) error {
	item := model.Item{}
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&item); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userID, ok := c.Get("user_id").(string)
	if !ok || userID == "" {
		return c.JSON(http.StatusUnauthorized, "user_id not found in context")
	}
	itemRes, err := ic.iu.CreateItem(item, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, itemRes)
}
