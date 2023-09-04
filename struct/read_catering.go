package _struct

type Read_Catering struct {
	Id_catering          string   `json:"id_catering"`
	Nama_catering        string   `json:"nama_catering"`
	Alamat_catering      string   `json:"alamat_catering"`
	Telp_catering        string   `json:"telp_catering"`
	Email_catering       string   `json:"email_catering"`
	Deskripsi_catering   string   `json:"deskripsi_catering"`
	Tipe_pemesanan       []string `json:"tipe_pemesanan"`
	Foto_profil_catering string   `json:"foto_profil_catering"`
	Rating               float32  `json:"rating"`
	Longtitude           string   `json:"longtitude"`
	Langtitude           string   `json:"langtitude"`
}
