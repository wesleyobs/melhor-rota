package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wesleyobs/melhor-rota/model"
	"github.com/wesleyobs/melhor-rota/service"
)

func MelhorRota(w http.ResponseWriter, r *http.Request) {
	rotaDesejada := model.RotaDesejada{
		Origem:  r.FormValue("origem"),
		Destino: r.FormValue("destino"),
	}
	melhorRotaEncontrada, erro := service.ObterMelhorRota(rotaDesejada)

	if naoTemErro(erro, w) {
		fmt.Fprintln(w, melhorRotaEncontrada)
	}

}

func NovaRota(w http.ResponseWriter, r *http.Request) {
	rota := model.NovaRota{}
	json.NewDecoder(r.Body).Decode(&rota)

	erro := service.GuardarNovaRotaCsv(rota)

	if naoTemErro(erro, w) {
		fmt.Fprintf(w, "Nova rota cadastrada com sucesso!")
	}

}

func naoTemErro(erro error, w http.ResponseWriter) bool {
	if erro != nil {
		fmt.Fprintf(w, erro.Error())
		return false
	}
	return true
}
