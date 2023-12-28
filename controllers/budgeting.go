package controllers

import (
	"KevinCatering/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//Input_Budgeting
func InputBudgeting(c echo.Context) error {
	id_catering := c.FormValue("id_catering")
	id_master_menu := c.FormValue("id_master_menu")
	total_porsi := c.FormValue("total_porsi")
	tanggal_budgeting := c.FormValue("tanggal_budgeting")

	tp, _ := strconv.Atoi(total_porsi)

	result, err := models.Input_Budgeting(id_catering, tp, tanggal_budgeting, id_master_menu)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Awal_Budgeting
func ReadAwalBudgeting(c echo.Context) error {
	id_catering := c.FormValue("id_catering")

	result, err := models.Read_List_Budgeting(id_catering)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Budgeting
func ReadBudgeting(c echo.Context) error {
	id_budgeting := c.FormValue("id_budgeting")

	result, err := models.Read_Budgeting(id_budgeting)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Update_Status
func UpdateStatusBudgeting(c echo.Context) error {
	id_budgeting := c.FormValue("id_budgeting")

	result, err := models.Update_Status_Budgeting(id_budgeting)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
