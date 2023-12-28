package Realisasi

type Read_Realisasi struct {
	Id_realisasi string `json:"id_realisasi"`
	Keterangan   string `json:"keterangan"`
	Harga_bahan  int    `json:"harga_bahan"`
	Jumlah_bahan int    `json:"jumlah_bahan"`
}

type Tabel_Realisasi struct {
	Id_bahan_menu     string  `json:"id_bahan_menu"`
	Nama_bahan        string  `json:"nama_bahan"`
	Total_bahan       float64 `json:"total_bahan"`
	Satuan            string  `json:"satuan"`
	Total_pengeluaran int     `json:"total_pengeluaran"`
}
