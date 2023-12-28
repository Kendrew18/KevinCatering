package models

import (
	"KevinCatering/db"
	"KevinCatering/struct/Realisasi"
	"KevinCatering/tools"
	"net/http"
	"strconv"
)

//Read_Tabel_Realisasi
func Read_Tabel_Realisasi(id_budgeting string, id_master_menu string) (tools.Response, error) {
	var res tools.Response
	var arr []Realisasi.Tabel_Realisasi
	var obj Realisasi.Tabel_Realisasi

	con := db.CreateCon()

	sqlStatement := "SELECT id_bahan_menu, nama_bahan, satuan_bahan FROM bahan_menu WHERE id_master_menu=?"

	rows, err := con.Query(sqlStatement, id_master_menu)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_bahan_menu, &obj.Nama_bahan, &obj.Satuan)
		if err != nil {
			return res, err
		}

		sql_statement_total := "SELECT ifnull(sum(jumlah_bahan),0), ifnull(sum(harga_bahan),0) FROM realisasi WHERE id_bahan_menu=? && id_budgeting=?"

		err = con.QueryRow(sql_statement_total, obj.Id_bahan_menu, id_budgeting).Scan(&obj.Total_bahan, &obj.Total_pengeluaran)

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

//Input_Realisasi
func Input_Realisasi(id_budgeting string, id_bahan_menu string, keterangan string, jumlah_bahan float64, harga_bahan int) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM realisasi ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_RL := "RL-" + strconv.Itoa(nm_str)

	sqlStatement := "INSERT INTO realisasi (co, id_realisasi, id_budgeting, id_bahan_menu, keterangan, harga_bahan, jumlah_bahan) values(?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_RL, id_budgeting, id_bahan_menu, keterangan, harga_bahan, jumlah_bahan)

	if err != nil {
		return res, err
	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read_Realisasi
func Read_Realisasi(id_bahan_menu string) (tools.Response, error) {
	var res tools.Response
	var arr []Realisasi.Read_Realisasi
	var obj Realisasi.Read_Realisasi

	con := db.CreateCon()

	sqlStatement := "SELECT id_realisasi, keterangan, harga_bahan, jumlah_bahan FROM realisasi WHERE id_bahan_menu=?"

	rows, err := con.Query(sqlStatement, id_bahan_menu)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_realisasi, &obj.Keterangan, &obj.Harga_bahan, &obj.Jumlah_bahan)
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
