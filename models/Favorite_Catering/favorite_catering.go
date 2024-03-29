package Favorite_Catering

import (
	"KevinCatering/db"
	"KevinCatering/tools"
	"net/http"
	"strconv"
)

func Input_Favorite_Catering(id_user string, id_catering string) (tools.Response, error) {
	var res tools.Response
	var id_favorite_catering string

	con := db.CreateCon()

	Sqlstatement := "SELECT id_favorite_catering FROM favorite_catering WHERE id_user=? && id_catering=?"

	_ = con.QueryRow(Sqlstatement, id_user, id_catering).Scan(&id_favorite_catering)

	if id_favorite_catering == "" {

		nm_str := 0

		Sqlstatement := "SELECT co FROM favorite_catering ORDER BY co DESC Limit 1"

		_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

		nm_str = nm_str + 1

		id_F_C := "FC-" + strconv.Itoa(nm_str)

		Sqlstatement = "INSERT INTO favorite_catering (co, id_favorite_catering, id_user, id_catering) values(?,?,?,?)"

		stmt, err := con.Prepare(Sqlstatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(nm_str, id_F_C, id_user, id_catering)

		if err != nil {
			return res, err
		}

		stmt.Close()

		res.Status = http.StatusOK
		res.Message = "Sukses"
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Sudah Terdaftar Sebagai Favorite"
	}

	return res, nil
}
