package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var HeaderMiddleware headerMiddleWareInterface = &headerMiddleware{}

type headerMiddleWareInterface interface {
	ServerHeader(next echo.HandlerFunc) echo.HandlerFunc
	CheckHeader(next echo.HandlerFunc) echo.HandlerFunc
}

type headerMiddleware struct{}

func (m *headerMiddleware) ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Hallo", "Welt")
		return next(c)
	}
}

func (m *headerMiddleware) CheckHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqHeader := c.Request().Header.Get("Hello")
		if reqHeader == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "request header not complete",
			})
		}
		return next(c)
	}
}
