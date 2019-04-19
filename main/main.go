package main

import (
	"log"
	"net/http"

	"github.com/wesleyobs/melhor-rota/controller"
)

func main() {

	http.Handle("/rota/melhor-rota", http.HandlerFunc(handler))
	http.Handle("/rota/nova-rota", http.HandlerFunc(handler))
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		controller.MelhorRota(w, r)
	} else if r.Method == http.MethodPost {
		controller.NovaRota(w, r)
	}
}
