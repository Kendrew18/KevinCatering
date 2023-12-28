package Master_Menu

type Read_Master_Menu struct {
	Id_master_menu           string                     `json:"id_master_menu"`
	Id_catering              string                     `json:"id_catering"`
	Nama_menu                string                     `json:"nama_menu"`
	Deskripsi_menu           string                     `json:"deskripsi_menu"`
	Detail_bahan_master_menu []Detail_Bahan_Master_Menu `json:"detail_bahan_master_menu"`
}

type Detail_Bahan_Master_Menu struct {
	Id_bahan_menu string  `json:"id_bahan_menu"`
	Nama_bahan    string  `json:"nama_bahan"`
	Jumlah_bahan  float64 `json:"jumlah_bahan"`
	Satuan_bahan  string  `json:"satuan_bahan"`
	Harga_bahan   int64   `json:"harga_bahan"`
}

type Drop_Down_Master_Menu struct {
	Id_master_menu string `json:"id_master_menu"`
	Nama_menu      string `json:"nama_menu"`
}
