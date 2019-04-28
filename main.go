package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	log.Fatal(http.ListenAndServe(":7878", r))
}
