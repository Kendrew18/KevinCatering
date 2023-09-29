# How TO Use API Order
__________
NB: id pengantar isa di dapet dari login :)

## Show Menu

Link: kostsoda.onthewifi.com:3333/ORD/read-order-pengantar

Method: GET

Controllers:

    id_pengantar := c.FormValue("id_pengantar")

NB: ID-> bisa id user atupun id catering

## Order Detail User

Link: kostsoda.onthewifi.com:3333/ORD/order_detail_user

Method: GET

Controllers:

    id_detail_order := c.FormValue("id_detail_order")

## Update Maps Pengantar

Link: kostsoda.onthewifi.com:3333/PA/update-Maps-pengantar

Method: PUT

Controllers:

    id_pengantar := c.FormValue("id_pengantar")
	langtitude := c.FormValue("langtitude")
	longtitude := c.FormValue("longtitude")

## Read Maps Pengantar

Link: kostsoda.onthewifi.com:3333/PA/read-Maps-pengantar

Method: GET

Controllers:

    id_pengantar := c.FormValue("id_pengantar")

## Read Maps Pembeli

Link: kostsoda.onthewifi.com:3333/ORD/read-location-user

Method: GET

Controllers:

    id_detail_order := c.FormValue("id_detail_order")
