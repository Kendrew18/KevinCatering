package models

import (
	"KevinCatering/db"
	_struct "KevinCatering/struct"
	"KevinCatering/struct/Order"
	"KevinCatering/tools"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//Input_Order
func Input_Order(id_catering string, id_user string, id_menu string, nama_menu string, jumlah string, harga_menu string, tanggal_menu string, tanggal_order string, alamat string, langtitude float64, longtitude float64, writer http.ResponseWriter, request *http.Request) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM order_catering ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_OD := "OR-" + strconv.Itoa(nm_str)

	fmt.Println(id_OD)

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
		path = "uploads/" + id_OD + ".jpg"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.jpg")
	}
	if strings.Contains(handler.Filename, "jpeg") {
		path = "uploads/" + id_OD + ".jpeg"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.jpeg")
	}
	if strings.Contains(handler.Filename, "png") {
		path = "uploads/" + id_OD + ".png"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.png")
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

	id_M := tools.String_Separator_To_String(id_menu)
	nama_mn := tools.String_Separator_To_String(nama_menu)
	jmlh_mn := tools.String_Separator_To_Int64(jumlah)
	harga_mn := tools.String_Separator_To_Int64(harga_menu)
	tgl_mn := tools.String_Separator_To_String(tanggal_menu)

	date2, _ := time.Parse("02-01-2006", tanggal_order)
	date_sql2 := date2.Format("2006-01-02")

	var total int64

	total = 0

	for i := 0; i < len(id_M); i++ {
		total = total + (jmlh_mn[i] * harga_mn[i])
	}

	sqlStatement := "INSERT INTO order_catering (co, id_order, id_user, id_catering, alamat, total, tanggal_order, longtitude, langtitude) values(?,?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_OD, id_user, id_catering, alamat, total, date_sql2, longtitude, langtitude)

	fmt.Println(id_M)

	for i := 0; i < len(id_M); i++ {
		nm_str2 := 0

		Sqlstatement := "SELECT co FROM detail_order ORDER BY co DESC Limit 1"

		_ = con.QueryRow(Sqlstatement).Scan(&nm_str2)

		nm_str2 = nm_str2 + 1

		id_DO := "DO-" + strconv.Itoa(nm_str2)

		date, _ := time.Parse("02-01-2006", tgl_mn[i])
		date_sql := date.Format("2006-01-02")

		sqlStatement := "INSERT INTO detail_order (co, id_detail_order, id_order, id_menu, nama_menu, tanggal_menu,jumlah, harga_menu, status_order) values(?,?,?,?,?,?,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(nm_str2, id_DO, id_OD, id_M[i], nama_mn[i], date_sql, jmlh_mn[i], harga_mn[i], "In-Progress")

		if err != nil {
			return res, err
		}
	}

	nm_str2 := 0

	Sqlstatement = "SELECT co FROM pembayaran ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str2)

	nm_str2 = nm_str2 + 1

	id_OD2 := "PBR-" + strconv.Itoa(nm_str2)

	sqlStatement = "INSERT INTO pembayaran (co, id_pembayaran,id_order,status_pembayaran,bukti_pembayaran) values(?,?,?,?,?)"

	stmt, err = con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str2, id_OD2, id_OD, 0, path)

	nm_str3 := 0

	Sqlstatement = "SELECT co FROM notif ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str3)

	nm_str3 = nm_str3 + 1

	id_ntf := "NTF-" + strconv.Itoa(nm_str3)

	sqlStatement = "INSERT INTO notif (co, id_notif, id_catering, id_order, catatan) values(?,?,?,?,?)"

	stmt, err = con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	ctt := "Pembayaran untuk pesanan dengan ID: " + id_OD

	_, err = stmt.Exec(nm_str3, id_ntf, id_catering, id_OD, ctt)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Input Pengantar
func Set_Pegantar(id_detail_order string, id_pengantar string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "UPDATE detail_order SET id_pengantar=?, status_order=? WHERE id_detail_order=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_pengantar, "On Delivery", id_detail_order)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"

	return res, nil
}

//confirm makanan sukses di terima
func Confirm_Order(id string, id_detail_order string) (tools.Response, error) {
	var res tools.Response
	var st_D int
	var st_P int

	con := db.CreateCon()

	sqlstatement := ""

	if strings.HasPrefix(id, "USP") {

		sqlstatement = "UPDATE detail_order SET status_pengantar=? WHERE id_detail_order=?"

	} else if strings.HasPrefix(id, "US") {

		sqlstatement = "UPDATE detail_order SET status_pembeli=? WHERE id_detail_order=?"

	}

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(1, id_detail_order)

	if err != nil {
		return res, err
	}

	sqlstatement = "SELECT status_pengantar,status_pembeli FROM detail_order WHERE id_detail_order=?"

	err = con.QueryRow(sqlstatement, id_detail_order).Scan(&st_D, &st_P)

	if st_D == 1 && st_P == 1 {

		sqlstatement = "UPDATE detail_order SET status_order=?, delivery_sukses=? WHERE id_detail_order=?"

		stmt, err := con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec("Complate", 1, id_detail_order)

		if err != nil {
			return res, err
		}
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"

	return res, nil
}

//Show order / menu / tanggal
func Show_Order_Menu(id string, tanggal_menu string) (tools.Response, error) {
	var res tools.Response
	var arr []Order.Read_Id_Order
	var obj Order.Read_Id_Order

	var arr_order_menu_fix []Order.Read_Menu_Order
	var obj_order_menu_fix Order.Read_Menu_Order

	var arr_order_menu []Order.Menu_Order_Dipesan
	var obj_order_menu Order.Menu_Order_Dipesan

	con := db.CreateCon()

	sqlStatement := ""
	if strings.HasPrefix(id, "US") {
		sqlStatement = "SELECT id_order FROM order_catering WHERE order_catering.id_user=?"
	} else if strings.HasPrefix(id, "CT") {
		sqlStatement = "SELECT id_order FROM order_catering WHERE order_catering.id_catering=?"
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr
	}

	rows, err := con.Query(sqlStatement, id)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_Order)
		if err != nil {
			return res, err
		}
		arr = append(arr, obj)
	}

	if arr != nil {
		if tanggal_menu == "" {

			q1 := " WHERE detail_order.id_order IN ("
			q2 := "ORDER BY tanggal_menu ASC"
			q3 := " && tanggal_menu=?"
			sqlStatement = "SELECT DISTINCT(detail_order.tanggal_menu) FROM detail_order"

			for i := 0; i < len(arr); i++ {
				if i == len(arr)-1 {
					q1 = q1 + "'" + arr[i].Id_Order + "') && status_order != 'Complate' "
				} else {
					q1 = q1 + "'" + arr[i].Id_Order + "' , "
				}
			}

			sqlStatement = sqlStatement + q1 + q2

			rows, err = con.Query(sqlStatement)

			defer rows.Close()

			if err != nil {
				return res, err
			}

			for rows.Next() {
				err = rows.Scan(&obj_order_menu_fix.Tanggal_menu)

				if err != nil {
					return res, err
				}

				sqlStatement2 := "SELECT detail_order.id_order, id_detail_order, oc.id_catering,c.nama_catering,id_pengantar, nama_menu, harga_menu, status_order, radius, m.longtitude, m.langtitude  FROM detail_order JOIN order_catering oc on detail_order.id_order = oc.id_order JOIN catering c on c.id_catering = oc.id_catering JOIN maps m on c.id_catering = m.id_catering"

				sqlStatement2 = sqlStatement2 + q1 + q3

				rows2, err := con.Query(sqlStatement2, obj_order_menu_fix.Tanggal_menu)

				defer rows2.Close()

				if err != nil {
					return res, err
				}

				for rows2.Next() {
					err = rows2.Scan(&obj_order_menu.Id_order, &obj_order_menu.Id_detail_order, &obj_order_menu.Id_catering, &obj_order_menu.Nama_catering, &obj_order_menu.Id_pengantar, &obj_order_menu.Nama_menu, &obj_order_menu.Harga_menu, &obj_order_menu.Status_order, &obj_order_menu.Radius, &obj_order_menu.Longtitude, &obj_order_menu.Langtitude)

					if err != nil {
						return res, err
					}

					arr_order_menu = append(arr_order_menu, obj_order_menu)
				}
				obj_order_menu_fix.Menu_Order_Dipesan = arr_order_menu

				arr_order_menu_fix = append(arr_order_menu_fix, obj_order_menu_fix)
			}

			if arr_order_menu_fix == nil {
				res.Status = http.StatusNotFound
				res.Message = "Not Found"
				res.Data = arr_order_menu_fix
			} else {
				res.Status = http.StatusOK
				res.Message = "Sukses"
				res.Data = arr_order_menu_fix
			}
		} else {
			q1 := " WHERE detail_order.id_order IN ("
			q3 := " && tanggal_menu=?"

			for i := 0; i < len(arr); i++ {
				if i == len(arr)-1 {
					q1 = q1 + "'" + arr[i].Id_Order + "') && status_order != 'Complate' "
				} else {
					q1 = q1 + "'" + arr[i].Id_Order + "' , "
				}
			}

			obj_order_menu_fix.Tanggal_menu = tanggal_menu

			date, _ := time.Parse("02-01-2006", tanggal_menu)
			tanggal_menu_sql := date.Format("2006-01-02")

			sqlStatement2 := "SELECT detail_order.id_order, id_detail_order, oc.id_catering,c.nama_catering,id_pengantar, nama_menu, harga_menu, status_order, radius, m.longtitude, m.langtitude  FROM detail_order JOIN order_catering oc on detail_order.id_order = oc.id_order JOIN catering c on c.id_catering = oc.id_catering JOIN maps m on c.id_catering = m.id_catering"

			sqlStatement2 = sqlStatement2 + q1 + q3

			rows2, err := con.Query(sqlStatement2, tanggal_menu_sql)

			defer rows2.Close()

			if err != nil {
				return res, err
			}

			for rows2.Next() {
				err = rows2.Scan(&obj_order_menu.Id_order, &obj_order_menu.Id_detail_order, &obj_order_menu.Id_catering, &obj_order_menu.Nama_catering, &obj_order_menu.Id_pengantar, &obj_order_menu.Nama_menu, &obj_order_menu.Harga_menu, &obj_order_menu.Status_order, &obj_order_menu.Radius, &obj_order_menu.Longtitude, &obj_order_menu.Langtitude)

				if err != nil {
					return res, err
				}

				arr_order_menu = append(arr_order_menu, obj_order_menu)
			}
			obj_order_menu_fix.Menu_Order_Dipesan = arr_order_menu

			arr_order_menu_fix = append(arr_order_menu_fix, obj_order_menu_fix)

			if arr_order_menu_fix == nil {
				res.Status = http.StatusNotFound
				res.Message = "Not Found"
				res.Data = arr_order_menu_fix
			} else {
				res.Status = http.StatusOK
				res.Message = "Sukses"
				res.Data = arr_order_menu_fix
			}
		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr
	}

	return res, nil
}

//Order_Detail_User
func Order_Detail_User(id_detail_order string) (tools.Response, error) {
	var res tools.Response
	var arr []Order.Read_Detail_Order_User
	var obj Order.Read_Detail_Order_User

	con := db.CreateCon()

	sqlStatement := "SELECT nama_menu, jumlah, harga_menu,u.nama, u.telp_user, status_order, nama_catering,p.langtitude, p.longtitude FROM detail_order JOIN pengantar p on detail_order.id_pengantar = p.id_pengantar JOIN user u on p.id_user = u.id_user JOIN order_catering oc on detail_order.id_order = oc.id_order JOIN catering c on oc.id_catering = c.id_catering WHERE id_detail_order=?"

	err := con.QueryRow(sqlStatement, id_detail_order).Scan(&obj.Nama_menu, &obj.Jumlah, &obj.Harga, &obj.Nama_pengantar, &obj.Nomer_telp, &obj.Status, &obj.Nama_catering, &obj.Langtitude, &obj.Longtitude)

	if err != nil {
		arr = append(arr, obj)
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr

		return res, nil
	}

	arr = append(arr, obj)

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

//History Order
func History_Order(id_user string, tanggal_menu string) (tools.Response, error) {
	var res tools.Response
	var arr []Order.Read_Id_Order
	var obj Order.Read_Id_Order

	var arr_order_menu_fix []Order.Read_Menu_Order_His
	var obj_order_menu_fix Order.Read_Menu_Order_His

	var arr_order_his []Order.History_Order
	var obj_order_his Order.History_Order

	con := db.CreateCon()

	sqlStatement := "SELECT id_order FROM order_catering WHERE order_catering.id_user=?"

	rows, err := con.Query(sqlStatement, id_user)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_Order)
		if err != nil {
			return res, err
		}
		arr = append(arr, obj)
	}

	if arr != nil {
		if tanggal_menu == "" {
			q1 := " WHERE detail_order.id_order IN ("
			q2 := "ORDER BY tanggal_menu ASC"
			q3 := " && tanggal_menu=?"

			sqlStatement = "SELECT DISTINCT(detail_order.tanggal_menu) FROM detail_order"

			for i := 0; i < len(arr); i++ {
				if i == len(arr)-1 {
					q1 = q1 + "'" + arr[i].Id_Order + "') && status_order = 'Complate' "
				} else {
					q1 = q1 + "'" + arr[i].Id_Order + "' , "
				}
			}

			sqlStatement = sqlStatement + q1 + q2

			rows, err = con.Query(sqlStatement)

			defer rows.Close()

			if err != nil {
				fmt.Println(sqlStatement)
				return res, err
			}

			for rows.Next() {
				err = rows.Scan(&obj_order_menu_fix.Tanggal_menu)

				if err != nil {
					return res, err
				}

				sqlStatement2 := "SELECT detail_order.id_order, id_detail_order, oc.id_catering,c.nama_catering, nama_menu, harga_menu, status_order FROM detail_order JOIN order_catering oc on detail_order.id_order = oc.id_order JOIN catering c on c.id_catering = oc.id_catering"

				sqlStatement2 = sqlStatement2 + q1 + q3

				rows2, err := con.Query(sqlStatement2, obj_order_menu_fix.Tanggal_menu)

				defer rows2.Close()

				if err != nil {
					return res, err
				}

				for rows2.Next() {
					err = rows2.Scan(&obj_order_his.Id_order, &obj_order_his.Id_detail_order, &obj_order_his.Id_catering, &obj_order_his.Nama_catering, &obj_order_his.Nama_menu, &obj_order_his.Harga_menu, &obj_order_his.Status_order)

					if err != nil {
						return res, err
					}

					sqlStatement_rating := "SELECT rating FROM rating WHERE id_detail_order=?"

					err := con.QueryRow(sqlStatement_rating, obj_order_his.Id_detail_order).Scan(&obj_order_his.Rating)

					if err != nil {
						fmt.Println(err)
						obj_order_his.Rating = 0
					}

					arr_order_his = append(arr_order_his, obj_order_his)
				}
				obj_order_menu_fix.History_Order = arr_order_his

				arr_order_menu_fix = append(arr_order_menu_fix, obj_order_menu_fix)
			}

			if arr_order_menu_fix == nil {
				res.Status = http.StatusNotFound
				res.Message = "Not Found"
				res.Data = arr_order_menu_fix
			} else {
				res.Status = http.StatusOK
				res.Message = "Sukses"
				res.Data = arr_order_menu_fix
			}
		} else {

			q1 := " WHERE detail_order.id_order IN ("
			q3 := " && tanggal_menu=?"

			for i := 0; i < len(arr); i++ {
				if i == len(arr)-1 {
					q1 = q1 + "'" + arr[i].Id_Order + "') && status_order = 'Complate' "
				} else {
					q1 = q1 + "'" + arr[i].Id_Order + "' , "
				}
			}

			obj_order_menu_fix.Tanggal_menu = tanggal_menu

			date, _ := time.Parse("02-01-2006", tanggal_menu)
			tanggal_menu_sql := date.Format("2006-01-02")

			sqlStatement2 := "SELECT detail_order.id_order, id_detail_order, oc.id_catering,c.nama_catering, nama_menu, harga_menu, status_order FROM detail_order JOIN order_catering oc on detail_order.id_order = oc.id_order JOIN catering c on c.id_catering = oc.id_catering"

			sqlStatement2 = sqlStatement2 + q1 + q3

			rows2, err := con.Query(sqlStatement2, tanggal_menu_sql)

			defer rows2.Close()

			if err != nil {
				return res, err
			}

			for rows2.Next() {
				err = rows2.Scan(&obj_order_his.Id_order, &obj_order_his.Id_detail_order, &obj_order_his.Id_catering, &obj_order_his.Nama_catering, &obj_order_his.Nama_menu, &obj_order_his.Harga_menu, &obj_order_his.Status_order)

				if err != nil {
					return res, err
				}

				sqlStatement_rating := "SELECT rating FROM rating WHERE id_detail_order=?"

				err := con.QueryRow(sqlStatement_rating, obj_order_his.Id_detail_order).Scan(&obj_order_his.Rating)

				if err != nil {
					fmt.Println(err)
					obj_order_his.Rating = 0
				}

				arr_order_his = append(arr_order_his, obj_order_his)
			}
			obj_order_menu_fix.History_Order = arr_order_his

			arr_order_menu_fix = append(arr_order_menu_fix, obj_order_menu_fix)

			if arr_order_menu_fix == nil {
				res.Status = http.StatusNotFound
				res.Message = "Not Found"
				res.Data = arr_order_menu_fix
			} else {
				res.Status = http.StatusOK
				res.Message = "Sukses"
				res.Data = arr_order_menu_fix
			}

		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr
	}

	return res, nil
}

//Read Order Pengantar
func Read_Order_Pengantar(id_pengantar string) (tools.Response, error) {
	var res tools.Response
	var arr_order_menu_fix []Order.Read_Menu_Order
	var obj_order_menu_fix Order.Read_Menu_Order

	var arr_order_menu []Order.Menu_Order_Dipesan
	var obj_order_menu Order.Menu_Order_Dipesan

	con := db.CreateCon()

	sqlStatement := "SELECT DISTINCT(detail_order.tanggal_menu) FROM detail_order WHERE id_pengantar=?"

	rows, err := con.Query(sqlStatement, id_pengantar)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj_order_menu_fix.Tanggal_menu)

		if err != nil {
			return res, err
		}

		sqlStatement2 := "SELECT detail_order.id_order, id_detail_order, oc.id_catering,c.nama_catering,id_pengantar, nama_menu, harga_menu, status_order, radius, m.longtitude, m.langtitude,oc.alamat  FROM detail_order JOIN order_catering oc on detail_order.id_order = oc.id_order JOIN catering c on c.id_catering = oc.id_catering JOIN maps m on c.id_catering = m.id_catering WHERE detail_order.id_pengantar=? && tanggal_menu=? && status_order=?"

		rows2, err := con.Query(sqlStatement2, id_pengantar, obj_order_menu_fix.Tanggal_menu, "On Delivery")

		defer rows2.Close()

		if err != nil {
			return res, err
		}

		for rows2.Next() {
			err = rows2.Scan(&obj_order_menu.Id_order, &obj_order_menu.Id_detail_order, &obj_order_menu.Id_catering, &obj_order_menu.Nama_catering, &obj_order_menu.Id_pengantar, &obj_order_menu.Nama_menu, &obj_order_menu.Harga_menu, &obj_order_menu.Status_order, &obj_order_menu.Radius, &obj_order_menu.Longtitude, &obj_order_menu.Langtitude, &obj_order_menu.Alamat)

			if err != nil {
				return res, err
			}

			arr_order_menu = append(arr_order_menu, obj_order_menu)
		}
		obj_order_menu_fix.Menu_Order_Dipesan = arr_order_menu

		tanggal, _ := time.Parse("2006-01-02", obj_order_menu_fix.Tanggal_menu)
		obj_order_menu_fix.Tanggal_menu = tanggal.Format("02-01-2006")

		arr_order_menu_fix = append(arr_order_menu_fix, obj_order_menu_fix)
	}

	if arr_order_menu_fix == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_order_menu_fix
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_order_menu_fix
	}

	return res, nil
}

//Read Location User
func Read_Location_User(id_detail_order string) (tools.Response, error) {
	var res tools.Response
	var arr []_struct.Read_Long_Lang
	var obj _struct.Read_Long_Lang

	con := db.CreateCon()

	sqlStatement := "SELECT oc.longtitude,oc.langtitude FROM detail_order JOIN order_catering oc on oc.id_order = detail_order.id_order WHERE id_detail_order=?"

	err := con.QueryRow(sqlStatement, id_detail_order).Scan(&obj.Longitude, &obj.Langitude)

	if err != nil {
		arr = append(arr, obj)
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr

		return res, nil
	}

	arr = append(arr, obj)

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

//Filter Order Menu
func Filter_Order_Menu(id string, tanggal_menu string) (tools.Response, error) {
	var res tools.Response
	var arr []Order.Read_Id_Order
	var obj Order.Read_Id_Order

	var arr_order_menu_fix []Order.Read_Menu_Order
	var obj_order_menu_fix Order.Read_Menu_Order

	var arr_order_menu []Order.Menu_Order_Dipesan
	var obj_order_menu Order.Menu_Order_Dipesan

	con := db.CreateCon()

	sqlStatement := ""
	if strings.HasPrefix(id, "US") {
		sqlStatement = "SELECT id_order FROM order_catering WHERE order_catering.id_user=?"
	} else if strings.HasPrefix(id, "CT") {
		sqlStatement = "SELECT id_order FROM order_catering WHERE order_catering.id_catering=?"
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr
	}

	rows, err := con.Query(sqlStatement, id)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_Order)
		if err != nil {
			return res, err
		}
		arr = append(arr, obj)
	}

	if arr != nil {

		q1 := " WHERE detail_order.id_order IN ("
		q3 := " && tanggal_menu=?"
		sqlStatement = "SELECT DISTINCT(detail_order.tanggal_menu) FROM detail_order"

		for i := 0; i < len(arr); i++ {
			if i == len(arr)-1 {
				q1 = q1 + "'" + arr[i].Id_Order + "') && status_order != 'Complate' "
			} else {
				q1 = q1 + "'" + arr[i].Id_Order + "' , "
			}
		}

		obj_order_menu_fix.Tanggal_menu = tanggal_menu

		date, _ := time.Parse("02-01-2006", tanggal_menu)
		tanggal_menu_sql := date.Format("2006-01-02")

		sqlStatement2 := "SELECT detail_order.id_order, id_detail_order, oc.id_catering,c.nama_catering,id_pengantar, nama_menu, harga_menu, status_order, radius, m.longtitude, m.langtitude  FROM detail_order JOIN order_catering oc on detail_order.id_order = oc.id_order JOIN catering c on c.id_catering = oc.id_catering JOIN maps m on c.id_catering = m.id_catering"

		sqlStatement2 = sqlStatement2 + q1 + q3

		rows2, err := con.Query(sqlStatement2, tanggal_menu_sql)

		defer rows2.Close()

		if err != nil {
			return res, err
		}

		for rows2.Next() {
			err = rows2.Scan(&obj_order_menu.Id_order, &obj_order_menu.Id_detail_order, &obj_order_menu.Id_catering, &obj_order_menu.Nama_catering, &obj_order_menu.Id_pengantar, &obj_order_menu.Nama_menu, &obj_order_menu.Harga_menu, &obj_order_menu.Status_order, &obj_order_menu.Radius, &obj_order_menu.Longtitude, &obj_order_menu.Langtitude)

			if err != nil {
				return res, err
			}

			arr_order_menu = append(arr_order_menu, obj_order_menu)
		}
		obj_order_menu_fix.Menu_Order_Dipesan = arr_order_menu

		arr_order_menu_fix = append(arr_order_menu_fix, obj_order_menu_fix)

		if arr_order_menu_fix == nil {
			res.Status = http.StatusNotFound
			res.Message = "Not Found"
			res.Data = arr_order_menu_fix
		} else {
			res.Status = http.StatusOK
			res.Message = "Sukses"
			res.Data = arr_order_menu_fix
		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr
	}

	return res, nil
}

//Filter History Order
func Filter_History_Order(id_user string, tanggal_menu string) (tools.Response, error) {
	var res tools.Response
	var arr []Order.Read_Id_Order
	var obj Order.Read_Id_Order

	var arr_order_menu_fix []Order.Read_Menu_Order_His
	var obj_order_menu_fix Order.Read_Menu_Order_His

	var arr_order_his []Order.History_Order
	var obj_order_his Order.History_Order

	con := db.CreateCon()

	sqlStatement := "SELECT id_order FROM order_catering WHERE order_catering.id_user=?"

	rows, err := con.Query(sqlStatement, id_user)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_Order)
		if err != nil {
			return res, err
		}
		arr = append(arr, obj)
	}

	if arr != nil {

		q1 := " WHERE detail_order.id_order IN ("
		q3 := " && tanggal_menu=?"

		for i := 0; i < len(arr); i++ {
			if i == len(arr)-1 {
				q1 = q1 + "'" + arr[i].Id_Order + "') && status_order = 'Complate' "
			} else {
				q1 = q1 + "'" + arr[i].Id_Order + "' , "
			}
		}

		obj_order_menu_fix.Tanggal_menu = tanggal_menu

		date, _ := time.Parse("02-01-2006", tanggal_menu)
		tanggal_menu_sql := date.Format("2006-01-02")

		sqlStatement2 := "SELECT detail_order.id_order, id_detail_order, oc.id_catering,c.nama_catering, nama_menu, harga_menu, status_order FROM detail_order JOIN order_catering oc on detail_order.id_order = oc.id_order JOIN catering c on c.id_catering = oc.id_catering"

		sqlStatement2 = sqlStatement2 + q1 + q3

		rows2, err := con.Query(sqlStatement2, tanggal_menu_sql)

		defer rows2.Close()

		if err != nil {
			return res, err
		}

		for rows2.Next() {
			err = rows2.Scan(&obj_order_his.Id_order, &obj_order_his.Id_detail_order, &obj_order_his.Id_catering, &obj_order_his.Nama_catering, &obj_order_his.Nama_menu, &obj_order_his.Harga_menu, &obj_order_his.Status_order)

			if err != nil {
				return res, err
			}

			sqlStatement_rating := "SELECT rating FROM rating WHERE id_detail_order=?"

			err := con.QueryRow(sqlStatement_rating, obj_order_his.Id_detail_order).Scan(&obj_order_his.Rating)

			if err != nil {
				fmt.Println(err)
				obj_order_his.Rating = 0
			}

			arr_order_his = append(arr_order_his, obj_order_his)
		}
		obj_order_menu_fix.History_Order = arr_order_his

		arr_order_menu_fix = append(arr_order_menu_fix, obj_order_menu_fix)

		if arr_order_menu_fix == nil {
			res.Status = http.StatusNotFound
			res.Message = "Not Found"
			res.Data = arr_order_menu_fix
		} else {
			res.Status = http.StatusOK
			res.Message = "Sukses"
			res.Data = arr_order_menu_fix
		}

	} else {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr
	}

	return res, nil
}
