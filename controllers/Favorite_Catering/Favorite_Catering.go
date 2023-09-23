package Favorite_Catering

import (
	"KevinCatering/models/Favorite_Catering"
	"github.com/labstack/echo/v4"
	"net/http"
)

func InputFavoriteCatering(c echo.Context) error {
	id_user := c.FormValue("id_user")
	id_catering := c.FormValue("id_catering")

	result, err := Favorite_Catering.Input_Favorite_Catering(id_user, id_catering)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
