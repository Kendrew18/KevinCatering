package Notif

import (
	"KevinCatering/models/Notif"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ShowAllNotif(c echo.Context) error {
	id_catering := c.FormValue("id_catering")

	result, err := Notif.Show_All_Notif(id_catering)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Notif Pembayaran
func ReadDetailNotif(c echo.Context) error {
	id_order := c.FormValue("id_order")

	result, err := Notif.Read_Detail_Notif(id_order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
