package controllers

import (
	"KevinCatering/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//Input_Maps
func InputMaps(c echo.Context) error {
	id_catering := c.FormValue("id_catering")
	longtitude := c.FormValue("longtitude")
	langtitude := c.FormValue("langtitude")
	radius := c.FormValue("radius")

	long, _ := strconv.ParseFloat(longtitude, 64)
	lang, _ := strconv.ParseFloat(langtitude, 64)
	rad, _ := strconv.Atoi(radius)

	result, err := models.Input_Maps(id_catering, long, lang, rad)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Maps
func ReadMaps(c echo.Context) error {
	id_catering := c.FormValue("id_catering")

	result, err := models.Read_Maps(id_catering)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
