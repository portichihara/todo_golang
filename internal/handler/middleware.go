package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"todo-api/pkg/auth"
)

func JWTMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing token"})
			}

			token, err := jwt.ParseWithClaims(tokenString[7:], &auth.JWTCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			}

			claims := token.Claims.(*auth.JWTCustomClaims)
			c.Set("userID", claims.UserID)
			return next(c)
		}
	}
}
