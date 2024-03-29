package controllers

import (
	"KevinCatering/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//Input_Order
func InputOrder(c echo.Context) error {
	id_user := c.FormValue("id_user")
	id_catering := c.FormValue("id_catering")
	id_menu := c.FormValue("id_menu")
	nama_menu := c.FormValue("nama_menu")
	jumlah_menu := c.FormValue("jumlah_menu")
	harga_menu := c.FormValue("harga_menu")
	tanggal_menu := c.FormValue("tanggal_menu")
	tanggal_order := c.FormValue("tanggal_order")
	alamat := c.FormValue("alamat")
	longtitude := c.FormValue("longtitude")
	langtitude := c.FormValue("langtitude")

	long, _ := strconv.ParseFloat(longtitude, 64)
	lang, _ := strconv.ParseFloat(langtitude, 64)

	result, err := models.Input_Order(id_catering, id_user, id_menu, nama_menu, jumlah_menu, harga_menu, tanggal_menu, tanggal_order, alamat, lang, long, c.Response(), c.Request())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Show_Order_Menu
func ShowOrderMenu(c echo.Context) error {
	id := c.FormValue("id")
	tanggal_menu := c.FormValue("tanggal_menu")

	result, err := models.Show_Order_Menu(id, tanggal_menu)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Set_Pengantar
func SetPegantar(c echo.Context) error {
	id_detail_order := c.FormValue("id_detail_order")
	id_pengantar := c.FormValue("id_pengantar")

	result, err := models.Set_Pegantar(id_detail_order, id_pengantar)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Confirm_Order
func ConfirmOrder(c echo.Context) error {
	id := c.FormValue("id")
	id_detail_order := c.FormValue("id_detail_order")

	result, err := models.Confirm_Order(id, id_detail_order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Order_Detail_User
func OrderDetailUser(c echo.Context) error {
	id_detail_order := c.FormValue("id_detail_order")

	result, err := models.Order_Detail_User(id_detail_order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//History Order
func HistoryOrder(c echo.Context) error {
	id_user := c.FormValue("id_user")
	tanggal_menu := c.FormValue("tanggal_menu")

	result, err := models.History_Order(id_user, tanggal_menu)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read Order Menu Pengantar
func ReadOrderMenuPengantar(c echo.Context) error {
	id_pengantar := c.FormValue("id_pengantar")

	result, err := models.Read_Order_Pengantar(id_pengantar)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

// Read Location User
func ReadLocationUser(c echo.Context) error {
	id_detail_order := c.FormValue("id_detail_order")

	result, err := models.Read_Location_User(id_detail_order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Filter Order Menu
func FilterOrderMenu(c echo.Context) error {
	id := c.FormValue("id")
	tanggal_menu := c.FormValue("tanggal_menu")

	result, err := models.Filter_Order_Menu(id, tanggal_menu)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

//Filter History Order
func FilterHistoryOrder(c echo.Context) error {
	id_user := c.FormValue("id_user")
	tanggal_menu := c.FormValue("tanggal_menu")

	result, err := models.Filter_History_Order(id_user, tanggal_menu)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
