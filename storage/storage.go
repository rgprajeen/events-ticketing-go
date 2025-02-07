package storage

import (
	"events-ticketing/models"
	"sync"
	"time"
)

type EventSystem struct {
	Tickets map[string][]models.Ticket
	mu      sync.RWMutex
}

func NewEventSystem() *EventSystem {
	return &EventSystem{
		Tickets: make(map[string][]models.Ticket),
	}
}

func (es *EventSystem) AddTicket(ticket models.Ticket) {
	es.mu.Lock()
	defer es.mu.Unlock()
	dateKey := ticket.EventDate.Format("2006-01-02")
	es.Tickets[dateKey] = append(es.Tickets[dateKey], ticket)
}

func (es *EventSystem) GetTicketsByEmail(email string) []models.Ticket {
	es.mu.RLock()
	defer es.mu.RUnlock()
	var userTickets []models.Ticket
	for _, tickets := range es.Tickets {
		for _, ticket := range tickets {
			if ticket.Email == email {
				userTickets = append(userTickets, ticket)
			}
		}
	}
	return userTickets
}

func (es *EventSystem) GetAttendeesByDate(date time.Time) []models.Ticket {
	es.mu.RLock()
	defer es.mu.RUnlock()
	dateKey := date.Format("2006-01-02")
	return es.Tickets[dateKey]
}

func (es *EventSystem) CancelTicket(email string, eventDate time.Time) bool {
	es.mu.Lock()
	defer es.mu.Unlock()
	dateKey := eventDate.Format("2006-01-02")
	tickets := es.Tickets[dateKey]
	for i, ticket := range tickets {
		if ticket.Email == email {
			es.Tickets[dateKey] = append(tickets[:i], tickets[i+1:]...)
			return true
		}
	}
	return false
}

func (es *EventSystem) ModifySeat(email string, eventDate time.Time, newSeat int) bool {
	es.mu.Lock()
	defer es.mu.Unlock()
	dateKey := eventDate.Format("2006-01-02")
	tickets := es.Tickets[dateKey]
	for i, ticket := range tickets {
		if ticket.Email == email {
			es.Tickets[dateKey][i].SeatNumber = newSeat
			return true
		}
	}
	return false
}
