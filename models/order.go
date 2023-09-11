package models

import (
	"KevinCatering/db"
	str "KevinCatering/struct"
	"KevinCatering/tools"
	"net/http"
	"strconv"
	"time"
)

//Input_Order
func Input_Order(id_catering string, id_user string, id_menu string, nama_menu string, jumlah string, harga_menu string, tanggal_menu string, tanggal_order string, langtitude float64, longtitude float64) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM order_catering ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_OD := "OR-" + strconv.Itoa(nm_str)

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

	sqlStatement := "INSERT INTO order_catering (co, id_order,id_catering,id_user,total,tanggal_order,longtitude,langtitude) values(?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_OD, id_catering, id_user, total, date_sql2, longtitude, langtitude)

	for i := 0; i < len(id_M); i++ {
		nm_str2 := 0

		Sqlstatement := "SELECT co FROM Detail_Order ORDER BY co DESC Limit 1"

		_ = con.QueryRow(Sqlstatement).Scan(&nm_str2)

		nm_str2 = nm_str2 + 1

		id_DO := "DO-" + strconv.Itoa(nm_str2)

		date, _ := time.Parse("02-01-2006", tgl_mn[i])
		date_sql := date.Format("2006-01-02")

		sqlStatement := "INSERT INTO Detail_Order (co,id_detail_order, id_order, id_menu, nama_menu, tanggal_menu,jumlah, harga_menu, status_order) values(?,?,?,?,?,?,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(nm_str2, id_DO, id_OD, id_M[i], nama_mn[i], date_sql, jmlh_mn[i], harga_mn[i], 0)

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

	_, err = stmt.Exec(nm_str2, id_OD2, id_OD, 0, "")

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read_Order
func Read_Order(id_user string) (tools.Response, error) {
	var res tools.Response
	var arr []str.Read_Order
	var obj str.Read_Order

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
func Read_Detail_Order(id_order string) (tools.Response, error) {
	var res tools.Response
	var arr []str.Read_Detail_Order
	var obj str.Read_Detail_Order

	var menu str.Menu_Order
	var arr_menu []str.Menu_Order

	con := db.CreateCon()

	sqlStatement := "SELECT id_order,order_catering.id_user,nama,u.telp_user,order_catering.id_catering, nama_catering,c.telp_catering,c.alamat_catering,tanggal_order,total,longtitude,langtitude FROM order_catering JOIN catering c on order_catering.id_catering = c.id_catering join user u on order_catering.id_user = u.id_user WHERE id_order=?"

	err := con.QueryRow(sqlStatement, id_order).Scan(&obj.Id_order, &obj.Id_user, &obj.Nama_user, &obj.No_telp_user,
		&obj.Id_catering, &obj.Nama_catering, &obj.No_telp_catering, &obj.Alamat_catering,
		&obj.Tanggal_order, &obj.Total, &obj.Longtitude, &obj.Langtitude)

	if err != nil {
		return res, err
	}

	sqlStatement = "SELECT Detail_Order.id_menu,Detail_Order.nama_menu,Detail_Order.tanggal_menu,Detail_Order.jumlah,Detail_Order.harga_menu,jam_pengiriman_awal,jam_pengiriman_akhir,Detail_Order.status_order FROM Detail_Order join menu m on Detail_Order.id_menu = m.id_menu WHERE id_order=?"

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

//Chage_Status_Order
func Change_Status_Menu_Order() {

}
