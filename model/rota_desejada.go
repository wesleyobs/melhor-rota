package model

type RotaDesejada struct {
	Origem  string  `json:"origem"`
	Destino string  `json:"destino"`
	Valor   float64 `json:"valor"`
}
