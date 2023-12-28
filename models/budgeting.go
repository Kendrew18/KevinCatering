package models

import (
	"KevinCatering/db"
	"KevinCatering/struct/Budgeting"
	"KevinCatering/tools"
	"net/http"
	"strconv"
	"time"
)

//Input_Budgeting
func Input_Budgeting(id_catering string, total_porsi int, tanggal_budgeting string, id_master_menu string) (tools.Response, error) {
	var res tools.Response
	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM budgeting ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_BD := "BD-" + strconv.Itoa(nm_str)

	date, _ := time.Parse("02-01-2006", tanggal_budgeting)
	date_sql := date.Format("2006-01-02")

	sqlStatement := "INSERT INTO budgeting (co, id_budgeting, id_master_menu, id_catering, total_porsi, tanggal_budgeting) values(?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_BD, id_master_menu, id_catering, total_porsi, date_sql)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read_Budgeting
func Read_List_Budgeting(id_catering string) (tools.Response, error) {
	var res tools.Response
	var arr []Budgeting.Read_Budgeting_Awal
	var obj Budgeting.Read_Budgeting_Awal

	con := db.CreateCon()

	sqlStatement := "SELECT id_budgeting, mm.nama_menu FROM budgeting join master_menu mm on mm.id_master_menu = budgeting.id_master_menu WHERE budgeting.id_catering=? && budgeting.status = 0"

	rows, err := con.Query(sqlStatement, id_catering)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_budgeting, &obj.Nama_menu)
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

//Read_Detail_Budgeting
func Read_Budgeting(id_budgeting string) (tools.Response, error) {
	var res tools.Response
	var obj Budgeting.Read_Budgeting
	var arr []Budgeting.Read_Budgeting

	var bahan Budgeting.Read_Bahan
	var arr_bahan []Budgeting.Read_Bahan

	con := db.CreateCon()

	sqlStatement := "SELECT id_budgeting,mm.id_master_menu, nama_menu, total_porsi, tanggal_budgeting FROM budgeting JOIN master_menu mm on mm.id_master_menu = budgeting.id_master_menu WHERE id_budgeting=? "

	err := con.QueryRow(sqlStatement, id_budgeting).Scan(&obj.Id_budgeting, &obj.Id_Master_Menu, &obj.Nama_menu, &obj.Total_porsi, &obj.Tanggal_budgeting)

	if err != nil {
		return res, err
	}

	sqlStatement = "SELECT id_bahan_menu, nama_bahan, jumlah_bahan, satuan_bahan, harga_bahan FROM bahan_menu WHERE id_master_menu=?"

	rows, err := con.Query(sqlStatement, obj.Id_Master_Menu)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&bahan.Id_bahan, &bahan.Nama_bahan, &bahan.Jumlah_bahan, &bahan.Satuan_bahan, &bahan.Harga_bahan)

		bahan.Jumlah_bahan = bahan.Jumlah_bahan * float64(obj.Total_porsi)

		bahan.Harga_bahan = bahan.Harga_bahan * obj.Total_porsi

		if err != nil {
			return res, err
		}
		arr_bahan = append(arr_bahan, bahan)
	}

	obj.Bahan = arr_bahan

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

//Update_Status
func Update_Status_Budgeting(id_budgeting string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "UPDATE budgeting SET status=? WHERE id_budgeting=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(1, id_budgeting)

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
