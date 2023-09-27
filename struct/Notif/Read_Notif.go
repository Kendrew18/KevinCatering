package Notif

type Read_Notif struct {
	Id_notif    string `json:"id_notif"`
	Id_catering string `json:"id_catering"`
	Id_order    string `json:"id_order"`
	Catatan     string `json:"catatan"`
}

type Read_Detail_Notif struct {
	Path_Foto        string             `json:"path_foto"`
	Menu_Order_Notif []Menu_Order_Notif `json:"menu_order_notif"`
}

type Menu_Order_Notif struct {
	Id_menu              string `json:"id_menu"`
	Tanggal_menu         string `json:"tanggal_menu"`
	Nama_menu            string `json:"nama_menu"`
	Jumlah_menu          int    `json:"jumlah_menu"`
	Harga_menu           int64  `json:"harga_menu"`
	Jam_pengiriman_awal  string `json:"jam_pengiriman_awal"`
	Jam_pengiriman_akhir string `json:"jam_pengiriman_akhir"`
	Status_order         string `json:"status_order"`
}
