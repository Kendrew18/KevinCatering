package controllers

import (
	"KevinCatering/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

//InputCatering
func InputCatering(c echo.Context) error {
	id_user := c.FormValue("id_user")
	nama_catering := c.FormValue("nama_catering")
	alamat_catering := c.FormValue("alamat_catering")
	telp_catering := c.FormValue("telp_catering")
	email_catering := c.FormValue("email_catering")
	deskripsi_catering := c.FormValue("deskripsi_catering")
	tipe_pemesanan := c.FormValue("tipe_pemesanan")

	result, err := models.Input_Catering(id_user, nama_catering, alamat_catering, telp_catering, email_catering, deskripsi_catering, tipe_pemesanan, c.Response(), c.Request())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Profile_Catering
func ReadProfileCatering(c echo.Context) error {
	id_user := c.FormValue("id_user")

	result, err := models.Read_Profile_Catering(id_user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Catering
func ReadCatering(c echo.Context) error {

	result, err := models.Read_Catering()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Edit_Profile_Catering
func EditProfileCatering(c echo.Context) error {
	id_catering := c.FormValue("id_catering")
	nama_catering := c.FormValue("nama_catering")
	alamat_catering := c.FormValue("alamat_catering")
	telp_catering := c.FormValue("telp_catering")
	email_catering := c.FormValue("email_catering")
	deskripsi_catering := c.FormValue("deskripsi_catering")

	result, err := models.Edit_Profile_Catering(id_catering, nama_catering, alamat_catering, telp_catering, email_catering, deskripsi_catering)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Get_QR_Catering
func GetQRCatering(c echo.Context) error {
	id_catering := c.FormValue("id_catering")

	result, err := models.Get_QR_Catering(id_catering)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
