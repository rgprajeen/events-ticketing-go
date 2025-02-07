package models

import "time"

type Ticket struct {
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	EventDate  time.Time `json:"event_date"`
	SeatNumber int       `json:"seat_number"`
}

type CancelRequest struct {
	Email     string    `json:"email"`
	EventDate time.Time `json:"event_date"`
}

type ModifyRequest struct {
	Email     string    `json:"email"`
	EventDate time.Time `json:"event_date"`
}
