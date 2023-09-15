package controllers

import (
	"KevinCatering/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ReadPembayaran(c echo.Context) error {
	id_order := c.FormValue("id_order")

	result, err := models.Read_Pembayaran(id_order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UploadFotoPembayaran(c echo.Context) error {
	id_order := c.FormValue("id_order")

	result, err := models.Upload_Foto_Pembayaran(id_order, c.Response(), c.Request())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Order_Recipe
func ReadOrder(c echo.Context) error {
	id_user := c.FormValue("id_user")

	result, err := models.Read_Order_Recipe(id_user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Detail_Order_Recipe
func ReadDetailOrder(c echo.Context) error {
	id_order := c.FormValue("id_order")

	result, err := models.Read_Detail_Order_Recipe(id_order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Confirm_Pembayaran
func ConfirmPembayaran(c echo.Context) error {
	id_order := c.FormValue("id_order")

	result, err := models.Confirm_Pembayaran(id_order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadFoto(c echo.Context) error {
	path := c.FormValue("path")
	return c.File(path)
}
