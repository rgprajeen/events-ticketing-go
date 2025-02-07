package handlers

import (
	"encoding/json"
	"events-ticketing/models"
	"events-ticketing/storage"
	"math/rand"
	"net/http"
	"time"
)

type Handlers struct {
	es *storage.EventSystem
}

func NewHandlers(es *storage.EventSystem) *Handlers {
	return &Handlers{es: es}
}

func (h *Handlers) ReserveTicket(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	err := json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ticket.SeatNumber = rand.Intn(100) + 1
	h.es.AddTicket(ticket)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Ticket reserved successfully",
		"ticket":  ticket,
	})
	w.WriteHeader(http.StatusCreated)
}

func (h *Handlers) ViewTicketDetails(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	userTickets := h.es.GetTicketsByEmail(email)
	if len(userTickets) == 0 {
		http.Error(w, "No tickets found for this email", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(userTickets)
}

func (h *Handlers) ViewAllAttendees(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		http.Error(w, "Date is required", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	attendees := h.es.GetAttendeesByDate(date)
	json.NewEncoder(w).Encode(attendees)
}

func (h *Handlers) CancelReservation(w http.ResponseWriter, r *http.Request) {
	var cancelRequest models.CancelRequest
	err := json.NewDecoder(r.Body).Decode(&cancelRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if h.es.CancelTicket(cancelRequest.Email, cancelRequest.EventDate) {
		json.NewEncoder(w).Encode(map[string]string{"message": "Reservation cancelled successfully"})
	} else {
		http.Error(w, "Reservation not found", http.StatusNotFound)
	}
}

func (h *Handlers) ModifySeatReservation(w http.ResponseWriter, r *http.Request) {
	var modifyRequest models.ModifyRequest
	err := json.NewDecoder(r.Body).Decode(&modifyRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newSeat := rand.Intn(100) + 1
	if h.es.ModifySeat(modifyRequest.Email, modifyRequest.EventDate, newSeat) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":    "Seat modified successfully",
			"new_seat":   newSeat,
			"event_date": modifyRequest.EventDate,
		})
	} else {
		http.Error(w, "Reservation not found", http.StatusNotFound)
	}
}
