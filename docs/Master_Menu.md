# How TO Use API Master Menu
__________ 

##  Input Master Menu

Method: POST

Link: kostsoda.onthewifi.com:3333/MM/master-menu

Controller:

    id_catering := c.FormValue("id_catering")
	nama_menu := c.FormValue("nama_menu")
	deskripsi_menu := c.FormValue("deskripsi_menu")
	nama_bahan := c.FormValue("nama_bahan")
	jumlah_bahan := c.FormValue("jumlah_bahan")
	satuan_bahan := c.FormValue("satuan_bahan")
	harga_bahan := c.FormValue("harga_bahan")

NB: Nama bahan, jumlah bahan, satuan bahan dan harga bahan merupakan string separator.

##  Read Master Menu

Method: GET

Link: kostsoda.onthewifi.com:3333/MM/master-menu

Controller:
    
    id_catering := c.FormValue("id_catering")

##  Drop Down Master Menu

Method: GET

Link: kostsoda.onthewifi.com:3333/MM/dropdown

Controller:

    id_catering := c.FormValue("id_catering")