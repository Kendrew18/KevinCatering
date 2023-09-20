package _struct

type Foto_Menu struct {
	Id_menu   string `json:"id_menu"`
	Path_Foto string `json:"path_foto"`
}

type Foto_Menu_Fix struct {
	Id_menu   []string `json:"id_menu"`
	Path_Foto []string `json:"path_foto"`
}

type Read_Path_Foto_QR struct {
	Id_catering string `json:"id_catering"`
	Path_QR     string `json:"path_qr"`
}
