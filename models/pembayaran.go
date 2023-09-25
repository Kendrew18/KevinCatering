package models

import (
	"KevinCatering/db"
	"KevinCatering/struct/Order"
	str "KevinCatering/struct/Pembayaran"
	"KevinCatering/tools"
	"fmt"
	"net/http"
	"strings"
	"time"
)

//Show_History_Order_Recipe [berubah]
func Read_Order_Recipe(id string, tanggal_recipe string) (tools.Response, error) {
	var res tools.Response
	var arr []Order.Read_Order
	var obj Order.Read_Order

	con := db.CreateCon()

	sqlStatement := ""

	rows, err := con.Query(sqlStatement)

	if tanggal_recipe != "" {

		date_recipe, _ := time.Parse("02-01-2006", tanggal_recipe)
		date_recipe_sql := date_recipe.Format("2006-01-02")

		if strings.HasPrefix(id, "US") {
			sqlStatement = "SELECT id_order, nama_catering,DATE_FORMAT(tanggal_order, '%d-%m-%Y') FROM order_catering JOIN catering c on order_catering.id_catering = c.id_catering WHERE order_catering.id_user=? && tanggal_order=? ORDER BY tanggal_order DESC "
		} else if strings.HasPrefix(id, "CT") {
			sqlStatement = "SELECT id_order, nama_catering,DATE_FORMAT(tanggal_order, '%d-%m-%Y') FROM order_catering JOIN catering c on order_catering.id_catering = c.id_catering WHERE order_catering.id_catering=? && tanggal_order=? ORDER BY tanggal_order DESC"
		} else {
			res.Status = http.StatusNotFound
			res.Message = "Not Found"
			res.Data = arr
		}

		rows, err = con.Query(sqlStatement, id, date_recipe_sql)
	} else {

		if strings.HasPrefix(id, "US") {
			sqlStatement = "SELECT id_order, nama_catering,DATE_FORMAT(tanggal_order, '%d-%m-%Y') FROM order_catering JOIN catering c on order_catering.id_catering = c.id_catering WHERE order_catering.id_user=? ORDER BY tanggal_order DESC"
		} else if strings.HasPrefix(id, "CT") {
			sqlStatement = "SELECT id_order, nama_catering,DATE_FORMAT(tanggal_order, '%d-%m-%Y') FROM order_catering JOIN catering c on order_catering.id_catering = c.id_catering WHERE order_catering.id_catering=? ORDER BY tanggal_order DESC"
		} else {
			res.Status = http.StatusNotFound
			res.Message = "Not Found"
			res.Data = arr
		}

		rows, err = con.Query(sqlStatement, id)
	}

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

	sqlStatement := "SELECT order_catering.id_order,order_catering.id_user,nama,u.telp_user,order_catering.id_catering, nama_catering,c.telp_catering,c.alamat_catering,tanggal_order,total,longtitude,langtitude,bukti_pembayaran FROM order_catering JOIN catering c on order_catering.id_catering = c.id_catering join user u on order_catering.id_user = u.id_user JOIN pembayaran p on order_catering.id_order = p.id_order WHERE order_catering.id_order=?"

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
		err = rows.Scan(&menu.Id_menu, &menu.Nama_menu, &menu.Tanggal_menu, &menu.Jumlah_menu, &menu.Harga_menu, &menu.Jam_pengiriman_awal, &menu.Jam_pengiriman_akhir, &menu.Status_order)
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

//Confirm_Pembayaran [aplikasi catering]
func Confirm_Pembayaran(id_order string) (tools.Response, error) {
	var res tools.Response

	var Read_Pembayaran_Fix str.Read_Pembayaran_Fix

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

	sqlstatement = "SELECT id_pembayaran, id_order, status_pembayaran FROM pembayaran WHERE id_order=?"

	err = con.QueryRow(sqlstatement, id_order).Scan(&Read_Pembayaran_Fix.Id_pembayaran, &Read_Pembayaran_Fix.Id_order, &Read_Pembayaran_Fix.Status_pembayaran)

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = Read_Pembayaran_Fix

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

//Read_Notif [aplikasi catering]
func Read_Notif(id_order string) (tools.Response, error) {
	var res tools.Response
	var notif str.Read_Notif
	var menu str.Menu_Order_Notif
	var arr_menu []str.Menu_Order_Notif

	con := db.CreateCon()

	sqlStatement := "SELECT detail_order.id_menu,detail_order.nama_menu,detail_order.tanggal_menu,detail_order.jumlah,detail_order.harga_menu,jam_pengiriman_awal,jam_pengiriman_akhir,detail_order.status_order FROM detail_order join menu m on detail_order.id_menu = m.id_menu WHERE detail_order.id_order=?"

	rows, err := con.Query(sqlStatement, id_order)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&menu.Id_menu, &menu.Nama_menu, &menu.Tanggal_menu, &menu.Jumlah_menu, &menu.Harga_menu, &menu.Jam_pengiriman_awal, &menu.Jam_pengiriman_akhir, &menu.Status_order)
		if err != nil {
			return res, err
		}
		arr_menu = append(arr_menu, menu)
	}

	sqlStatement = "SELECT bukti_pembayaran FROM pembayaran WHERE id_order=?"

	err = con.QueryRow(sqlStatement, id_order).Scan(&notif.Path_Foto)

	notif.Menu_Order_Notif = arr_menu

	if arr_menu == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = notif
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = notif
	}

	return res, nil
}
