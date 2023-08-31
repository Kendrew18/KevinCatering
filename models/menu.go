package models

import (
	"KevinCatering/db"
	str "KevinCatering/struct"
	"KevinCatering/tools"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Generate_Id_Menu() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_menu FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_menu=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

//input menu
func Input_Menu(id_catering string, nama_menu string, harga_menu int64, tanggal_menu string,
	jam_pengiriman_awal string, jam_pengiriman_akhir string, status int) (tools.Response, error) {
	var res tools.Response
	var id str.Read_Id_Menu
	con := db.CreateCon()

	nm := Generate_Id_Menu()

	nm_str := strconv.Itoa(nm)

	id_MN := "MN-" + nm_str

	date, _ := time.Parse("02-01-2006", tanggal_menu)
	date_sql := date.Format("2006-01-02")

	sqlStatement := "INSERT INTO menu (id_catering, id_menu, nama_menu, harga_menu, tanggal_menu, jam_pengiriman_awal, jam_pengiriman_akhir, status_menu,foto_menu) values(?,?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_catering, id_MN, nama_menu, harga_menu, date_sql, jam_pengiriman_awal,
		jam_pengiriman_akhir, status, "photo_menu/photo.jpg")

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
	var obj str.Read_Menu_fix
	var arr []str.Read_Menu_fix

	con := db.CreateCon()

	if tanggal_menu2 == "" {

		date, _ := time.Parse("02-01-2006", tanggal_menu)
		date_sql := date.Format("2006-01-02")

		sqlStatement := "SELECT id_menu, nama_menu, harga_menu, jam_pengiriman_awal, jam_pengiriman_akhir, status_menu, foto_menu FROM menu WHERE tanggal_menu=? && id_catering=? "

		rows, err := con.Query(sqlStatement, date_sql, id_catering)

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
			arr = append(arr, obj)
		}

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

			sqlStatement := "SELECT id_menu, nama_menu, harga_menu, jam_pengiriman_awal, jam_pengiriman_akhir, status_menu, foto_menu FROM menu WHERE tanggal_menu=? && id_catering=? "

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
				arr = append(arr, obj)
			}

		}
	}
	fmt.Println(arr)
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

//read_menu_2_tgl
/*func Read_Menu_2_tgl(id_catering string, tanggal_menu string, tanggal_menu2 string) (tools.Response, error) {
	var res tools.Response
	var rm str.Read_Menu
	var obj str.Read_Menu_fix
	var menu str.Read_Menu_String
	var arr []str.Read_Menu

	con := db.CreateCon()

	date, _ := time.Parse("02-01-2006", tanggal_menu)
	date_sql := date.Format("2006-01-02")

	date2, _ := time.Parse("02-01-2006", tanggal_menu2)
	date_sql2 := date2.Format("2006-01-02")

	sqlStatement := "SELECT id_catering,tanggal_menu, id_menu, nama_menu, harga_menu, jam_pengiriman_awal, jam_pengiriman_akhir, status_menu, foto_menu FROM menu WHERE tanggal_menu>=? && tanggal_menu<=? && id_catering=? "

	err := con.QueryRow(sqlStatement, date_sql, date_sql2, id_catering).Scan(&rm.Id_catering, &rm.Tanggal_menu,
		&menu.Id_menu, &menu.Nama_menu, &menu.Harga_menu, &menu.Jam_pengiriman_awal,
		&menu.Jam_pengiriman_akhir, &menu.Status_menu, &menu.Foto_menu)

	if err != nil {
		return res, err
	}

	id_mn_all := tools.String_Separator_To_String(menu.Id_menu)
	nm_mn := tools.String_Separator_To_String(menu.Nama_menu)
	hg_mn := tools.String_Separator_To_Int64(menu.Harga_menu)
	j_awal := tools.String_Separator_To_String(menu.Jam_pengiriman_awal)
	j_akhir := tools.String_Separator_To_String(menu.Jam_pengiriman_akhir)
	st := tools.String_Separator_To_Int(menu.Status_menu)
	fm := tools.String_Separator_To_String(menu.Foto_menu)

	for i := 0; i < len(id_mn_all); i++ {
		obj.Id_menu = id_mn_all[i]
		obj.Nama_menu = nm_mn[i]
		obj.Harga_menu = hg_mn[i]
		obj.Jam_pengiriman_awal = j_awal[i]
		obj.Jam_pengiriman_akhir = j_akhir[i]
		obj.Status_menu = st[i]
		obj.Foto_menu = fm[i]
		rm.Menu = append(rm.Menu, obj)
	}
	arr = append(arr, rm)

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
*/
//edit menu
func Edit_Menu(id_catering string, id_menu string, nama_menu string, harga_menu int64,
	jam_pengiriman_awal string, jam_pengiriman_akhir string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	condition := ""

	sqlStatement := "SELECT id_detail_order FROM Detail_Order WHERE id_menu=? "

	_ = con.QueryRow(sqlStatement, id_menu).Scan(&condition)

	if condition == "" {

		sqlstatement := "UPDATE menu SET nama_menu=?,harga_menu=?,jam_pengiriman_awal=?,jam_pengiriman_akhir=? WHERE id_catering=? && id_menu=?"

		stmt, err := con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(nama_menu, harga_menu, jam_pengiriman_awal, jam_pengiriman_akhir, id_catering, id_menu)

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

	sqlStatement := "SELECT id_detail_order FROM Detail_Order WHERE id_menu=? "

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

//upload foto
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
