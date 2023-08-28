package models

import (
	"KevinCatering/db"
	str "KevinCatering/struct"
	"KevinCatering/tools"
	"fmt"
	"net/http"
	"strconv"
)

//Generate_id_catering
func Generate_Id_Catering() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_catering FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_catering=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

//Input_Catering
func Input_Catering(id_user string, nama_catering string, alamat_catering string, telp_catering string,
	email_catering string, deskripsi_catering string, tipe_catering string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Catering()

	nm_str := strconv.Itoa(nm)

	id_ct := "CT-" + nm_str

	sqlStatement := "INSERT INTO catering (id_catering, id_user, nama_catering, alamat_catering, telp_catering, email_catering, deskripsi_catering, tipe_pemesanan_catering, foto_profil_catering, rating) values(?,?,?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_ct, id_user, nama_catering, alamat_catering, telp_catering, email_catering,
		deskripsi_catering, tipe_catering, "uploads/images.png", 0.0)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//see profile
func Read_Profile_Catering(id_user string) (tools.Response, error) {
	var res tools.Response
	var arr []str.Read_Catering
	var obj str.Read_Catering
	var obj_c str.Tipe_Catering

	con := db.CreateCon()

	sqlStatement := "SELECT id_catering, nama_catering, alamat_catering, telp_catering, email_catering, deskripsi_catering, tipe_pemesanan_catering, foto_profil_catering,rating FROM catering WHERE id_user=?"

	rows, err := con.Query(sqlStatement, id_user)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_catering, &obj.Nama_catering, &obj.Alamat_catering,
			&obj.Telp_catering, &obj.Email_catering, &obj.Deskripsi_catering,
			&obj_c.Tipe_catering, &obj.Foto_profil_catering, &obj.Rating)
		obj.Tipe_pemesanan = tools.String_Separator_To_String(obj_c.Tipe_catering)
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

//see_catering
func Read_Catering(id_catering string) (tools.Response, error) {
	var res tools.Response
	var arr []str.Read_Catering
	var obj str.Read_Catering
	var obj_c str.Tipe_Catering

	con := db.CreateCon()

	sqlStatement := "SELECT id_catering, nama_catering, alamat_catering, telp_catering, email_catering, deskripsi_catering, tipe_pemesanan_catering, foto_profil_catering,rating FROM catering WHERE id_catering=?"

	rows, err := con.Query(sqlStatement, id_catering)

	fmt.Println(id_catering)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_catering, &obj.Nama_catering, &obj.Alamat_catering,
			&obj.Telp_catering, &obj.Email_catering, &obj.Deskripsi_catering,
			&obj_c.Tipe_catering, &obj.Foto_profil_catering, &obj.Rating)
		obj.Tipe_pemesanan = tools.String_Separator_To_String(obj_c.Tipe_catering)
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

//read buat awal
func Read_Awal_Catering() (tools.Response, error) {
	var res tools.Response
	var arr []str.Read_Awal_Catering
	var obj str.Read_Awal_Catering
	var obj_c str.Tipe_Catering

	con := db.CreateCon()

	sqlStatement := "SELECT id_catering, nama_catering, alamat_catering,tipe_pemesanan_catering,rating FROM catering "

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_catering, &obj.Nama_catering, &obj.Alamat_catering,
			&obj_c.Tipe_catering, &obj.Rating)
		obj.Tipe_pemesanan = tools.String_Separator_To_String(obj_c.Tipe_catering)
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

//edit profile
func Edit_Profile_Catering(id_catering string, nama_catering string, alamat_catering string,
	telp_catering string, email_catering string, deskripsi_catering string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "UPDATE catering SET nama_catering=?,alamat_catering=?,telp_catering=?,email_catering=?,deskripsi_catering=? WHERE id_catering=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama_catering, alamat_catering, telp_catering, email_catering, deskripsi_catering, id_catering)

	if err != nil {
		return res, err
	}

	rowschanged, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": rowschanged,
	}

	return res, nil
}
