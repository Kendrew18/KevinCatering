package models

import (
	"KevinCatering/db"
	str "KevinCatering/struct"
	"KevinCatering/tools"
	"fmt"
	"net/http"
	"strconv"
)

//Generate_id_user
func Generate_Id_User() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_user FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_user=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

//sign UP
func Sign_up(nama_user string, telp_user string, email_user string,
	username_user string, password_user string,
	status_user int) (tools.Response, error) {
	var res tools.Response

	var arr str.Login

	con := db.CreateCon()

	sqlStatement := "SELECT id_user FROM user where username_user=?"

	_ = con.QueryRow(sqlStatement, username_user).Scan(&arr.Id_user)

	if arr.Id_user == "" {

		nm := Generate_Id_User()

		nm_str := strconv.Itoa(nm)

		id_US := "US-" + nm_str

		sqlStatement = "INSERT INTO user (id_user,nama,telp_user,email_user,username_user,password_user,foto_user,status_user) values(?,?,?,?,?,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(id_US, nama_user, telp_user, email_user, username_user, password_user, "uploads/images.png", status_user)

		stmt.Close()

		arr.Id_user = id_US
		arr.Status_user = status_user

		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
	}

	return res, nil
}

//login
func Login(username string, password string) (tools.Response, error) {
	var arr str.Login
	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "SELECT id_user,status_user FROM user where username_user=? && password_user=?"

	err := con.QueryRow(sqlStatement, username, password).Scan(&arr.Id_user, &arr.Status_user)

	if err != nil || arr.Id_user == "" {
		arr.Id_user = ""
		arr.Status_user = 0
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr

	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr
	}

	fmt.Println(arr)
	return res, nil
}

//see profile
func Read_Profile(id_user string) (tools.Response, error) {
	var res tools.Response
	var arr []str.Read_user
	var obj str.Read_user

	con := db.CreateCon()

	sqlStatement := "SELECT id_user, nama, telp_user, email_user, username_user, password_user, foto_user, status_user FROM user WHERE user.id_user=?"

	rows, err := con.Query(sqlStatement, id_user)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_user, &obj.Nama_user, &obj.Telp_user, &obj.Email_user,
			&obj.Username_user, &obj.Password_user, &obj.Foto_user, &obj.Status_user)
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

//edit profile
func Edit_Profile(id_user string, nama_user string, telp_user string,
	email_user string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "UPDATE user SET nama=?,telp_user=?,email_user=? WHERE id_user=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama_user, telp_user, email_user, id_user)

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
