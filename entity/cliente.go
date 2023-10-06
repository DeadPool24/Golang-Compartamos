package entity

type Cliente struct {
	Dni        string `json:"dni"`
	Nombres    string `json:"nombres"`
	Apellidos  string `json:"apellidos"`
	Nacimiento string `json:"nacimiento"`
	Ciudad     string `json:"ciudad"`
	Direccion  string `json:"direccion"`
	Correo     string `json:"correo"`
	Telefono   int64  `json:"telefono"`
}
