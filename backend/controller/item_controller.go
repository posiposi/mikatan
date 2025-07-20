// Package controller handles HTTP request/response processing and input validation.
package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/posiposi/project/backend/presenter"
	"github.com/posiposi/project/backend/usecase"
	"github.com/posiposi/project/backend/usecase/request"
)

type IItemController interface {
	GetAllItems(c echo.Context) error
	CreateItem(c echo.Context) error
}

type itemController struct {
	iu usecase.IItemUsecase
	ip presenter.IItemPresenter
}

func NewItemController(iu usecase.IItemUsecase) IItemController {
	ip := presenter.NewItemPresenter()
	return &itemController{iu, ip}
}

func (ic *itemController) GetAllItems(c echo.Context) error {
	items, err := ic.iu.GetAllItems()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	response := ic.ip.ToJSONList(items)
	return c.JSON(http.StatusOK, response)
}

func (ic *itemController) CreateItem(c echo.Context) error {
	var req struct {
		ItemName    string `json:"item_name" validate:"required"`
		Stock       bool   `json:"stock"`
		Description string `json:"description"`
	}
	
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	userId, ok := c.Get("user_id").(string)
	if !ok || userId == "" {
		return c.JSON(http.StatusUnauthorized, "user_id not found in context")
	}
	
	createReq := request.CreateItemRequest{
		ItemName:    req.ItemName,
		Stock:       req.Stock,
		Description: req.Description,
		UserId:      userId,
	}
	
	createdItem, err := ic.iu.CreateItem(createReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	response := ic.ip.ToJSON(createdItem)
	return c.JSON(http.StatusCreated, response)
}
