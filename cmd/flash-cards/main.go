package main

import (
	"net/http"

	"github.com/MauricioSauza/flash-cards/internal/routes"
)

func main() {
	r := routes.NewRouter()
	http.ListenAndServe(":8080", r)
}
