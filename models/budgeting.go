package models

import (
	"KevinCatering/db"
	str "KevinCatering/struct"
	"KevinCatering/tools"
	"net/http"
	"strconv"
	"time"
)

//Generate_id_catering
func Generate_Id_Budgeting() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_budgeting FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_budgeting=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

func Generate_Id_Bahan() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_bahan FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_bahan=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

//Input_Budgeting
func Input_Budgeting(id_catering string, nama_menu string, total_porsi int, tanggal_budgeting string,
	nama_bahan string, jumlah_bahan string, satuan_bahan string, harga_bahan string) (tools.Response, error) {
	var res tools.Response
	var id str.Read_Id_Menu
	con := db.CreateCon()

	nm := Generate_Id_Budgeting()

	nm_str := strconv.Itoa(nm)

	id_BD := "BD-" + nm_str

	date, _ := time.Parse("02-01-2006", tanggal_budgeting)
	date_sql := date.Format("2006-01-02")

	sqlStatement := "INSERT INTO budgeting (id_budgeting,id_catering, nama_menu, total_porsi, tanggal_budgeting, id_bahan, nama_bahan, jumlah_bahan, satuan_bahan, harga_bahan, status_budgeting) values(?,?,?,?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	nb := tools.String_Separator_To_String(nama_bahan)

	id_bhn_fix := ""

	for i := 0; i < len(nb); i++ {
		nm := Generate_Id_Bahan()
		nm_str := strconv.Itoa(nm)
		id_Bhn := "B-" + nm_str
		id_bhn_fix = id_bhn_fix + "|" + id_Bhn + "|"
	}

	_, err = stmt.Exec(id_BD, id_catering, nama_menu, total_porsi, date_sql, id_bhn_fix,
		nama_bahan, jumlah_bahan, satuan_bahan, harga_bahan, 0)

	stmt.Close()

	sqlStatement = "SELECT id_bahan FROM budgeting WHERE id_budgeting=?"

	_ = con.QueryRow(sqlStatement, id_BD).Scan(&id.Id_menu)

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = id

	return res, nil
}

//Read_awal_budgeting
func Read_Budgeting_Awal(id_catering string) (tools.Response, error) {
	var res tools.Response
	var arr []str.Read_Budgeting_Awal
	var obj str.Read_Budgeting_Awal

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

//Read_budgeting
func Read_Budgeting(id_budgeting string) (tools.Response, error) {
	var res tools.Response
	var rm str.Read_Budgeting
	var obj str.Read_Bahan_fix
	var menu str.Read_Bahan
	var arr []str.Read_Budgeting

	con := db.CreateCon()

	sqlStatement := "SELECT id_budgeting, nama_menu, total_porsi, tanggal_budgeting, id_bahan, nama_bahan, jumlah_bahan, satuan_bahan, harga_bahan, status_budgeting FROM budgeting WHERE id_budgeting=? "

	err := con.QueryRow(sqlStatement, id_budgeting).Scan(&rm.Id_budgeting, &rm.Nama_menu,
		&rm.Total_porsi, &rm.Tanggal_budgeting, &menu.Id_bahan, &menu.Nama_bahan, &menu.Jumlah_bahan,
		&menu.Satuan_bahan, &menu.Harga_bahan, &rm.Status_budgeting)

	if err != nil {
		return res, err
	}

	id_mn_all := tools.String_Separator_To_String(menu.Id_bahan)
	nm_mn := tools.String_Separator_To_String(menu.Nama_bahan)
	hg_mn := tools.String_Separator_To_float64(menu.Jumlah_bahan)
	j_awal := tools.String_Separator_To_String(menu.Satuan_bahan)
	j_akhir := tools.String_Separator_To_Int(menu.Harga_bahan)

	for i := 0; i < len(id_mn_all); i++ {
		obj.Id_bahan = id_mn_all[i]
		obj.Nama_bahan = nm_mn[i]
		obj.Jumlah_bahan = hg_mn[i]
		obj.Satuan_bahan = j_awal[i]
		obj.Harga_bahan = j_akhir[i]
		rm.Bahan = append(rm.Bahan, obj)
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


