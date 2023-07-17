package _struct

type Read_Budgeting struct {
	Id_budgeting      string           `json:"id_budgeting"`
	Nama_menu         string           `json:"nama_menu"`
	Total_porsi       int              `json:"total_porsi"`
	Tanggal_budgeting string           `json:"tanggal_budgeting"`
	Status_budgeting  string           `json:"status_budgeting"`
	Bahan             []Read_Bahan_fix `json:"bahan"`
}
type Read_Bahan_fix struct {
	Id_bahan     string  `json:"id_bahan"`
	Nama_bahan   string  `json:"nama_bahan"`
	Jumlah_bahan float64 `json:"jumlah_bahan"`
	Satuan_bahan string  `json:"satuan"`
	Harga_bahan  int     `json:"harga_bahan"`
}
type Read_Bahan struct {
	Id_bahan     string `json:"id_bahan"`
	Nama_bahan   string `json:"nama_bahan"`
	Jumlah_bahan string `json:"jumlah_bahan"`
	Satuan_bahan string `json:"satuan"`
	Harga_bahan  string `json:"harga_bahan"`
}
