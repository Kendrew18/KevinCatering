package controllers

import (
	"KevinCatering/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//Input Realisasi
func InputRealisasi(c echo.Context) error {
	id_bahan_menu := c.FormValue("id_bahan_menu")
	keterangan := c.FormValue("keterangan")
	jumlah := c.FormValue("jumlah")
	harga := c.FormValue("harga")

	jmlh, _ := strconv.ParseFloat(jumlah, 64)
	hrg, _ := strconv.Atoi(harga)

	result, err := models.Input_Realisasi(id_bahan_menu, keterangan, jmlh, hrg)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read Realisasi
func ReadRealisasi(c echo.Context) error {
	id_bahan_menu := c.FormValue("id_bahan_menu")

	result, err := models.Read_Realisasi(id_bahan_menu)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
