package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type IAdminAuthController interface {
	CheckAdminAuth(c echo.Context) error
}

type adminAuthController struct{}

func NewAdminAuthController() IAdminAuthController {
	return &adminAuthController{}
}

func (aac *adminAuthController) CheckAdminAuth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Admin access granted",
		"admin":   true,
	})
}