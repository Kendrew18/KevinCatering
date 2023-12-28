# How TO Use API Budgeting dan realisasi
___
##  Input Budgeting (APK catering)

Link: kostsoda.onthewifi.com:3333/BD/input-budgeting

Method: POST

Controllers:

    id_catering := c.FormValue("id_catering")
	id_master_menu := c.FormValue("id_master_menu")
	total_porsi := c.FormValue("total_porsi")
	tanggal_budgeting := c.FormValue("tanggal_budgeting")

##  Read Awal Budgeting (APK catering)

Link: kostsoda.onthewifi.com:3333/BD/read-awal-budgeting

Method: GET

Controllers:

    id_catering := c.FormValue("id_catering")

##  Read Detail Budgeting (APK catering)

Link: kostsoda.onthewifi.com:3333/BD/read-budgeting

Method: GET

Controllers:

    id_budgeting := c.FormValue("id_budgeting")

##  Update Status (APK catering)

Link: kostsoda.onthewifi.com:3333/BD/update-status

Method: PUT

Controllers:

    id_budgeting := c.FormValue("id_budgeting")

##  Read Tabel Realisasi (APK catering)

Link: kostsoda.onthewifi.com:3333/RL/read-tabel-realisasi

Method: GET

Controllers:

    id_budgeting := c.FormValue("id_budgeting")
    id_master_menu := c.FormValue("id_master_menu")

##  Input Realisasi (APK catering)

Link: kostsoda.onthewifi.com:3333/RL/input-realisasi

Method: POST

Controllers:

    id_budgeting := c.FormValue("id_budgeting")
	id_bahan_menu := c.FormValue("id_bahan_menu")
	keterangan := c.FormValue("keterangan")
	jumlah := c.FormValue("jumlah")
	harga := c.FormValue("harga")

##  Read Realisasi (APK catering)

Link: kostsoda.onthewifi.com:3333/RL/read-realisasi

Method: GET

Controllers:

    id_budgeting := c.FormValue("id_budgeting")
	id_bahan_menu := c.FormValue("id_bahan_menu")