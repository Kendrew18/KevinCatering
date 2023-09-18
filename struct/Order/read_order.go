package Order

type Read_Order struct {
	Id_order      string `json:"id_order"`
	Nama_catering string `json:"nama_catering"`
	Tanggal_order string `json:"tanggal_order"`
}

type Read_Id_Order struct {
	Id_Order string `json:"id_order"`
}

type Read_Menu_Order struct {
	Tanggal_menu       string               `json:"tanggal_menu"`
	Menu_Order_Dipesan []Menu_Order_Dipesan `json:"menu_order_dipesan"`
}
type Menu_Order_Dipesan struct {
	Id_order        string `json:"id_order"`
	Id_detail_order string `json:"id_detail_order"`
	Id_catering     string `json:"id_catering"`
	Nama_catering   string `json:"nama_catering"`
	Id_pengantar    string `json:"id_pengantar"`
	Nama_menu       string `json:"nama_menu"`
	Harga_menu      string `json:"harga_menu"`
	Status_order    string `json:"status_order"`
}

type Read_Detail_Order_User struct {
	Nama_pengantar string `json:"nama_pengantar"`
	Status         string `json:"status"`
	Nomer_telp     string `json:"nomer_telp"`
	Nama_catering  string `json:"nama_catering"`
	Nama_menu      string `json:"nama_menu"`
	Jumlah         int    `json:"jumlah"`
	Harga          int    `json:"harga"`
}
