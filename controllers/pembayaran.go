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

func ReadFoto(c echo.Context) error {
	path := c.FormValue("path")
	return c.File(path)
}
