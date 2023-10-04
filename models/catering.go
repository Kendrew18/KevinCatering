package models

import (
	"KevinCatering/db"
	_struct "KevinCatering/struct"
	"KevinCatering/struct/Catering"
	"KevinCatering/tools"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//Input_Catering
func Input_Catering(id_user string, nama_catering string, alamat_catering string, telp_catering string, email_catering string, deskripsi_catering string, tipe_catering string, writer http.ResponseWriter, request *http.Request) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM catering ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_ct := "CT-" + strconv.Itoa(nm_str)

	sqlStatement := "INSERT INTO catering (co, id_catering, id_user, nama_catering, alamat_catering, telp_catering, email_catering, deskripsi_catering, tipe_pemesanan_catering, foto_profil_catering, rating,path_foto_qr) values(?,?,?,?,?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_ct, id_user, nama_catering, alamat_catering, telp_catering, email_catering, deskripsi_catering, tipe_catering, "uploads/images.png", 0.0, "")

	request.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err := request.FormFile("photo")
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	defer file.Close()

	fmt.Println("File Info")
	fmt.Println("File Name : ", handler.Filename)
	fmt.Println("File Size : ", handler.Size)
	fmt.Println("File Type : ", handler.Header.Get("Content-Type"))

	var tempFile *os.File
	path := ""

	if strings.Contains(handler.Filename, "jpg") {
		path = "foto_qr_catering/" + id_ct + ".jpg"
		tempFile, err = ioutil.TempFile("foto_qr_catering/", "Read"+"*.jpg")
	}
	if strings.Contains(handler.Filename, "jpeg") {
		path = "foto_qr_catering/" + id_ct + ".jpeg"
		tempFile, err = ioutil.TempFile("foto_qr_catering/", "Read"+"*.jpeg")
	}
	if strings.Contains(handler.Filename, "png") {
		path = "foto_qr_catering/" + id_ct + ".png"
		tempFile, err = ioutil.TempFile("foto_qr_catering/", "Read"+"*.png")
	}

	if err != nil {
		return res, err
	}

	fileBytes, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		return res, err2
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return res, err
	}

	fmt.Println("Success!!")
	fmt.Println(tempFile.Name())
	tempFile.Close()

	err = os.Rename(tempFile.Name(), path)
	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	fmt.Println("new path:", tempFile.Name())

	sqlstatement := "UPDATE catering SET path_foto_qr=? WHERE id_catering=?"

	stmt, err = con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(path, id_ct)

	if err != nil {
		return res, err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return res, err
	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//see profile
func Read_Profile_Catering(id_user string) (tools.Response, error) {
	var res tools.Response
	var arr []Catering.Read_Catering
	var obj Catering.Read_Catering
	var obj_c Catering.Tipe_Catering

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
func Read_Catering() (tools.Response, error) {
	var res tools.Response
	var arr []Catering.Read_Catering
	var obj Catering.Read_Catering
	var obj_c Catering.Tipe_Catering

	con := db.CreateCon()

	sqlStatement := "SELECT catering.id_catering, nama_catering, alamat_catering, telp_catering, email_catering, deskripsi_catering, tipe_pemesanan_catering, foto_profil_catering,rating,longtitude,langtitude,radius FROM catering JOIN maps m on catering.id_catering = m.id_catering"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_catering, &obj.Nama_catering, &obj.Alamat_catering,
			&obj.Telp_catering, &obj.Email_catering, &obj.Deskripsi_catering,
			&obj_c.Tipe_catering, &obj.Foto_profil_catering, &obj.Rating,
			&obj.Longtitude, &obj.Langtitude, &obj.Radius)
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
func Edit_Profile_Catering(id_catering string, nama_catering string, alamat_catering string, telp_catering string, email_catering string, deskripsi_catering string) (tools.Response, error) {
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

//Get_QR_Catering
func Get_QR_Catering(id_catering string) (tools.Response, error) {
	var res tools.Response

	var obj _struct.Read_Path_Foto_QR

	con := db.CreateCon()

	sqlstatement := "SELECT path_foto_qr FROM catering WHERE id_catering=?"

	err := con.QueryRow(sqlstatement, id_catering).Scan(&obj.Path_QR)

	obj.Id_catering = id_catering

	if err != nil {
		fmt.Println(obj.Path_QR)
		return res, err
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = obj
	}

	return res, nil

}

//Set Favorite Catering
func Set_Favorite_Catering(id_user string, id_catering string) (tools.Response, error) {
	var res tools.Response

	return res, nil
}

//Filter_Catering
func Filter_Catering(tipe int, id_user string) (tools.Response, error) {
	var res tools.Response
	var arr []Catering.Read_Catering
	var obj Catering.Read_Catering
	var obj_c Catering.Tipe_Catering

	con := db.CreateCon()

	sqlStatement := ""

	rows, err := con.Query(sqlStatement)

	if tipe == 0 {
		sqlStatement = "SELECT catering.id_catering, nama_catering, alamat_catering, telp_catering, email_catering, deskripsi_catering, tipe_pemesanan_catering, foto_profil_catering,rating,longtitude,langtitude,radius FROM catering JOIN maps m on catering.id_catering = m.id_catering"

		rows, err = con.Query(sqlStatement)
	} else if tipe == 1 {
		sqlStatement = "SELECT catering.id_catering, nama_catering, alamat_catering, telp_catering, email_catering, deskripsi_catering, tipe_pemesanan_catering, foto_profil_catering,rating,longtitude,langtitude,radius FROM catering JOIN maps m on catering.id_catering = m.id_catering ORDER BY rating DESC "

		rows, err = con.Query(sqlStatement)
	} else if tipe == 2 {
		sqlStatement = "SELECT c.id_catering, nama_catering, alamat_catering, telp_catering, email_catering, deskripsi_catering, tipe_pemesanan_catering, foto_profil_catering,rating,longtitude,langtitude,radius FROM favorite_catering JOIN catering c on c.id_catering = favorite_catering.id_catering JOIN maps m on favorite_catering.id_catering = m.id_catering WHERE favorite_catering.id_user=? ORDER BY favorite_catering.co ASC"

		rows, err = con.Query(sqlStatement, id_user)
	} else if tipe == 3 {
		sqlStatement = "SELECT catering.id_catering, nama_catering, alamat_catering, telp_catering, email_catering, deskripsi_catering, tipe_pemesanan_catering, foto_profil_catering,rating,longtitude,langtitude,radius FROM catering JOIN maps m on catering.id_catering = m.id_catering WHERE tipe_pemesanan_catering LIKE '%Harian%' ORDER BY catering.co"

		rows, err = con.Query(sqlStatement)
	} else if tipe == 4 {
		sqlStatement = "SELECT catering.id_catering, nama_catering, alamat_catering, telp_catering, email_catering, deskripsi_catering, tipe_pemesanan_catering, foto_profil_catering,rating,longtitude,langtitude,radius FROM catering JOIN maps m on catering.id_catering = m.id_catering WHERE tipe_pemesanan_catering LIKE '%Mingguan%' ORDER BY catering.co"

		rows, err = con.Query(sqlStatement)
	} else if tipe == 5 {
		sqlStatement = "SELECT catering.id_catering, nama_catering, alamat_catering, telp_catering, email_catering, deskripsi_catering, tipe_pemesanan_catering, foto_profil_catering,rating,longtitude,langtitude,radius FROM catering JOIN maps m on catering.id_catering = m.id_catering WHERE tipe_pemesanan_catering LIKE '%Bulanan%' ORDER BY catering.co ASC"

		rows, err = con.Query(sqlStatement)
	}

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_catering, &obj.Nama_catering, &obj.Alamat_catering,
			&obj.Telp_catering, &obj.Email_catering, &obj.Deskripsi_catering,
			&obj_c.Tipe_catering, &obj.Foto_profil_catering, &obj.Rating,
			&obj.Longtitude, &obj.Langtitude, &obj.Radius)

		id_favorite := ""
		sqlstatement := "SELECT id_favorite_catering FROM favorite_catering WHERE id_catering=? && id_user=?"

		_ = con.QueryRow(sqlstatement, obj.Id_catering, id_user).Scan(&id_favorite)

		if id_favorite == "" {
			obj.Favorite = 0
		} else {
			obj.Favorite = 1
		}

		obj.Tipe_pemesanan = tools.String_Separator_To_String(obj_c.Tipe_catering)
		if err != nil {
			return res, err
		}
		arr = append(arr, obj)
	}

	if arr == nil {
		arr = append(arr, obj)
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
