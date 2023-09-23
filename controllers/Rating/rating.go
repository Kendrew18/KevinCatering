package Rating

import (
	"KevinCatering/models/Rating"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func InputRating(c echo.Context) error {
	id_detail_order := c.FormValue("id_detail_order")
	id_catering := c.FormValue("id_catering")
	rating := c.FormValue("rating")
	review := c.FormValue("review")

	rt, _ := strconv.Atoi(rating)

	result, err := Rating.Input_Rating(id_detail_order, id_catering, rt, review)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
