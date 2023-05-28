package _struct

type Foto_Menu struct {
	Id_menu   string `json:"id_menu"`
	Path_Foto string `json:"path_foto"`
}

type Foto_Menu_Fix struct {
	Id_menu   []string `json:"id_menu"`
	Path_Foto []string `json:"path_foto"`
}
