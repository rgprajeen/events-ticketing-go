package handlers

import (
	"bytes"
	"encoding/json"
	"events-ticketing/models"
	"events-ticketing/storage"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() (*chi.Mux, *Handlers) {
	es := storage.NewEventSystem()
	h := NewHandlers(es)

	r := chi.NewRouter()
	r.Post("/reserve", h.ReserveTicket)
	r.Get("/view", h.ViewTicketDetails)
	r.Get("/attendees", h.ViewAllAttendees)
	r.Post("/cancel", h.CancelReservation)
	r.Post("/modify", h.ModifySeatReservation)

	return r, h
}

func TestReserveTicket(t *testing.T) {
	r, _ := setupTestRouter()

	ticket := models.Ticket{
		Name:      "John Doe",
		Email:     "john@example.com",
		EventDate: time.Now(),
	}
	body, _ := json.Marshal(ticket)

	req, _ := http.NewRequest("POST", "/reserve", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, "Ticket reserved successfully", response["message"])
}

func TestViewTicketDetails(t *testing.T) {
	r, h := setupTestRouter()

	// Add a ticket
	h.ReserveTicket(httptest.NewRecorder(), httptest.NewRequest("POST", "/reserve", bytes.NewBufferString(`{"name":"John Doe","email":"john@example.com","event_date":"2023-05-20T00:00:00Z"}`)))

	req, _ := http.NewRequest("GET", "/view?email=john@example.com", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var tickets []models.Ticket
	json.Unmarshal(rr.Body.Bytes(), &tickets)
	assert.Equal(t, 1, len(tickets))
	assert.Equal(t, "John Doe", tickets[0].Name)
}

func TestViewAllAttendees(t *testing.T) {
	r, h := setupTestRouter()

	// Add a ticket
	h.ReserveTicket(httptest.NewRecorder(), httptest.NewRequest("POST", "/reserve", bytes.NewBufferString(`{"name":"John Doe","email":"john@example.com","event_date":"2023-05-20T00:00:00Z"}`)))

	req, _ := http.NewRequest("GET", "/attendees?date=2023-05-20", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var attendees []models.Ticket
	json.Unmarshal(rr.Body.Bytes(), &attendees)
	assert.Equal(t, 1, len(attendees))
	assert.Equal(t, "John Doe", attendees[0].Name)
}

func TestCancelReservation(t *testing.T) {
	r, h := setupTestRouter()

	// Add a ticket
	h.ReserveTicket(httptest.NewRecorder(), httptest.NewRequest("POST", "/reserve", bytes.NewBufferString(`{"name":"John Doe","email":"john@example.com","event_date":"2023-05-20T00:00:00Z"}`)))

	cancelRequest := models.CancelRequest{
		Email:     "john@example.com",
		EventDate: time.Date(2023, 5, 20, 0, 0, 0, 0, time.UTC),
	}
	body, _ := json.Marshal(cancelRequest)

	req, _ := http.NewRequest("POST", "/cancel", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, "Reservation cancelled successfully", response["message"])
}

func TestModifySeatReservation(t *testing.T) {
	r, h := setupTestRouter()

	// Add a ticket
	h.ReserveTicket(httptest.NewRecorder(), httptest.NewRequest("POST", "/reserve", bytes.NewBufferString(`{"name":"John Doe","email":"john@example.com","event_date":"2023-05-20T00:00:00Z"}`)))

	modifyRequest := models.ModifyRequest{
		Email:     "john@example.com",
		EventDate: time.Date(2023, 5, 20, 0, 0, 0, 0, time.UTC),
	}
	body, _ := json.Marshal(modifyRequest)

	req, _ := http.NewRequest("POST", "/modify", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, "Seat modified successfully", response["message"])
	assert.NotNil(t, response["new_seat"])
}
