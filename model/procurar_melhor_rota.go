package model

import (
	"regexp"
	"strings"
)

const origem int = 0
const destino int = 2
const tamanhoMinimoArrayLinhaCsv int = 2

type ProcurarRota struct {
	RotaDesejada          RotaDesejada
	MelhorPrecoEncontrado float64
	ValorPassagem         float64
	MelhorRotaEncontrada  string
}

func (procurarRota *ProcurarRota) SetValorPassagem(valorPassagem float64) {
	procurarRota.ValorPassagem = valorPassagem
}

func (procurarRota *ProcurarRota) SetMelhorRotaEncontrada(arrayLinhaCsv []string) {
	procurarRota.MelhorRotaEncontrada = strings.Join(arrayLinhaCsv, " - ")
}

func (procurarRota *ProcurarRota) SetMelhorPrecoEncontrado(melhorPrecoEncontrado float64) {
	procurarRota.MelhorPrecoEncontrado = melhorPrecoEncontrado
}

func (procurarRota *ProcurarRota) SetRotaDesejada(rotaDesejada RotaDesejada) {
	procurarRota.RotaDesejada = rotaDesejada
}

func (procurarRota *ProcurarRota) RotaEncontrada(arrayLinhaCsv []string) bool {
	rotaDesejada := procurarRota.RotaDesejada

	if len(arrayLinhaCsv) >= tamanhoMinimoArrayLinhaCsv &&
		rotaDesejada.Origem == arrayLinhaCsv[origem] &&
		rotaDesejada.Destino == arrayLinhaCsv[len(arrayLinhaCsv)-destino] {
		return true
	}
	return false
}

func (procurarRota *ProcurarRota) PrimeiraVarredura() bool {
	return procurarRota.MelhorPrecoEncontrado == 0.0
}

/**
* retorna verdadeiro caso o valor da passagem da linha corrente do arquivo for menor que o valor da passagem encontrada anteriormente
 */
func (procurarRota *ProcurarRota) ValorPassagemTemMenorPrecoAteOMomento() bool {
	return procurarRota.ValorPassagem < procurarRota.MelhorPrecoEncontrado
}

func (procurarRota *ProcurarRota) RotaFormatada() string {
	if procurarRota.MelhorRotaEncontrada == "" {
		return "Rota não encontrada! Tente novamente."
	}

	var regexCompile = regexp.MustCompile("-([^-]*)$")
	return "Melhor opção encontrada " + strings.ReplaceAll(
		regexCompile.ReplaceAllString(procurarRota.MelhorRotaEncontrada, ">$1"),
		"> ",
		"> $")
}
