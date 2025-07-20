package middleware

import (
	"net/http"

	"github.com/posiposi/project/backend/domain"
	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	GetUserByID(userID string) (*domain.User, error)
}

func AdminMiddleware(userRepo UserRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := c.Get("user_id")
			if userID == nil {
				return c.JSON(http.StatusUnauthorized, "user not authenticated")
			}

			userIDStr, ok := userID.(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, "invalid user id")
			}

			user, err := userRepo.GetUserByID(userIDStr)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "user not found")
			}

			if user == nil {
				return c.JSON(http.StatusUnauthorized, "user is nil")
			}

			if user.Role() == nil {
				return c.JSON(http.StatusUnauthorized, "user role is nil")
			}

			if user.Role().Value() != "ADMINISTRATOR" {
				return c.JSON(http.StatusForbidden, "admin permission required")
			}

			return next(c)
		}
	}
}