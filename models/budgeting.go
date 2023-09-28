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
func Input_Budgeting(id_catering string, nama_menu string, total_porsi int, tanggal_budgeting string, nama_bahan string, jumlah_bahan string, satuan_bahan string, harga_bahan string) (tools.Response, error) {
	var res tools.Response
	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM budgeting ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_BD := "BD-" + strconv.Itoa(nm_str)

	date, _ := time.Parse("02-01-2006", tanggal_budgeting)
	date_sql := date.Format("2006-01-02")

	sqlStatement := "INSERT INTO budgeting (co, id_budgeting, id_catering, nama_menu, total_porsi, tanggal_budgeting) values(?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_BD, id_catering, nama_menu, total_porsi, date_sql)

	nama_bahan_SS := tools.String_Separator_To_String(nama_bahan)
	jumlah_bahan_SS := tools.String_Separator_To_float64(jumlah_bahan)
	satuan_bahan_SS := tools.String_Separator_To_String(satuan_bahan)
	harga_bahan_SS := tools.String_Separator_To_Int64(harga_bahan)

	for i := 0; i < len(nama_bahan_SS); i++ {
		nm_str_BHN := 0

		Sqlstatement_BHN := "SELECT co FROM bahan_menu ORDER BY co DESC Limit 1"

		_ = con.QueryRow(Sqlstatement_BHN).Scan(&nm_str_BHN)

		nm_str_BHN = nm_str_BHN + 1

		id_BHN := "BHN-" + strconv.Itoa(nm_str_BHN)

		Sqlstatement_BHN = "INSERT INTO bahan_menu (co, id_bahan_menu, id_budgeting, nama_bahan, jumlah_bahan, satuan_bahan, harga_bahan) values(?,?,?,?,?,?,?)"

		stmt_BHN, err := con.Prepare(Sqlstatement_BHN)

		if err != nil {
			return res, err
		}

		_, err = stmt_BHN.Exec(nm_str_BHN, id_BHN, id_BD, nama_bahan_SS[i], jumlah_bahan_SS[i], satuan_bahan_SS[i], harga_bahan_SS[i])

		if err != nil {
			return res, err
		}
	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read_Budgeting
func Read_Budgeting_Awal(id_catering string) (tools.Response, error) {
	var res tools.Response
	var arr []Budgeting.Read_Budgeting_Awal
	var obj Budgeting.Read_Budgeting_Awal

	con := db.CreateCon()

	sqlStatement := "SELECT id_budgeting,nama_menu FROM budgeting WHERE id_catering=?"

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

	sqlStatement := "SELECT id_budgeting, nama_menu, total_porsi, tanggal_budgeting FROM budgeting WHERE id_budgeting=? "

	err := con.QueryRow(sqlStatement, id_budgeting).Scan(&obj.Id_budgeting, &obj.Nama_menu, &obj.Total_porsi, &obj.Tanggal_budgeting)

	if err != nil {
		return res, err
	}

	sqlStatement = "SELECT id_bahan_menu, nama_bahan, jumlah_bahan, satuan_bahan, harga_bahan FROM bahan_menu WHERE id_budgeting=?"

	rows, err := con.Query(sqlStatement, id_budgeting)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&bahan.Id_bahan, &bahan.Nama_bahan, &bahan.Jumlah_bahan, &bahan.Satuan_bahan, &bahan.Harga_bahan)
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
