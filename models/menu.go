package models

import (
	"KevinCatering/db"
	str "KevinCatering/struct"
	"KevinCatering/struct/Menu"
	"KevinCatering/tools"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//input menu
func Input_Menu(id_catering string, id_master_menu string, harga_menu int64, tanggal_menu string, jam_pengiriman_awal string, jam_pengiriman_akhir string, status int, writer http.ResponseWriter, request *http.Request) (tools.Response, error) {
	var res tools.Response
	var id Menu.Read_Id_Menu

	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM menu ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_MN := "MN-" + strconv.Itoa(nm_str)

	request.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err := request.FormFile("photo")

	path := ""

	if file != nil {
		defer file.Close()

		fmt.Println("File Info")
		fmt.Println("File Name : ", handler.Filename)
		fmt.Println("File Size : ", handler.Size)
		fmt.Println("File Type : ", handler.Header.Get("Content-Type"))

		var tempFile *os.File

		if strings.Contains(handler.Filename, "jpg") {
			path = "photo_menu/" + id_MN + ".jpg"
			tempFile, err = ioutil.TempFile("photo_menu/", "Read"+"*.jpg")
		}
		if strings.Contains(handler.Filename, "jpeg") {
			path = "photo_menu/" + id_MN + ".jpeg"
			tempFile, err = ioutil.TempFile("photo_menu/", "Read"+"*.jpeg")
		}
		if strings.Contains(handler.Filename, "png") {
			path = "photo_menu/" + id_MN + ".png"
			tempFile, err = ioutil.TempFile("photo_menu/", "Read"+"*.png")
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

	} else {
		path = "photo_menu/photo.jpg"
	}

	date, _ := time.Parse("02-01-2006", tanggal_menu)
	date_sql := date.Format("2006-01-02")

	sqlStatement := "INSERT INTO menu (co, id_menu, id_catering, id_master_menu, harga_menu, tanggal_menu, jam_pengiriman_awal, jam_pengiriman_akhir, status_menu, foto_menu) values(?,?,?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_catering, id_MN, id_master_menu, harga_menu, date_sql, jam_pengiriman_awal, jam_pengiriman_akhir, status, path)

	stmt.Close()

	id.Id_menu = id_MN

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = id

	return res, nil
}

//read_menu
func Read_Menu(id_catering string, tanggal_menu string, tanggal_menu2 string) (tools.Response, error) {
	var res tools.Response
	var obj Menu.Read_Menu_fix
	var arr []Menu.Read_Menu_fix
	var obj_fix Menu.Read_Menu
	var arr_fix []Menu.Read_Menu

	con := db.CreateCon()

	if tanggal_menu2 == "" {

		date, _ := time.Parse("02-01-2006", tanggal_menu)
		date_sql := date.Format("2006-01-02")

		sqlStatement := "SELECT id_menu, nama_menu, harga_menu, jam_pengiriman_awal, jam_pengiriman_akhir, status_menu, foto_menu FROM menu join master_menu mm on mm.id_master_menu = menu.id_master_menu WHERE tanggal_menu=? && menu.id_catering=? "

		rows, err := con.Query(sqlStatement, date_sql, id_catering)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.Id_menu, &obj.Nama_menu, &obj.Harga_menu, &obj.Jam_pengiriman_awal, &obj.Jam_pengiriman_akhir, &obj.Status_menu, &obj.Foto_menu)
			if err != nil {
				return res, err
			}
			arr = append(arr, obj)
		}

		obj_fix.Id_catering = id_catering
		obj_fix.Tanggal_menu = tanggal_menu
		obj_fix.Menu = arr

		arr_fix = append(arr_fix, obj_fix)

	} else if tanggal_menu2 != "" {

		var tgl str.Read_Tanggal
		var arr_tgl []str.Read_Tanggal

		date, _ := time.Parse("02-01-2006", tanggal_menu)
		date_sql := date.Format("2006-01-02")

		date2, _ := time.Parse("02-01-2006", tanggal_menu2)
		date_sql2 := date2.Format("2006-01-02")

		sqlStatement := "SELECT DISTINCT(tanggal_menu) FROM menu WHERE tanggal_menu>=? && tanggal_menu<=? && id_catering=? "

		rows, err := con.Query(sqlStatement, date_sql, date_sql2, id_catering)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&tgl.Tanggal)
			if err != nil {
				return res, err
			}
			arr_tgl = append(arr_tgl, tgl)
		}

		for i := 0; i < len(arr_tgl); i++ {
			var obj Menu.Read_Menu_fix
			var arr2 []Menu.Read_Menu_fix

			sqlStatement := "SELECT id_menu, nama_menu, harga_menu, jam_pengiriman_awal, jam_pengiriman_akhir, status_menu, foto_menu FROM menu join master_menu mm on mm.id_master_menu = menu.id_master_menu WHERE tanggal_menu=? && menu.id_catering=?"

			rows, err := con.Query(sqlStatement, arr_tgl[i].Tanggal, id_catering)

			defer rows.Close()

			if err != nil {
				return res, err
			}

			for rows.Next() {
				err = rows.Scan(&obj.Id_menu, &obj.Nama_menu, &obj.Harga_menu,
					&obj.Jam_pengiriman_awal, &obj.Jam_pengiriman_akhir, &obj.Status_menu, &obj.Foto_menu)
				if err != nil {
					return res, err
				}
				arr2 = append(arr2, obj)
			}

			date_catering, _ := time.Parse("2006-01-02", arr_tgl[i].Tanggal)
			date_catering_fix := date_catering.Format("02-01-2006")

			obj_fix.Id_catering = id_catering
			obj_fix.Tanggal_menu = date_catering_fix
			obj_fix.Menu = arr2

			arr = arr2

			arr_fix = append(arr_fix, obj_fix)

		}
	}

	if arr == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_fix
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_fix
	}

	return res, nil
}

//edit menu
func Edit_Menu(id_catering string, id_menu string, harga_menu int64, jam_pengiriman_awal string, jam_pengiriman_akhir string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	condition := ""

	sqlStatement := "SELECT id_detail_order FROM detail_order WHERE id_menu=? "

	_ = con.QueryRow(sqlStatement, id_menu).Scan(&condition)

	if condition == "" {

		sqlstatement := "UPDATE menu SET harga_menu=?,jam_pengiriman_awal=?,jam_pengiriman_akhir=? WHERE id_catering=? && id_menu=?"

		stmt, err := con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(harga_menu, jam_pengiriman_awal, jam_pengiriman_akhir, id_catering, id_menu)

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
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Tidak Dapat Diupdate"
	}
	return res, nil
}

//delete menu
func Delete_Menu(id_catering string, id_menu string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "SELECT id_detail_order FROM detail_order WHERE id_menu=? "

	condition := ""

	_ = con.QueryRow(sqlStatement, id_menu).Scan(&condition)

	if condition == "" {

		sqlStatement = "DELETE FROM menu WHERE id_menu=? && id_catering=?"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(id_menu, id_catering)

		if err != nil {
			return res, err
		}

		_, err = result.RowsAffected()

		if err != nil {
			return res, err
		}

		stmt.Close()

		res.Status = http.StatusOK
		res.Message = "Suksess"

	} else {
		res.Status = http.StatusNotFound
		res.Message = "Tidak Dapat Diupdate"
	}
	return res, nil
}

//upload foto (buat update foto)
func Upload_Foto_Menu(id_catering string, id_menu string, writer http.ResponseWriter, request *http.Request) (tools.Response, error) {
	var res tools.Response
	var fotm str.Foto_Menu

	con := db.CreateCon()

	request.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err := request.FormFile("photo")
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	defer file.Close()

	sqlStatement := "SELECT foto_menu FROM menu WHERE id_menu=? && id_catering=? "

	err = con.QueryRow(sqlStatement, id_menu, id_catering).Scan(&fotm.Path_Foto)

	fmt.Println(fotm.Path_Foto)

	if err != nil {
		return res, err
	}

	new_path := ""

	if fotm.Path_Foto == "photo_menu/photo.jpg" {

		fmt.Println("File Info")
		fmt.Println("File Name : ", handler.Filename)
		fmt.Println("File Size : ", handler.Size)
		fmt.Println("File Type : ", handler.Header.Get("Content-Type"))

		var tempFile *os.File
		path := ""

		if strings.Contains(handler.Filename, "jpg") {
			path = "photo_menu/" + id_menu + ".jpg"
			tempFile, err = ioutil.TempFile("photo_menu/", "Read"+"*.jpg")
		}
		if strings.Contains(handler.Filename, "jpeg") {
			path = "photo_menu/" + id_menu + ".jpeg"
			tempFile, err = ioutil.TempFile("photo_menu/", "Read"+"*.jpeg")
		}
		if strings.Contains(handler.Filename, "png") {
			path = "photo_menu/" + id_menu + ".png"
			tempFile, err = ioutil.TempFile("photo_menu/", "Read"+"*.png")
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

		new_path = path
	} else {
		_ = os.Remove("./" + fotm.Path_Foto)

		fmt.Println("File Info")
		fmt.Println("File Name : ", handler.Filename)
		fmt.Println("File Size : ", handler.Size)
		fmt.Println("File Type : ", handler.Header.Get("Content-Type"))

		var tempFile *os.File
		path := ""

		if strings.Contains(handler.Filename, "jpg") {
			path = "photo_menu/" + id_menu + ".jpg"
			tempFile, err = ioutil.TempFile("photo_menu/", "Read"+"*.jpg")
		}
		if strings.Contains(handler.Filename, "jpeg") {
			path = "photo_menu/" + id_menu + ".jpeg"
			tempFile, err = ioutil.TempFile("photo_menu/", "Read"+"*.jpeg")
		}
		if strings.Contains(handler.Filename, "png") {
			path = "photo_menu/" + id_menu + ".png"
			tempFile, err = ioutil.TempFile("photo_menu/", "Read"+"*.png")
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

		new_path = path
	}

	sqlstatement := "UPDATE menu SET foto_menu=? WHERE id_catering=? && id_menu=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(new_path, id_catering, id_menu)

	if err != nil {
		return res, err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}
