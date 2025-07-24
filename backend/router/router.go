// Package router defines API routes and middleware configuration.
package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/posiposi/project/backend/controller"
	authMiddleware "github.com/posiposi/project/backend/middleware"
	"github.com/posiposi/project/backend/validator"
)

func NewRouter(uc controller.IUserController, ic controller.IItemController, aic controller.IAdminItemController, aac controller.IAdminAuthController, userRepo authMiddleware.UserRepository) *echo.Echo {
	e := echo.New()
	e.Validator = validator.NewValidator()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://labstack.com", "https://labstack.net", "http://localhost:3000", "https://localhost:3000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	g := e.Group("/v1")
	g.POST("/signup", uc.SignUp)
	g.POST("/login", uc.LogIn)
	g.POST("/logout", uc.LogOut)
	g.GET("/auth/check", uc.CheckAuth, authMiddleware.AuthMiddleware())
	i := g.Group("/items")
	i.GET("", ic.GetAllItems)
	i.POST("", ic.CreateItem, authMiddleware.AuthMiddleware())
	
	admin := g.Group("/admin", authMiddleware.AuthMiddleware(), authMiddleware.AdminMiddleware(userRepo))
	admin.GET("/auth/check", aac.CheckAdminAuth)
	adminItems := admin.Group("/items")
	adminItems.GET("", aic.GetAllItems)
	adminItems.GET("/:id", aic.GetItemByID)
	adminItems.POST("", aic.CreateItem)
	adminItems.PUT("/:id", aic.UpdateItem)
	adminItems.DELETE("/:id", aic.DeleteItem)
	
	return e
}
