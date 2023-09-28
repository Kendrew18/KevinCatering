# How TO Use API Budgeting dan realisasi
___
##  Input Budgeting (APK catering)

Link: kostsoda.onthewifi.com:3333/BD/input-budgeting

Method: POST

Controllers:

    id_catering := c.FormValue("id_catering")
	nama_menu := c.FormValue("nama_menu")
	total_porsi := c.FormValue("total_porsi")
	tanggal_budgeting := c.FormValue("tanggal_budgeting")
	nama_bahan := c.FormValue("nama_bahan")
	jumlah_bahan := c.FormValue("jumlah_bahan")
	satuan_bahan := c.FormValue("satuan_bahan")
	harga_bahan := c.FormValue("harga_bahan")

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

##  Input Budgeting (APK catering)

Link: kostsoda.onthewifi.com:3333/RL/input-realisasi

Method: POST

Controllers:

    id_bahan_menu := c.FormValue("id_bahan_menu")
	keterangan := c.FormValue("keterangan")
	jumlah := c.FormValue("jumlah")
	harga := c.FormValue("harga")

##  Read Awal Budgeting (APK catering)

Link: kostsoda.onthewifi.com:3333/RL/read-realisasi

Method: GET

Controllers:

    id_bahan_menu := c.FormValue("id_bahan_menu")