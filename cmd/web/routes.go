package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Get("/pay", app.PaymentTerminal)
	mux.Post("/payment-successful", app.PaymentSuccessful)
	return mux
}
