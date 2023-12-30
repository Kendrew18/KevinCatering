package Menu

type Read_Menu_fix struct {
	Id_menu              string       `json:"id_menu"`
	Id_master_menu       string       `json:"id_master_menu"`
	Nama_menu            string       `json:"nama_menu"`
	Harga_menu           int64        `json:"harga_menu"`
	Jam_pengiriman_awal  string       `json:"jam_pengiriman_awal"`
	Jam_pengiriman_akhir string       `json:"jam_pengiriman_akhir"`
	Deskripsi_menu       string       `json:"deskripsi_menu"`
	Status_menu          int          `json:"status_menu"`
	Foto_menu            string       `json:"foto_menu"`
	Bahan_menu           []Bahan_Menu `json:"bahan_menu"`
}

type Read_Menu struct {
	Id_catering  string          `json:"id_catering"`
	Tanggal_menu string          `json:"tanggal_menu"`
	Menu         []Read_Menu_fix `json:"menu"`
}

type Bahan_Menu struct {
	Id_bahan_menu   string `json:"id_bahan_menu"`
	Nama_bahan_menu string `json:"nama_bahan_menu"`
}
