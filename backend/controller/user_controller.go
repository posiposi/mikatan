package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/posiposi/project/backend/presenter"
	"github.com/posiposi/project/backend/usecase"
	"github.com/posiposi/project/backend/usecase/request"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CheckAuth(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
	up presenter.IUserPresenter
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	up := presenter.NewUserPresenter()
	return &userController{uu, up}
}

func (uc *userController) SignUp(c echo.Context) error {
	var req struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}
	
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	signUpReq := request.SignUpRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	
	user, err := uc.uu.SignUp(signUpReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	response := uc.up.ToJSON(user)
	return c.JSON(http.StatusCreated, response)
}

func (uc *userController) LogIn(c echo.Context) error {
	var req struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	logInReq := request.LogInRequest{
		Email:    req.Email,
		Password: req.Password,
	}
	
	tokenString, user, err := uc.uu.Login(logInReq)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	
	response := uc.up.ToLoginJSON(tokenString, user)
	return c.JSON(http.StatusOK, response)
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CheckAuth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]bool{"authenticated": true})
}
