package main

import (
	"fmt"
	"net/http"

	"events-ticketing/handlers"
	"events-ticketing/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	es := storage.NewEventSystem()
	h := handlers.NewHandlers(es)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/reserve", h.ReserveTicket)
	r.Get("/view", h.ViewTicketDetails)
	r.Get("/attendees", h.ViewAllAttendees)
	r.Post("/cancel", h.CancelReservation)
	r.Post("/modify", h.ModifySeatReservation)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", r)
}
