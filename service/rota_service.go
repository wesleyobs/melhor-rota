package service

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/wesleyobs/melhor-rota/model"
)

const valor int = 1

var csvPath = "../resources/" + os.Getenv("ENVIRONMENT_APP") + "/destinos.csv"

func GuardarNovaRotaCsv(novaRota model.NovaRota) error {
	arquivo, err := os.OpenFile(csvPath, os.O_APPEND|os.O_WRONLY, 0777)

	if err != nil {
		return errors.New("Erro ao guardar uma nova rota: " + err.Error())
	}

	e := csv.NewWriter(arquivo)

	itinerarioComValor := model.GetItinerarioComValor(novaRota)

	e.Write(itinerarioComValor)

	e.Flush()
	arquivo.Close()

	return nil
}

func ObterMelhorRota(rotaDesejada model.RotaDesejada) (string, error) {
	arquivoDiretorio, erro := os.Open(csvPath)

	if erro != nil {
		return "", errors.New("Erro ao obter a melhor rota: " + erro.Error())
	}

	arquivo := csv.NewReader(arquivoDiretorio)
	procurarRota := new(model.ProcurarRota)
	procurarRota.SetRotaDesejada(rotaDesejada)

	for {
		arrayLinhaCsv, erro := arquivo.Read()
		if erro == io.EOF {
			break
		}

		if procurarRota.RotaEncontrada(arrayLinhaCsv) {
			valorPassagem, erro := strconv.ParseFloat(arrayLinhaCsv[len(arrayLinhaCsv)-valor], 64)

			procurarRota.SetValorPassagem(valorPassagem)

			if erro != nil {
				return "", errors.New("Erro ao converter o valor do itinerario para double: " + erro.Error())
			}

			if procurarRota.PrimeiraVarredura() ||
				procurarRota.ValorPassagemTemMenorPrecoAteOMomento() {

				procurarRota.SetMelhorRotaEncontrada(arrayLinhaCsv)
				procurarRota.SetMelhorPrecoEncontrado(valorPassagem)
			}
		}
	}

	return procurarRota.RotaFormatada(), nil

}
