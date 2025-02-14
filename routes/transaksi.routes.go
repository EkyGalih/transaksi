package routes

import (
	controllertransaksi "github.com/EkyGalih/transaksi/controller/controller_transaksi"
	"github.com/labstack/echo"
)

func TransaksiRoutes(e *echo.Echo) {
	e.GET("/transaksi", controllertransaksi.Index)
	e.GET("/transaksi/:id", controllertransaksi.GetDetailTransaksi)
	e.POST("/transaksi", controllertransaksi.StoreTransaksi)
	e.PUT("/transaksi/:id", controllertransaksi.UpdateTransaksi)
	e.DELETE("/transaksi/:id", controllertransaksi.DeleteTransaksi)
}