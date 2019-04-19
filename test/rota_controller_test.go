package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/wesleyobs/melhor-rota/controller"
)

var csvPath = "../resources/" + os.Getenv("ENVIRONMENT_APP") + "/destinos.csv"

func TestMain(m *testing.M) {
	inserirItinerarioArquivo()
	retCode := m.Run()
	os.Exit(retCode)
	os.Truncate(csvPath, 0)
}

func TestObterMelhorRota(t *testing.T) {
	requestBuscarMelhorRota, _ := http.NewRequest("GET", "localhost:8001/rota/melhor-rota?origem=A&destino=D", nil)

	recorderBuscarMelhorRota := httptest.NewRecorder()

	controller.MelhorRota(recorderBuscarMelhorRota, requestBuscarMelhorRota)

	result := recorderBuscarMelhorRota.Result()
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		t.Errorf("Esperava um status code 200, mas foi obtido: %v", result.Status)
	}

	body, erro := ioutil.ReadAll(result.Body)
	if erro != nil {
		t.Fatalf("Não foi possível obter a resposta")
	}

	rotaEsperada := strings.TrimSpace("Melhor opção encontrada A - B - C - D > $17.00")

	resultadoObtido := strings.TrimSpace(string(body))
	if !strings.EqualFold(resultadoObtido, rotaEsperada) {
		t.Fatalf("Resultado esperado: %v - Resultado obtido: %v ", rotaEsperada, resultadoObtido)
	}
}

func inserirItinerarioArquivo() {
	os.Truncate(csvPath, 0)

	requestNovaRota1, _ := http.NewRequest("POST", "localhost:8001/rota/nova-rota", strings.NewReader(`{"itinerario": ["A","B","C","D"], "valor":"20.00"}`))
	requestNovaRota2, _ := http.NewRequest("POST", "localhost:8001/rota/nova-rota", strings.NewReader(`{"itinerario": ["A","A1","A2","D"], "valor":"19.00"}`))
	requestNovaRota3, _ := http.NewRequest("POST", "localhost:8001/rota/nova-rota", strings.NewReader(`{"itinerario": ["A","A3","A4","D"], "valor":"18.00"}`))
	requestNovaRota4, _ := http.NewRequest("POST", "localhost:8001/rota/nova-rota", strings.NewReader(`{"itinerario": ["A","B1","B2","A"], "valor":"03.00"}`))
	requestNovaRota5, _ := http.NewRequest("POST", "localhost:8001/rota/nova-rota", strings.NewReader(`{"itinerario": ["A","B","C","D"], "valor":"17.00"}`))

	recorderNovaRota := httptest.NewRecorder()
	controller.NovaRota(recorderNovaRota, requestNovaRota1)
	controller.NovaRota(recorderNovaRota, requestNovaRota2)
	controller.NovaRota(recorderNovaRota, requestNovaRota3)
	controller.NovaRota(recorderNovaRota, requestNovaRota4)
	controller.NovaRota(recorderNovaRota, requestNovaRota5)
}
