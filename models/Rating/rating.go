package Rating

import (
	"KevinCatering/db"
	"KevinCatering/tools"
	"net/http"
	"strconv"
)

func Input_Rating(id_detail_order string, id_catering string, rating int, review string) (tools.Response, error) {
	var res tools.Response
	var avg float64

	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM rating ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_RT := "RT-" + strconv.Itoa(nm_str)

	Sqlstatement = "INSERT INTO rating (co, id_rating, id_detail_order, id_catering, rating, review) values(?,?,?,?,?,?)"

	stmt, err := con.Prepare(Sqlstatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_RT, id_detail_order, id_catering, rating, review)

	if err != nil {
		return res, err
	}

	Sqlstatement = "SELECT AVG(rating) FROM rating WHERE id_catering=?"

	err = con.QueryRow(Sqlstatement, id_catering).Scan(&avg)

	if err != nil {
		return res, err
	}

	sqlstatement := "UPDATE catering SET rating=? WHERE id_catering=?"

	stmt, err = con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(avg, id_catering)

	if err != nil {
		return res, err
	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}
