package Master_Menu

import (
	"KevinCatering/db"
	"KevinCatering/struct/Master_Menu"
	"KevinCatering/tools"
	"net/http"
	"strconv"
)

func Input_Master_Menu(id_catering string, nama_menu string, deskripsi_menu string, nama_bahan string, jumlah_bahan string, satuan_bahan string, harga_bahan string) (tools.Response, error) {
	var res tools.Response
	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM master_menu ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_MM := "MM-" + strconv.Itoa(nm_str)

	sqlStatement := "INSERT INTO master_menu (co, id_master_menu, id_catering, nama_menu, deskripsi_menu) values(?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_MM, id_catering, nama_menu, deskripsi_menu)

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

		Sqlstatement_BHN = "INSERT INTO bahan_menu (co, id_bahan_menu, id_master_menu, nama_bahan, jumlah_bahan, satuan_bahan, harga_bahan) values(?,?,?,?,?,?,?)"

		stmt_BHN, err := con.Prepare(Sqlstatement_BHN)

		if err != nil {
			return res, err
		}

		_, err = stmt_BHN.Exec(nm_str_BHN, id_BHN, id_MM, nama_bahan_SS[i], jumlah_bahan_SS[i], satuan_bahan_SS[i], harga_bahan_SS[i])

		if err != nil {
			return res, err
		}
	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

func Read_Master_Menu(id_catering string) (tools.Response, error) {
	var res tools.Response

	var arr []Master_Menu.Read_Master_Menu
	var obj Master_Menu.Read_Master_Menu

	con := db.CreateCon()

	sqlStatement := "SELECT id_master_menu, id_catering, nama_menu, deskripsi_menu FROM master_menu WHERE id_catering=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_catering)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		var arr_detail []Master_Menu.Detail_Bahan_Master_Menu
		var obj_detail Master_Menu.Detail_Bahan_Master_Menu

		err = rows.Scan(&obj.Id_master_menu, &obj.Id_catering, &obj.Nama_menu, &obj.Deskripsi_menu)

		if err != nil {
			return res, err
		}

		sqlstatement_detail := "SELECT id_bahan_menu, nama_bahan, jumlah_bahan, satuan_bahan, harga_bahan FROM bahan_menu WHERE id_master_menu = ? ORDER BY co ASC "

		rows_detail, err := con.Query(sqlstatement_detail, obj.Id_master_menu)

		defer rows_detail.Close()

		for rows_detail.Next() {
			err = rows_detail.Scan(&obj_detail.Id_bahan_menu, &obj_detail.Nama_bahan, &obj_detail.Jumlah_bahan, &obj_detail.Satuan_bahan, &obj_detail.Harga_bahan)

			if err != nil {
				return res, err
			}

			arr_detail = append(arr_detail, obj_detail)
		}

		obj.Detail_bahan_master_menu = arr_detail

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

func Drop_Down_Master_Menu(id_catering string) (tools.Response, error) {
	var res tools.Response
	var arr []Master_Menu.Drop_Down_Master_Menu
	var obj Master_Menu.Drop_Down_Master_Menu

	con := db.CreateCon()

	sqlStatement := "SELECT id_master_menu,nama_menu FROM master_menu WHERE id_catering=? ORDER BY co DESC "

	rows, err := con.Query(sqlStatement, id_catering)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_master_menu, &obj.Nama_menu)
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
