package models

import (
	"KevinCatering/db"
	str "KevinCatering/struct"
	"KevinCatering/tools"
	"net/http"
	"strconv"
)

//Generate_Id_Maps
func Generate_Id_Maps() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_maps FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_maps=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

//input_maps
func Input_Maps(Id_catering string, longtitude float64, langtitude float64, radius int) (tools.Response, error) {
	var res tools.Response
	var maps str.Read_maps

	con := db.CreateCon()

	sqlstatement := "SELECT id_catering FROM maps WHERE id_catering=?"

	_ = con.QueryRow(sqlstatement, Id_catering).Scan(&maps.Id_catering)

	if maps.Id_catering == "" {

		nm := Generate_Id_Maps()

		nm_str := strconv.Itoa(nm)

		id_ct := "MP-" + nm_str

		sqlStatement := "INSERT INTO maps (id_maps, id_catering, longtitude, langtitude, radius) values(?,?,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(id_ct, Id_catering, longtitude, langtitude, radius)

		stmt.Close()

		res.Status = http.StatusOK
		res.Message = "Sukses"

	} else {

		sqlstatement := "UPDATE maps SET longtitude=?,langtitude=?,radius=? WHERE id_catering=?"

		stmt, err := con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(longtitude, langtitude, radius)

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
	var maps str.Read_maps
	var maps_arr []str.Read_maps

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
