# How TO Use API Pembayaran
__________ 
##  Read Pembayaran
Link: kostsoda.onthewifi.com:3333/PBR/read-pembayaran

Method: GET

Controllers:

    id_order := c.FormValue("id_order")

## Upload Foto Pembayaran
Link: kostsoda.onthewifi.com:3333/PBR/upload-foto

Method: POST

Controllers:

    id_order := c.FormValue("id_order")

## Confirm Pembayaran
Link: kostsoda.onthewifi.com:3333/PBR/confirm-pembayaran

Method: PUT

Controllers:

    id_order := c.FormValue("id_order")

## Read Recipe Order
Link: kostsoda.onthewifi.com:3333/PBR/read-recipe-order

Method: GET

Controllers:

    id_user := c.FormValue("id_user")

## Read Detail Rescipe order
Link: kostsoda.onthewifi.com:3333/PBR/read-detail-rescipe-order

Method: GET

Controllers:

    id_order := c.FormValue("id_order")

## Read Notif Pembayaran
Link: kostsoda.onthewifi.com:3333/PBR/read_notif_pembayaran

Method: GET

Controller:

    id_order := c.FormValue("id_order")

## Read QR Catering
Link: kostsoda.onthewifi.com:3333/cat/get-QR-catering

Method: GET

Contoller:

    id_catering := c.FormValue("id_catering")