package models

import (
	"KevinCatering/db"
	str "KevinCatering/struct"
	"KevinCatering/tools"
	"net/http"
	"strconv"
)

func Generate_Id_Pengantar() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_pengantar FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_pengantar=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no

}

func Sign_Up_Pengantar(id_catering string, nama_user string, telp_user string, email_user string,
	username_user string, password_user string,
	status_user int) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_User()

	nm_str := strconv.Itoa(nm)

	id_US := "US-" + nm_str

	nmp := Generate_Id_Pengantar()

	nm_strp := strconv.Itoa(nmp)

	id_USP := "USP-" + nm_strp

	sqlStatement := "INSERT INTO user (id_user,nama,telp_user,email_user,username_user,password_user,foto_user,status_user) values(?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_US, nama_user, telp_user, email_user, username_user, password_user, "uploads/images.png", status_user)

	sqlStatement = "INSERT INTO pengantar (id_pengantar, id_user, id_catering) values(?,?,?)"

	stmt, err = con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_USP, id_US, id_catering)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

func Read_Pengantar(id_catering string) (tools.Response, error) {
	var res tools.Response
	var arr []str.Read_pengantar
	var obj str.Read_pengantar

	con := db.CreateCon()

	sqlStatement := "SELECT id_pengantar,id_catering,u.nama,u.telp_user FROM pengantar JOIN user u on pengantar.id_user = u.id_user WHERE id_catering=?"

	rows, err := con.Query(sqlStatement, id_catering)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_pengantar, &obj.Id_catering, &obj.Nama_Pengantar, &obj.Nomor_telp_pengantar)
		if err != nil {
			return res, err
		}
		arr = append(arr, obj)
	}

	if arr == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr
	}

	return res, nil
}

func Update_Maps(id_user string, langtitude float64, longtitude float64) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "UPDATE pengantar SET longtitude=?,langtitude=? WHERE id_user=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(longtitude, langtitude, id_user)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"

	return res, nil
}

func Read_Maps_Pengantar(id_user string) (tools.Response, error) {
	var res tools.Response
	var maps str.Read_Maps_Pengantar
	var maps_arr []str.Read_Maps_Pengantar

	con := db.CreateCon()

	sqlstatement := "SELECT longtitude, langtitude FROM pengantar WHERE id_user=?"

	err := con.QueryRow(sqlstatement, id_user).Scan(&maps.Longtitude, &maps.Langtitude)

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
