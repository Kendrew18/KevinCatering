# Tambahan dan Perubahan
__________ 
##  Input Rating
Link: kostsoda.onthewifi.com:3333/RT/input-rating

Method: POST

Controllers:

    id_detail_order := c.FormValue("id_detail_order")
	id_catering := c.FormValue("id_catering")
	rating := c.FormValue("rating")
	review := c.FormValue("review")

##  History Order
Link: kostsoda.onthewifi.com:3333/ORD/history-order

Method: GET

Controllers:

    id_user := c.FormValue("id_user")

## Filter Catering
Link: kostsoda.onthewifi.com:3333/cat/filter-catering

Method: GET 

Controllers:

    id_user := c.FormValue("id_user")
    tipe := c.FormValue("tipe")

NB: tipe itu interger. angka 1 buat rating tertinggi, angka 2 buat favorite, angka 3 harian, angka 4 mingguan, angka 5 bulanan

## Favorite Catering
Link: kostsoda.onthewifi.com:3333/cat/favorite-catering

Method: POST

Controllers:

    id_user := c.FormValue("id_user")
    id_catering := c.FormValue("id_catering")

## Input Menu
Link: kostsoda.onthewifi.com:3333/mn/input-menu

Method: POST

Controllers:
    
    id_catering := c.FormValue("id_catering")
    nama_menu := c.FormValue("nama_menu")
    harga_menu := c.FormValue("harga_menu")
    tanggal_menu := c.FormValue("tanggal_menu")
    jam_pengiriman_awal := c.FormValue("jam_pengiriman_awal")
    jam_pengiriman_akhir := c.FormValue("jam_pengiriman_akhir")
    status := c.FormValue("status")
    file, handler, err := request.FormFile("photo")

NB: lek de e gak mau ngasih foto kosongi ae photo e :))

## Input Order
Link: kostsoda.onthewifi.com:3333/ORD/input-order

Method: POST

Controllers:

    id_user := c.FormValue("id_user")
    id_catering := c.FormValue("id_catering")
    id_menu := c.FormValue("id_menu")
    nama_menu := c.FormValue("nama_menu")
    jumlah_menu := c.FormValue("jumlah_menu")
    harga_menu := c.FormValue("harga_menu")
    tanggal_menu := c.FormValue("tanggal_menu")
    tanggal_order := c.FormValue("tanggal_order")
    longtitude := c.FormValue("longtitude")
    langtitude := c.FormValue("langtitude")
    file, handler, err := request.FormFile("photo")