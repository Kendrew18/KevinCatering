# How TO Use API Menu
__________ 

##  Input Menu

Method: POST

Link: kostsoda.onthewifi.com:3333/mn/input-menu

Controller:

    id_catering := c.FormValue("id_catering")
	id_master_menu := c.FormValue("id_master_menu")
	harga_menu := c.FormValue("harga_menu")
	tanggal_menu := c.FormValue("tanggal_menu")
	jam_pengiriman_awal := c.FormValue("jam_pengiriman_awal")
	jam_pengiriman_akhir := c.FormValue("jam_pengiriman_akhir")
	status := c.FormValue("status")
    file, handler, err := request.FormFile("photo")

##  Read Menu

Method: GET

Link: kostsoda.onthewifi.com:3333/mn/read-menu

Controller:

    id_catering := c.FormValue("id_catering")
	tanggal_menu := c.FormValue("tanggal_menu")
	tanggal_menu2 := c.FormValue("tanggal_menu2")

##  Edit Menu

Method: PUT

Link: kostsoda.onthewifi.com:3333/mn/update-menu

Controller:

    id_catering := c.FormValue("id_catering")
	id_menu := c.FormValue("id_menu")
	harga_menu := c.FormValue("harga_menu")
	jam_pengiriman_awal := c.FormValue("jam_pengiriman_awal")
	jam_pengiriman_akhir := c.FormValue("jam_pengiriman_akhir")

##  Delete Menu

Method: DELETE

Link: kostsoda.onthewifi.com:3333/mn/delete-menu

Controller:

    id_catering := c.FormValue("id_catering")
    id_menu := c.FormValue("id_menu")

##  Upload Foto Menu

Method: POST

Link: kostsoda.onthewifi.com:3333/mn/upload-foto-menu

Controller:

    id_catering := c.FormValue("id_catering")
	id_menu := c.FormValue("id_menu")
    file, handler, err := request.FormFile("photo")