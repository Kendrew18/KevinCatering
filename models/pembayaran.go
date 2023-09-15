package models

import (
	"KevinCatering/db"
	str "KevinCatering/struct"
	"KevinCatering/struct/Order"
	"KevinCatering/tools"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//upload-foto-pembayaran
func Upload_Foto_Pembayaran(id_order string, writer http.ResponseWriter, request *http.Request) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

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
		path = "uploads/" + id_order + ".jpg"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.jpg")
	}
	if strings.Contains(handler.Filename, "jpeg") {
		path = "uploads/" + id_order + ".jpeg"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.jpeg")
	}
	if strings.Contains(handler.Filename, "png") {
		path = "uploads/" + id_order + ".png"
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

	sqlstatement := "UPDATE pembayaran SET bukti_pembayaran=? WHERE id_order=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(path, id_order)

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

//Show_History_Order_Recipe
func Read_Order_Recipe(id_user string) (tools.Response, error) {
	var res tools.Response
	var arr []Order.Read_Order
	var obj Order.Read_Order

	con := db.CreateCon()

	sqlStatement := "SELECT id_order, nama_catering,tanggal_order FROM order_catering JOIN catering c on order_catering.id_catering = c.id_catering WHERE order_catering.id_user=?"

	rows, err := con.Query(sqlStatement, id_user)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_order, &obj.Nama_catering, &obj.Tanggal_order)
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

//Read_Detail_Order
func Read_Detail_Order_Recipe(id_order string) (tools.Response, error) {
	var res tools.Response
	var arr []Order.Read_Detail_Order
	var obj Order.Read_Detail_Order

	var menu Order.Menu_Order
	var arr_menu []Order.Menu_Order

	con := db.CreateCon()

	sqlStatement := "SELECT id_order,order_catering.id_user,nama,u.telp_user,order_catering.id_catering, nama_catering,c.telp_catering,c.alamat_catering,tanggal_order,total,longtitude,langtitude,bukti_pembayaran FROM order_catering JOIN catering c on order_catering.id_catering = c.id_catering join user u on order_catering.id_user = u.id_user JOIN pembayaran p on order_catering.id_order = p.id_order WHERE order_catering.id_order=?"

	err := con.QueryRow(sqlStatement, id_order).Scan(&obj.Id_order, &obj.Id_user, &obj.Nama_user, &obj.No_telp_user, &obj.Id_catering, &obj.Nama_catering, &obj.No_telp_catering, &obj.Alamat_catering, &obj.Tanggal_order, &obj.Total, &obj.Longtitude, &obj.Langtitude, &obj.Foto_Bukti_pembayaran)

	if err != nil {
		return res, err
	}

	sqlStatement = "SELECT detail_order.id_menu,detail_order.nama_menu,detail_order.tanggal_menu,detail_order.jumlah,detail_order.harga_menu,jam_pengiriman_awal,jam_pengiriman_akhir,detail_order.status_order FROM detail_order join menu m on detail_order.id_menu = m.id_menu WHERE id_order=?"

	rows, err := con.Query(sqlStatement, id_order)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&menu.Id_menu, &menu.Nama_menu, &menu.Tanggal_menu, &menu.Jumlah_menu,
			&menu.Harga_menu, &menu.Jam_pengiriman_awal, &menu.Jam_pengiriman_akhir, &menu.Status_order)
		if err != nil {
			return res, err
		}
		arr_menu = append(arr_menu, menu)
	}

	obj.Menu_order = arr_menu

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

//Read_Notif
func Read_Notif(id_order string) (tools.Response, error) {
	var res tools.Response
	var menu Order.Menu_Order
	var arr_menu []Order.Menu_Order

	con := db.CreateCon()

	sqlStatement := "SELECT detail_order.id_menu,detail_order.nama_menu,detail_order.tanggal_menu,detail_order.jumlah,detail_order.harga_menu,jam_pengiriman_awal,jam_pengiriman_akhir,detail_order.status_order FROM detail_order join menu m on detail_order.id_menu = m.id_menu WHERE id_order=?"

	rows, err := con.Query(sqlStatement, id_order)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&menu.Id_menu, &menu.Nama_menu, &menu.Tanggal_menu, &menu.Jumlah_menu,
			&menu.Harga_menu, &menu.Jam_pengiriman_awal, &menu.Jam_pengiriman_akhir, &menu.Status_order)
		if err != nil {
			return res, err
		}
		arr_menu = append(arr_menu, menu)
	}

	if arr_menu == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_menu
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_menu
	}

	return res, nil
}

//Confirm_Pembayaran
func Confirm_Pembayaran(id_order string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "UPDATE pembayaran SET status_pembayaran=? WHERE id_order=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(1, id_order)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"

	return res, nil
}

//Read_Pembayaran
func Read_Pembayaran(id_order string) (tools.Response, error) {
	var res tools.Response
	var arr []str.Read_pembayaran
	var obj str.Read_pembayaran

	con := db.CreateCon()

	sqlStatement := "SELECT id_pembayaran,id_order,bukti_pembayaran,status_pembayaran FROM pembayaran WHERE id_order=?"

	fmt.Println(id_order)

	rows, err := con.Query(sqlStatement, id_order)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_pembayaran, &obj.Id_order, &obj.Bukti_pembayaran, &obj.Status_pembayaran)
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
