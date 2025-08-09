package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/posiposi/project/backend/presenter"
	"github.com/posiposi/project/backend/usecase"
	"github.com/posiposi/project/backend/usecase/request"
)

type IAdminItemController interface {
	GetAllItems(c echo.Context) error
	GetItemByID(c echo.Context) error
	CreateItem(c echo.Context) error
	UpdateItem(c echo.Context) error
	DeleteItem(c echo.Context) error
}

type adminItemController struct {
	iu usecase.IItemUsecase
	ip presenter.IItemPresenter
}

func NewAdminItemController(iu usecase.IItemUsecase) IAdminItemController {
	ip := presenter.NewItemPresenter()
	return &adminItemController{iu, ip}
}

func (aic *adminItemController) GetAllItems(c echo.Context) error {
	items, err := aic.iu.GetAllItems()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	response := aic.ip.ToJSONList(items)
	return c.JSON(http.StatusOK, response)
}

func (aic *adminItemController) GetItemByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "ID is required")
	}
	
	item, err := aic.iu.GetItemByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	response := aic.ip.ToJSON(item)
	return c.JSON(http.StatusOK, response)
}

func (aic *adminItemController) CreateItem(c echo.Context) error {
	var req struct {
		ItemName        string   `json:"item_name" validate:"required"`
		Stock           bool     `json:"stock"`
		Description     string   `json:"description"`
		PriceWithoutTax *int     `json:"price_without_tax"`
		TaxRate         *float64 `json:"tax_rate"`
		Currency        *string  `json:"currency"`
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
		ItemName:        req.ItemName,
		Stock:           req.Stock,
		Description:     req.Description,
		UserId:          userId,
		PriceWithoutTax: req.PriceWithoutTax,
		TaxRate:         req.TaxRate,
		Currency:        req.Currency,
	}

	// 税込み料金を計算
	if req.PriceWithoutTax != nil && req.TaxRate != nil {
		priceWithTax := int(float64(*req.PriceWithoutTax) * (1 + *req.TaxRate/100))
		createReq.PriceWithTax = &priceWithTax
	}

	createdItem, err := aic.iu.CreateItem(createReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	response := aic.ip.ToJSON(createdItem)
	return c.JSON(http.StatusCreated, response)
}

func (aic *adminItemController) UpdateItem(c echo.Context) error {
	id := c.Param("id")
	
	var req struct {
		ItemName        string   `json:"item_name" validate:"required"`
		Stock           bool     `json:"stock"`
		Description     string   `json:"description"`
		PriceWithoutTax *int     `json:"price_without_tax"`
		TaxRate         *float64 `json:"tax_rate"`
		Currency        *string  `json:"currency"`
	}
	
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updateReq := request.UpdateItemRequest{
		ItemId:          id,
		ItemName:        req.ItemName,
		Stock:           req.Stock,
		Description:     req.Description,
		PriceWithoutTax: req.PriceWithoutTax,
		TaxRate:         req.TaxRate,
		Currency:        req.Currency,
	}

	if req.PriceWithoutTax != nil && req.TaxRate != nil {
		priceWithTax := int(float64(*req.PriceWithoutTax) * (1 + *req.TaxRate/100))
		updateReq.PriceWithTax = &priceWithTax
	}

	updatedItem, err := aic.iu.UpdateItem(updateReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	response := aic.ip.ToJSON(updatedItem)
	return c.JSON(http.StatusOK, response)
}

func (aic *adminItemController) DeleteItem(c echo.Context) error {
	id := c.Param("id")

	err := aic.iu.DeleteItem(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent, nil)
}