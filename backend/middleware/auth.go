// Package middleware provides HTTP middleware functions for authentication and authorization.
package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("token")
			var tokenString string

			if err == nil && cookie != nil {
				tokenString = cookie.Value
			} else {
				auth := c.Request().Header.Get("Authorization")
				if auth == "" {
					return c.JSON(http.StatusUnauthorized, "missing authentication token")
				}

				parts := strings.Split(auth, " ")
				if len(parts) != 2 || parts[0] != "Bearer" {
					return c.JSON(http.StatusUnauthorized, "invalid authorization header format")
				}
				tokenString = parts[1]
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(os.Getenv("SECRET")), nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, "invalid or expired token")
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				if userID, exists := claims["user_id"].(string); exists {
					c.Set("user_id", userID)
					return next(c)
				}
			}

			return c.JSON(http.StatusUnauthorized, "invalid token claims")
		}
	}
}
