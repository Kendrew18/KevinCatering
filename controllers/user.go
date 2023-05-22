package controllers

import (
	"KevinCatering/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//Sign_up
func SignUP(c echo.Context) error {
	nama_user := c.FormValue("nama_user")
	telp_user := c.FormValue("telp_user")
	email_user := c.FormValue("email_user")
	username_user := c.FormValue("username_user")
	password_user := c.FormValue("password_user")
	status_user := c.FormValue("status_user")

	stu, _ := strconv.Atoi(status_user)

	result, err := models.Sign_up(nama_user, telp_user, email_user, username_user, password_user, stu)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//login
func LoginUser(c echo.Context) error {
	username := c.FormValue("username_user")
	password := c.FormValue("password_user")

	result, err := models.Login(username, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Profile
func Read_Profile(c echo.Context) error {
	id_user := c.FormValue("id_user")

	result, err := models.Read_Profile(id_user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Edit_Profile
func EditProfile(c echo.Context) error {
	id_user := c.FormValue("id_user")
	nama_user := c.FormValue("nama_user")
	telp_user := c.FormValue("telp_user")
	email_user := c.FormValue("email_user")

	result, err := models.Edit_Profile(id_user, nama_user, telp_user, email_user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
