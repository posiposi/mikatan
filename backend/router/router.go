package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/posiposi/project/backend/controller"
)

func NewRouter(uc controller.IUserController, bc controller.IBookController, ic controller.IItemController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net", "http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	g := e.Group("/v1")
	u := g.Group("/secret-admin")
	u.POST("/signup", uc.SignUp)
	u.POST("/login", uc.LogIn)
	u.POST("/logout", uc.LogOut)
	b := g.Group("/books")
	b.GET("", bc.GetAllBooks)
	b.GET("/:bookId", bc.GetBookByBookId)
	// TODO GET以外のメソッドはログインしているユーザーのみアクセス可能
	b.POST("", bc.CreateBook)
	b.PUT("/:bookId", bc.UpdateBook)
	b.DELETE("/:bookId", bc.DeleteBook)
	i := g.Group("/items")
	i.GET("", ic.GetAllItems)
	return e
}
