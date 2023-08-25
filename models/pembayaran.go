package models

import (
	"KevinCatering/db"
	str "KevinCatering/struct"
	"KevinCatering/tools"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//upload-foto-pembayaran
func Upload_Foto_Pembayaran(id_order string, writer http.ResponseWriter, request *http.Request) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	request.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err := request.FormFile("photo")
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	defer file.Close()

	fmt.Println("File Info")
	fmt.Println("File Name : ", handler.Filename)
	fmt.Println("File Size : ", handler.Size)
	fmt.Println("File Type : ", handler.Header.Get("Content-Type"))

	var tempFile *os.File
	path := ""

	if strings.Contains(handler.Filename, "jpg") {
		path = "uploads/" + id_order + ".jpg"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.jpg")
	}
	if strings.Contains(handler.Filename, "jpeg") {
		path = "uploads/" + id_order + ".jpeg"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.jpeg")
	}
	if strings.Contains(handler.Filename, "png") {
		path = "uploads/" + id_order + ".png"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.png")
	}

	if err != nil {
		return res, err
	}

	fileBytes, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		return res, err2
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return res, err
	}

	fmt.Println("Success!!")
	fmt.Println(tempFile.Name())
	tempFile.Close()

	err = os.Rename(tempFile.Name(), path)
	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	fmt.Println("new path:", tempFile.Name())

	sqlstatement := "UPDATE pembayaran SET bukti_pembayaran=? WHERE id_order=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(path, id_order)

	if err != nil {
		return res, err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read_Pembayaran
func Read_Pembayaran(id_order string) (tools.Response, error) {
	var res tools.Response
	var arr []str.Read_pembayaran
	var obj str.Read_pembayaran

	con := db.CreateCon()

	sqlStatement := "SELECT id_pembayaran,id_order,bukti_pembayaran,status_pembayaran FROM pembayaran WHERE id_order=?"

	fmt.Println(id_order)

	rows, err := con.Query(sqlStatement, id_order)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_pembayaran, &obj.Id_order, &obj.Bukti_pembayaran, &obj.Status_pembayaran)
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
