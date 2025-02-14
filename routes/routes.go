package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is web api transaksi!")
	})

	TransaksiRoutes(e)

	return e
}
