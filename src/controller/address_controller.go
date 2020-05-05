package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/maei/protocol_buffer_go/src/service"
	"github.com/maei/shared_utils_go/rest_errors"
	"io/ioutil"
	"net/http"
)

var AddressController addressControllerInterface = &addressController{}

type addressControllerInterface interface {
	WriteAddress(echo.Context) error
	ReadAddress(c echo.Context) error
}

type addressController struct{}

func (a *addressController) WriteAddress(c echo.Context) error {
	defer c.Request().Body.Close()

	jsonBody := c.Request().Body

	bs, err := ioutil.ReadAll(jsonBody)
	if err != nil {
		return rest_errors.NewInternalServerError("failed reading json body", err)
	}

	ab, writeErr := service.AddressService.WriteAddress(bs)
	if writeErr != nil {
		return c.JSON(writeErr.Status(), writeErr)
	}
	return c.JSON(http.StatusOK, ab)
}

func (a *addressController) ReadAddress(c echo.Context) error {
	ab, writeErr := service.AddressService.ReadAddress()
	if writeErr != nil {
		return c.JSON(writeErr.Status(), writeErr)
	}
	return c.JSON(http.StatusOK, ab)
}
