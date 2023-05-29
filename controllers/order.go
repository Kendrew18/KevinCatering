package controllers

import (
	"KevinCatering/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func InputOrder(c echo.Context) error {
	id_user := c.FormValue("id_user")
	id_catering := c.FormValue("id_catering")
	id_menu := c.FormValue("id_menu")
	nama_menu := c.FormValue("nama_menu")
	harga_menu := c.FormValue("harga_menu")
	tanggal_menu := c.FormValue("tanggal_menu")
	tanggal_order := c.FormValue("tanggal_order")
	status_order := c.FormValue("status_order")
	longtitude := c.FormValue("longtitude")
	langtitude := c.FormValue("langtitude")

	long, _ := strconv.ParseFloat(longtitude, 64)
	lang, _ := strconv.ParseFloat(langtitude, 64)

	result, err := models.Input_Order(id_catering, id_user, id_menu, nama_menu, harga_menu,
		tanggal_order, tanggal_menu, status_order, lang, long)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadOrder(c echo.Context) error {
	id_user := c.FormValue("id_user")

	result, err := models.Read_Order(id_user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadDetailOrder(c echo.Context) error {
	id_order := c.FormValue("id_order")

	result, err := models.Read_Detail_Order(id_order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
