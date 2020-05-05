package middleware

import (
	"github.com/labstack/echo/v4"
)

var HeaderMiddleware headerMiddleWareInterface = &headerMiddleware{}

type headerMiddleWareInterface interface {
	ServerHeader(next echo.HandlerFunc) echo.HandlerFunc
}

type headerMiddleware struct{}

func (m *headerMiddleware) ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Hallo", "Welt")
		return next(c)
	}
}
