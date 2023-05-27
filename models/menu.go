package models

import (
	"KevinCatering/db"
	str "KevinCatering/struct"
	"KevinCatering/tools"
	"net/http"
	"strconv"
	"time"
)

func Generate_Id_Menu() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_menu FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_menu=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

//input menu
func Input_Menu(id_catering string, nama_menu string, harga_menu string, tanggal_menu string,
	jam_pengiriman_awal string, jam_pengiriman_akhir string, status string) (tools.Response, error) {
	var res tools.Response
	var id str.Read_Id_Menu
	con := db.CreateCon()

	nm := Generate_Id_Menu()

	nm_str := strconv.Itoa(nm)

	id_MN := "MN-" + nm_str

	date, _ := time.Parse("02-01-2006", tanggal_menu)
	date_sql := date.Format("2006-01-02")

	sqlStatement := "SELECT id_menu FROM menu WHERE tanggal_menu=? && id_catering=?"

	_ = con.QueryRow(sqlStatement, date_sql, id_catering).Scan(&id.Id_menu)

	if id.Id_menu == "" {

		id_mn_all := "|" + id_MN + "|"
		nm_mn := "|" + nama_menu + "|"
		hg_mn := "|" + harga_menu + "|"
		j_awal := "|" + jam_pengiriman_awal + "|"
		j_akhir := "|" + jam_pengiriman_akhir + "|"
		st := "|" + status + "|"

		sqlStatement := "INSERT INTO menu (id_catering, id_menu, nama_menu, harga_menu, tanggal_menu, jam_pengiriman_awal, jam_pengiriman_akhir, status_menu,foto_menu) values(?,?,?,?,?,?,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(id_catering, id_mn_all, nm_mn, hg_mn, date_sql, j_awal,
			j_akhir, st, "|uploads/images.png|")

		stmt.Close()

	} else {

		var menu str.Read_Menu_String

		sqlStatement := "SELECT id_menu, nama_menu, harga_menu, jam_pengiriman_awal, jam_pengiriman_akhir, status_menu, foto_menu FROM menu WHERE tanggal_menu=? && id_catering=? "

		_ = con.QueryRow(sqlStatement, date_sql, id_catering).Scan(&menu.Id_menu,
			&menu.Nama_menu, &menu.Harga_menu, &menu.Jam_pengiriman_awal,
			&menu.Jam_pengiriman_akhir, &menu.Status_menu)

		id_mn_all := menu.Id_menu + "|" + id_MN + "|"
		nm_mn := menu.Nama_menu + "|" + nama_menu + "|"
		hg_mn := menu.Harga_menu + "|" + harga_menu + "|"
		j_awal := menu.Jam_pengiriman_awal + "|" + jam_pengiriman_awal + "|"
		j_akhir := menu.Jam_pengiriman_akhir + "|" + jam_pengiriman_akhir + "|"
		st := menu.Status_menu + "|" + status + "|"

		sqlstatement := "UPDATE menu SET id_menu=?,nama_menu=?,harga_menu=?,jam_pengiriman_awal=?,jam_pengiriman_akhir=?,status_menu=? WHERE id_catering=? && tanggal_menu=?"

		stmt, err := con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(id_mn_all, nm_mn, hg_mn, j_awal, j_akhir, st, id_catering, date_sql)

		if err != nil {
			return res, err
		}

		stmt.Close()
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//read_menu
func Read_Menu(id_catering string, tanggal_menu string) (tools.Response, error) {
	var res tools.Response
	var rm str.Read_Menu
	var obj str.Read_Menu_fix
	var menu str.Read_Menu_String
	var arr []str.Read_Menu

	con := db.CreateCon()

	date, _ := time.Parse("02-01-2006", tanggal_menu)
	date_sql := date.Format("2006-01-02")

	sqlStatement := "SELECT id_catering,tanggal_menu, id_menu, nama_menu, harga_menu, jam_pengiriman_awal, jam_pengiriman_akhir, status_menu, foto_menu FROM menu WHERE tanggal_menu=? && id_catering=? "

	err := con.QueryRow(sqlStatement, date_sql, id_catering).Scan(&rm.Id_catering, &rm.Tanggal_menu,
		&menu.Id_menu, &menu.Nama_menu, &menu.Harga_menu, &menu.Jam_pengiriman_awal,
		&menu.Jam_pengiriman_akhir, &menu.Status_menu, &menu.Foto_menu)

	if err != nil {
		return res, err
	}

	id_mn_all := tools.String_Separator_To_String(menu.Id_menu)
	nm_mn := tools.String_Separator_To_String(menu.Nama_menu)
	hg_mn := tools.String_Separator_To_Int64(menu.Harga_menu)
	j_awal := tools.String_Separator_To_String(menu.Jam_pengiriman_awal)
	j_akhir := tools.String_Separator_To_String(menu.Jam_pengiriman_akhir)
	st := tools.String_Separator_To_Int(menu.Status_menu)

	for i := 0; i < len(id_mn_all); i++ {
		obj.Id_menu = id_mn_all[i]
		obj.Nama_menu = nm_mn[i]
		obj.Harga_menu = hg_mn[i]
		obj.Jam_pengiriman_awal = j_awal[i]
		obj.Jam_pengiriman_akhir = j_akhir[i]
		obj.Status_menu = st[i]
		rm.Menu = append(rm.Menu, obj)
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

//edit menu
func Edit_Menu(id_catering string, id_menu string, nama_menu string, harga_menu string,
	tanggal_menu string, jam_pengiriman_awal string, jam_pengiriman_akhir string) (tools.Response, error) {
	var res tools.Response
	var menu str.Read_Menu_String

	con := db.CreateCon()

	date, _ := time.Parse("02-01-2006", tanggal_menu)
	date_sql := date.Format("2006-01-02")

	sqlStatement := "SELECT id_menu, nama_menu, harga_menu, jam_pengiriman_awal, jam_pengiriman_akhir, status_menu, foto_menu FROM menu WHERE tanggal_menu=? && id_catering=? "

	err := con.QueryRow(sqlStatement, date_sql, id_catering).Scan(&menu.Id_menu,
		&menu.Nama_menu, &menu.Harga_menu, &menu.Jam_pengiriman_awal,
		&menu.Jam_pengiriman_akhir, &menu.Status_menu)

	if err != nil {
		return res, err
	}

	id_mn_all := tools.String_Separator_To_String(menu.Id_menu)
	nm_mn := tools.String_Separator_To_String(menu.Nama_menu)
	hg_mn := tools.String_Separator_To_String(menu.Harga_menu)
	j_awal := tools.String_Separator_To_String(menu.Jam_pengiriman_awal)
	j_akhir := tools.String_Separator_To_String(menu.Jam_pengiriman_akhir)
	st := tools.String_Separator_To_Int(menu.Status_menu)

	condition := -1

	for i := 0; i < len(id_mn_all); i++ {
		if id_menu == id_mn_all[i] {
			condition = st[i]
		}
	}

	if condition == 0 {

		nm_s := ""
		hg := ""
		jaw := ""
		jak := ""

		for i := 0; i < len(id_mn_all); i++ {
			if id_menu == id_mn_all[i] {
				nm_mn[i] = nama_menu
				hg_mn[i] = harga_menu
				j_awal[i] = jam_pengiriman_awal
				j_akhir[i] = jam_pengiriman_akhir
			}
			nm_s = nm_s + "|" + nm_mn[i] + "|"
			hg = hg + "|" + hg_mn[i] + "|"
			jaw = jaw + "|" + j_awal[i] + "|"
			jak = jak + "|" + j_akhir[i] + "|"
		}

		sqlstatement := "UPDATE menu SET nama_menu=?,harga_menu=?,jam_pengiriman_awal=?,jam_pengiriman_akhir=? WHERE id_catering=? && tanggal_menu=?"

		stmt, err := con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(nm_s, hg, jaw, jak, id_catering, tanggal_menu)

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
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Tidak Dapat Diupdate"
	}
	return res, nil
}

func Delete_Menu(id_catering string, id_menu string, tanggal_menu string) (tools.Response, error) {
	var res tools.Response
	var menu str.Read_Menu_String

	con := db.CreateCon()

	date, _ := time.Parse("02-01-2006", tanggal_menu)
	date_sql := date.Format("2006-01-02")

	sqlStatement := "SELECT id_menu, nama_menu, harga_menu, jam_pengiriman_awal, jam_pengiriman_akhir, status_menu, foto_menu FROM menu WHERE tanggal_menu=? && id_catering=? "

	err := con.QueryRow(sqlStatement, date_sql, id_catering).Scan(&menu.Id_menu,
		&menu.Nama_menu, &menu.Harga_menu, &menu.Jam_pengiriman_awal,
		&menu.Jam_pengiriman_akhir, &menu.Status_menu, &menu.Foto_menu)

	if err != nil {
		return res, err
	}

	id_mn_all := tools.String_Separator_To_String(menu.Id_menu)
	nm_mn := tools.String_Separator_To_String(menu.Nama_menu)
	hg_mn := tools.String_Separator_To_String(menu.Harga_menu)
	j_awal := tools.String_Separator_To_String(menu.Jam_pengiriman_awal)
	j_akhir := tools.String_Separator_To_String(menu.Jam_pengiriman_akhir)
	st := tools.String_Separator_To_Int(menu.Status_menu)
	st_s := tools.String_Separator_To_String(menu.Status_menu)
	foto := tools.String_Separator_To_String(menu.Foto_menu)

	condition := -1

	for i := 0; i < len(id_mn_all); i++ {
		if id_menu == id_mn_all[i] {
			condition = st[i]
		}
	}

	if condition == 0 {

		nm_s := ""
		hg := ""
		jaw := ""
		jak := ""
		id := ""
		st_st := ""
		ft := ""

		for i := 0; i < len(id_mn_all); i++ {
			if id_menu == id_mn_all[i] {

			} else {
				id = id + "|" + id_mn_all[i] + "|"
				nm_s = nm_s + "|" + nm_mn[i] + "|"
				hg = hg + "|" + hg_mn[i] + "|"
				jaw = jaw + "|" + j_awal[i] + "|"
				jak = jak + "|" + j_akhir[i] + "|"
				st_st = st_st + "|" + st_s[i] + "|"
				ft = ft + "|" + foto[i] + "|"

			}
		}

		sqlstatement := "UPDATE menu SET id_menu=?,nama_menu=?,harga_menu=?,jam_pengiriman_awal=?,jam_pengiriman_akhir=?,status_menu=?,foto_menu=? WHERE id_catering=? && tanggal_menu=?"

		stmt, err := con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(id, nm_s, hg, jaw, jak, st_st, ft, id_catering, tanggal_menu)

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
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Tidak Dapat Diupdate"
	}
	return res, nil
}
