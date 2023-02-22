package admin

import (
	"github.com/labstack/echo/v4"
)

func saveOptionsEndpoint(c echo.Context) error {
	var options []Option
	err := c.Bind(&options)
	if err != nil {
		return err
	}
	return saveOptions(options)
}

func getOptionsEndpoint(c echo.Context) error {
	return c.JSON(200, getOptions())
}
