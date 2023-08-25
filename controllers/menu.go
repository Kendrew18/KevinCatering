package controllers

import (
	"KevinCatering/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//input_menu
func InputMenu(c echo.Context) error {
	id_catering := c.FormValue("id_catering")
	nama_menu := c.FormValue("nama_menu")
	harga_menu := c.FormValue("harga_menu")
	tanggal_menu := c.FormValue("tanggal_menu")
	jam_pengiriman_awal := c.FormValue("jam_pengiriman_awal")
	jam_pengiriman_akhir := c.FormValue("jam_pengiriman_akhir")
	status := c.FormValue("status")

	st, _ := strconv.Atoi(status)

	hg, _ := strconv.ParseInt(harga_menu, 10, 64)

	result, err := models.Input_Menu(id_catering, nama_menu, hg, tanggal_menu, jam_pengiriman_awal, jam_pengiriman_akhir, st)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//read_menu
func ReadMenu(c echo.Context) error {
	id_catering := c.FormValue("id_catering")
	tanggal_menu := c.FormValue("tanggal_menu")
	tanggal_menu2 := c.FormValue("tanggal_menu2")

	result, err := models.Read_Menu(id_catering, tanggal_menu, tanggal_menu2)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//edit_menu
func EditMenu(c echo.Context) error {
	id_catering := c.FormValue("id_catering")
	id_menu := c.FormValue("id_menu")
	nama_menu := c.FormValue("nama_menu")
	harga_menu := c.FormValue("harga_menu")
	jam_pengiriman_awal := c.FormValue("jam_pengiriman_awal")
	jam_pengiriman_akhir := c.FormValue("jam_pengiriman_akhir")

	hg, _ := strconv.ParseInt(harga_menu, 10, 64)

	result, err := models.Edit_Menu(id_catering, id_menu, nama_menu, hg, jam_pengiriman_awal, jam_pengiriman_akhir)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//delete_menu
func DeleteMenu(c echo.Context) error {
	id_catering := c.FormValue("id_catering")
	id_menu := c.FormValue("id_menu")

	result, err := models.Delete_Menu(id_catering, id_menu)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//upload_foto_menu
func UploadFotoMenu(c echo.Context) error {
	id_catering := c.FormValue("id_catering")
	id_menu := c.FormValue("id_menu")

	result, err := models.Upload_Foto_Menu(id_catering, id_menu, c.Response(), c.Request())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
