package model

type NovaRota struct {
	Itinerario []string `json:"itinerario"`
	Valor      string   `json:"valor"`
}

func GetItinerarioComValor(rota NovaRota) []string {
	return append(rota.Itinerario, rota.Valor)
}
