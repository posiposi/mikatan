// Package router defines API routes and middleware configuration.
package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/posiposi/project/backend/controller"
	authMiddleware "github.com/posiposi/project/backend/middleware"
	"github.com/posiposi/project/backend/validator"
)

func NewRouter(uc controller.IUserController, ic controller.IItemController) *echo.Echo {
	e := echo.New()
	e.Validator = validator.NewValidator()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net", "http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	g := e.Group("/v1")
	g.POST("/signup", uc.SignUp)
	g.POST("/login", uc.LogIn)
	g.POST("/logout", uc.LogOut)
	i := g.Group("/items")
	i.GET("", ic.GetAllItems)
	i.POST("", ic.CreateItem, authMiddleware.AuthMiddleware())
	return e
}
