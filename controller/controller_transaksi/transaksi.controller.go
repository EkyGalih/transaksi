package controllertransaksi

import (
	"net/http"

	"github.com/EkyGalih/transaksi/entities"
	modeltransaksi "github.com/EkyGalih/transaksi/model/model_transaksi"
	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	result, err := modeltransaksi.FetchAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetDetailTransaksi(c echo.Context) error {
	id := c.Param("id")

	res, err := modeltransaksi.Detail(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(res.Status, res)
}

func StoreTransaksi(c echo.Context) error {
	var transaksi entities.Transaction

	if err := c.Bind(&transaksi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := modeltransaksi.Store(transaksi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateTransaksi(c echo.Context) error {
	id := c.Param("id")

	var transaksi entities.Transaction
	if err := c.Bind(&transaksi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid input"})
	}

	res, err := modeltransaksi.Update(id, transaksi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func DeleteTransaksi(c echo.Context) error {
	id := c.Param("id")

	res, err := modeltransaksi.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(res.Status, res)
}
