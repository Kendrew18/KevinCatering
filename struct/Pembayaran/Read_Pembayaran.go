package Pembayaran

type Read_pembayaran struct {
	Id_pembayaran     string `json:"id_pembayaran"`
	Id_order          string `json:"id_order"`
	Bukti_pembayaran  string `json:"bukti_pembayaran"`
	Status_pembayaran int    `json:"status_pembayaran"`
}

type Read_Pembayaran_Fix struct {
	Id_pembayaran     string `json:"id_pembayaran"`
	Id_order          string `json:"id_order"`
	Status_pembayaran int    `json:"status_pembayaran"`
}
