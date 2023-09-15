package Realisasi

type Read_Realisasi struct {
	Id_realisasi string `json:"id_realisasi"`
	Keterangan   string `json:"keterangan"`
	Harga_bahan  int    `json:"harga_bahan"`
	Jumlah_bahan int    `json:"jumlah_bahan"`
}
