package models

import (
	"KevinCatering/db"
	str "KevinCatering/struct"
	"KevinCatering/tools"
	"net/http"
	"strconv"
	"time"
)

//Generate_id_user
func Generate_Id_Order() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_order FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_order=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

func Generate_Id_Pembayaran() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_order FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_order=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

//Input_Order
func Input_Order(id_catering string, id_user string, id_menu string, nama_menu string, harga_menu string,
	tanggal_order string, tanggal_menu string, status_order string, langtitude float64, longtitude float64) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Order()

	nm_str := strconv.Itoa(nm)

	id_OD := "OR-" + nm_str

	date, _ := time.Parse("02-01-2006", tanggal_menu)
	date_sql := date.Format("2006-01-02")

	date2, _ := time.Parse("02-01-2006", tanggal_order)
	date_sql2 := date2.Format("2006-01-02")

	sqlStatement := "INSERT INTO order_catering (id_order,id_catering,id_user, id_menu, nama_menu, harga_menu, total, tanggal_menu,tanggal_order, status_order,longtitude,langtitude) values(?,?,?,?,?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	hg_mn := tools.String_Separator_To_Int64(harga_menu)

	var total int64

	total = 0

	for i := 0; i < len(hg_mn); i++ {
		total = total + hg_mn[i]
	}

	_, err = stmt.Exec(id_OD, id_catering, id_user, id_menu, nama_menu, harga_menu,
		total, date_sql, date_sql2, status_order, longtitude, langtitude)

	nm2 := Generate_Id_Pembayaran()

	nm_str2 := strconv.Itoa(nm2)

	id_OD2 := "PBR-" + nm_str2

	sqlStatement = "INSERT INTO pembayaran (id_pembayaran,id_order,status_pembayaran) values(?,?,?)"

	stmt, err = con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_OD, id_OD2, 0)

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

	sqlStatement := "SELECT id_order, nama_catering,tanggal_order FROM order_catering JOIN catering c on order_catering.id_catering = c.id_catering WHERE id_user=?"

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
	var obj_str str.Menu_Order_String

	var menu str.Menu_Order
	var arr_menu []str.Menu_Order

	con := db.CreateCon()

	sqlStatement := "SELECT id_order,order_catering.id_user,nama,u.telp_user,order_catering.id_catering, nama_catering,c.telp_catering,c.alamat_catering,tanggal_order,total,longtitude,langtitude, id_menu,nama_menu,harga_menu,tanggal_menu,status_order FROM order_catering JOIN catering c on order_catering.id_catering = c.id_catering join user u on order_catering.id_user = u.id_user WHERE id_order=?"

	err := con.QueryRow(sqlStatement, id_order).Scan(&obj.Id_order, &obj.Id_user, &obj.Nama_user, &obj.No_telp_user,
		&obj.Id_catering, &obj.Nama_catering, &obj.No_telp_catering, &obj.Alamat_catering,
		&obj.Tanggal_order, &obj.Total, &obj.Longtitude, &obj.Langtitude, &obj_str.Id_menu, &obj_str.Nama_menu, &obj_str.Harga_menu,
		&obj_str.Tanggal_menu, &obj_str.Status_order)

	if err != nil {
		return res, err
	}

	id_mn_all := tools.String_Separator_To_String(obj_str.Id_menu)
	nm_mn := tools.String_Separator_To_String(obj_str.Nama_menu)
	hg_mn := tools.String_Separator_To_Int64(obj_str.Harga_menu)
	tgl_mn := tools.String_Separator_To_String(obj_str.Tanggal_menu)
	st := tools.String_Separator_To_Int(obj_str.Status_order)

	for i := 0; i < len(id_mn_all); i++ {
		menu.Id_menu = id_mn_all[i]
		menu.Nama_menu = nm_mn[i]
		menu.Harga_menu = hg_mn[i]
		menu.Tanggal_menu = tgl_mn[i]
		menu.Status_order = st[i]
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
