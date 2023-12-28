package Master_Menu

import (
	"KevinCatering/models/Master_Menu"
	"github.com/labstack/echo/v4"
	"net/http"
)

func InputMasterMenu(c echo.Context) error {
	id_catering := c.FormValue("id_catering")
	nama_menu := c.FormValue("nama_menu")
	deskripsi_menu := c.FormValue("deskripsi_menu")
	nama_bahan := c.FormValue("nama_bahan")
	jumlah_bahan := c.FormValue("jumlah_bahan")
	satuan_bahan := c.FormValue("satuan_bahan")
	harga_bahan := c.FormValue("harga_bahan")

	result, err := Master_Menu.Input_Master_Menu(id_catering, nama_menu, deskripsi_menu, nama_bahan, jumlah_bahan, satuan_bahan, harga_bahan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadMasterMenu(c echo.Context) error {
	id_catering := c.FormValue("id_catering")

	result, err := Master_Menu.Read_Master_Menu(id_catering)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DropDownMasterMenu(c echo.Context) error {
	id_catering := c.FormValue("id_catering")

	result, err := Master_Menu.Drop_Down_Master_Menu(id_catering)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
