package Order

type Menu_Order_String struct {
	Id_menu              string `json:"id_menu"`
	Tanggal_menu         string `json:"tanggal_menu"`
	Nama_menu            string `json:"nama_menu"`
	Jumlah_menu          string `json:"jumlah_menu"`
	Harga_menu           string `json:"harga_menu"`
	Jam_pengiriman_awal  string `json:"jam_pengiriman_awal"`
	Jam_pengiriman_akhir string `json:"jam_pengiriman_akhir"`
	Status_order         string `json:"status_order"`
}

type Menu_Order struct {
	Id_menu              string `json:"id_menu"`
	Tanggal_menu         string `json:"tanggal_menu"`
	Nama_menu            string `json:"nama_menu"`
	Jumlah_menu          int    `json:"jumlah_menu"`
	Harga_menu           int64  `json:"harga_menu"`
	Jam_pengiriman_awal  string `json:"jam_pengiriman_awal"`
	Jam_pengiriman_akhir string `json:"jam_pengiriman_akhir"`
	Status_order         int    `json:"status_order"`
}

type Read_Detail_Order struct {
	Id_order              string       `json:"id_order"`
	Id_user               string       `json:"id_user"`
	Nama_user             string       `json:"nama_user"`
	Alamat_user           string       `json:"alamat_user"`
	No_telp_user          string       `json:"no_telp_user"`
	Id_catering           string       `json:"id_catering"`
	Nama_catering         string       `json:"nama_catering"`
	Alamat_catering       string       `json:"alamat_catering"`
	No_telp_catering      string       `json:"no_telp_catering"`
	Tanggal_order         string       `json:"tanggal_order"`
	Longtitude            float64      `json:"longtitude"`
	Langtitude            float64      `json:"langtitude"`
	Foto_Bukti_pembayaran string       `json:"foto_bukti_pembayaran"`
	Menu_order            []Menu_Order `json:"menu_order"`
	Total                 int64        `json:"total"`
}
