package Budgeting

type Read_Budgeting_Awal struct {
	Id_budgeting string `json:"id_budgeting"`
	Nama_menu    string `json:"id_menu"`
}

type Read_Budgeting struct {
	Id_budgeting      string       `json:"id_budgeting"`
	Nama_menu         string       `json:"nama_menu"`
	Total_porsi       int          `json:"total_porsi"`
	Tanggal_budgeting string       `json:"tanggal_budgeting"`
	Bahan             []Read_Bahan `json:"bahan"`
}
type Read_Bahan struct {
	Id_bahan     string  `json:"id_bahan"`
	Nama_bahan   string  `json:"nama_bahan"`
	Jumlah_bahan float64 `json:"jumlah_bahan"`
	Satuan_bahan string  `json:"satuan"`
	Harga_bahan  int     `json:"harga_bahan"`
}
