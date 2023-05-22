package _struct

type Tipe_Catering struct {
	Tipe_catering string `json:"tipe_catering"`
}
type Read_Awal_Catering struct {
	Id_catering     string   `json:"id_catering"`
	Nama_catering   string   `json:"nama_catering"`
	Alamat_catering string   `json:"alamat_catering"`
	Tipe_pemesanan  []string `json:"tipe_pemesanan"`
	Rating          float32  `json:"rating"`
}
