package Notif

import (
	"KevinCatering/db"
	"KevinCatering/struct/Notif"
	"KevinCatering/tools"
	"net/http"
)

func Show_All_Notif(id_catering string) (tools.Response, error) {
	var res tools.Response
	var Read_notif Notif.Read_Notif
	var Arr_Read_notif []Notif.Read_Notif

	con := db.CreateCon()

	sqlStatement := "SELECT id_notif, id_catering, id_order, catatan FROM notif WHERE id_catering=?"

	rows, err := con.Query(sqlStatement, id_catering)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&Read_notif.Id_notif, &Read_notif.Id_catering, &Read_notif.Id_order, &Read_notif.Catatan)
		if err != nil {
			return res, err
		}
		Arr_Read_notif = append(Arr_Read_notif, Read_notif)
	}

	if Arr_Read_notif == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = Arr_Read_notif
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = Arr_Read_notif
	}

	return res, nil
}

//Read_Notif [aplikasi catering]
func Read_Detail_Notif(id_order string) (tools.Response, error) {
	var res tools.Response
	var notif Notif.Read_Detail_Notif
	var menu Notif.Menu_Order_Notif
	var arr_menu []Notif.Menu_Order_Notif

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
