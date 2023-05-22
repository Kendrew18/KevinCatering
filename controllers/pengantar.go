package controllers

import (
	"KevinCatering/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//Sign-up-pengantar
func SignUpPengantar(c echo.Context) error {
	id_catering := c.FormValue("id_catering")
	nama_user := c.FormValue("nama_user")
	telp_user := c.FormValue("telp_user")
	email_user := c.FormValue("email_user")
	username_user := c.FormValue("username_user")
	password_user := c.FormValue("password_user")
	status_user := c.FormValue("status_user")

	stu, _ := strconv.Atoi(status_user)

	result, err := models.Sign_Up_Pengantar(id_catering, nama_user, telp_user, email_user, username_user, password_user, stu)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
func ReadPengantar(c echo.Context) error {
	id_catering := c.FormValue("id_catering")

	result, err := models.Read_Pengantar(id_catering)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
