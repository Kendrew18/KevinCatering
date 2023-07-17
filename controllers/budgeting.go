package controllers

import (
	"KevinCatering/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//InputCatering
func InputBudgeting(c echo.Context) error {
	id_catering := c.FormValue("id_catering")
	nama_menu := c.FormValue("nama_menu")
	total_porsi := c.FormValue("total_porsi")
	tanggal_budgeting := c.FormValue("tanggal_budgeting")
	nama_bahan := c.FormValue("nama_bahan")
	jumlah_bahan := c.FormValue("jumlah_bahan")
	satuan_bahan := c.FormValue("satuan_bahan")
	harga_bahan := c.FormValue("harga_bahan")

	tp, _ := strconv.Atoi(total_porsi)

	result, err := models.Input_Budgeting(id_catering, nama_menu, tp, tanggal_budgeting,
		nama_bahan, jumlah_bahan, satuan_bahan, harga_bahan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Profile_Catering
func ReadAwalBudgeting(c echo.Context) error {
	id_catering := c.FormValue("id_catering")

	result, err := models.Read_Budgeting_Awal(id_catering)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Catering
func ReadBudgeting(c echo.Context) error {
	id_budgeting := c.FormValue("id_budgeting")

	result, err := models.Read_Budgeting(id_budgeting)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
