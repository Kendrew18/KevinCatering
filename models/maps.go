package models

import (
	"KevinCatering/db"
	"KevinCatering/struct/Maps"
	"KevinCatering/tools"
	"net/http"
	"strconv"
)

//input_maps
func Input_Maps(Id_catering string, longtitude float64, langtitude float64, radius int) (tools.Response, error) {
	var res tools.Response
	var maps Maps.Read_maps

	con := db.CreateCon()

	sqlstatement := "SELECT id_catering FROM maps WHERE id_catering=?"

	_ = con.QueryRow(sqlstatement, Id_catering).Scan(&maps.Id_catering)

	if maps.Id_catering == "" {

		nm_str := 0

		Sqlstatement := "SELECT co FROM maps ORDER BY co DESC Limit 1"

		_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

		nm_str = nm_str + 1

		id_ct := "MP-" + strconv.Itoa(nm_str)

		sqlStatement := "INSERT INTO maps (co, id_maps, id_catering, longtitude, langtitude, radius) values(?,?,?,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(nm_str, id_ct, Id_catering, longtitude, langtitude, radius)

		stmt.Close()

		res.Status = http.StatusOK
		res.Message = "Sukses"

	} else {

		sqlstatement := "UPDATE maps SET longtitude=?,langtitude=?,radius=? WHERE id_catering=?"

		stmt, err := con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(longtitude, langtitude, radius, Id_catering)

		if err != nil {
			return res, err
		}

		res.Status = http.StatusOK
		res.Message = "Suksess"
	}

	return res, nil
}

//Read_Maps
func Read_Maps(Id_catering string) (tools.Response, error) {
	var res tools.Response
	var maps Maps.Read_maps
	var maps_arr []Maps.Read_maps

	con := db.CreateCon()

	sqlstatement := "SELECT id_maps, id_catering, longtitude, langtitude, radius FROM maps WHERE id_catering=?"

	err := con.QueryRow(sqlstatement, Id_catering).Scan(&maps.Id_maps,
		&maps.Id_catering, &maps.Longtitude, &maps.Langtitude, &maps.Radius)

	if err != nil {
		return res, err
	}

	maps_arr = append(maps_arr, maps)

	if maps_arr == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = maps_arr
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = maps_arr
	}

	return res, nil
}
