# How TO Use API Order
__________
##  Input Order (APK user)

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

## Show Menu (APK user and catering)

Link: kostsoda.onthewifi.com:3333/ORD/show-order-menu

Method: GET

Controllers:

    id := c.FormValue("id")

NB: ID-> bisa id user atupun id catering

## Set Pegantar (APK catering)

Link: kostsoda.onthewifi.com:3333/ORD/set-pengantar

Method: PUT

Controllers:

    id_detail_order := c.FormValue("id_detail_order")
    id_pengantar := c.FormValue("id_pengantar")

## Confirm Order (APK user and Pengantar)

Link: kostsoda.onthewifi.com:3333/ORD/confirm-order

Method: PUT

Controllers:

    id := c.FormValue("id")
	id_detail_order := c.FormValue("id_detail_order")

NB: ID-> bisa id user atupun id pengantar

## Order Detail User

Link: kostsoda.onthewifi.com:3333/ORD/order_detail_user

Method: GET

Controllers:

    id_detail_order := c.FormValue("id_detail_order")
