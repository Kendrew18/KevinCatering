package controllers

import (
	"KevinCatering/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

//Read_Pembayaran
func ReadPembayaran(c echo.Context) error {
	id_order := c.FormValue("id_order")

	result, err := models.Read_Pembayaran(id_order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Order_Recipe
func ReadOrderRecipe(c echo.Context) error {
	id := c.FormValue("id")
	tanggal_recipe := c.FormValue("tanggal_recipe")

	result, err := models.Read_Order_Recipe(id, tanggal_recipe)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Detail_Order_Recipe
func ReadDetailOrderRecipe(c echo.Context) error {
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
