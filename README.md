# Event Ticket Reservation System

This is a Go-based Event Ticket Reservation System that provides APIs for reserving tickets, viewing ticket details, listing attendees, canceling reservations, and modifying seat reservations.

## Prerequisites
- Go 1.16 or higher
- Git

## Installation
1. Clone the repository:
```
git clone https://github.com/rgprajeen/events-ticketing-go.git
cd events-ticketing-go
```

2. Install the required dependencies:
```
go get -u github.com/go-chi/chi/v5
go get -u github.com/stretchr/testify
```

## Running the Application
To start the server, run:
```
go run main.go
```
The server will start on `http://localhost:8080`.

## Running Tests
To run the unit tests, execute:
```
go test ./...
```

## API Endpoints
Here are the available API endpoints and how to use them with curl:
1. Reserve a Ticket
```
curl -X POST http://localhost:8080/reserve -H "Content-Type: application/json" -d '{"name":"John Doe","email":"john@example.com","event_date":"2023-05-20T00:00:00Z"}'
```

2. View Ticket Details
```
curl http://localhost:8080/view?email=john@example.com
```

3. View All Attendees for a Date
```
curl http://localhost:8080/attendees?date=2023-05-20
```

4. Cancel a Reservation
```
curl -X POST http://localhost:8080/cancel -H "Content-Type: application/json" -d '{"email":"john@example.com","event_date":"2023-05-20T00:00:00Z"}'
```

5. Modify a Seat Reservation
```
curl -X POST http://localhost:8080/modify -H "Content-Type: application/json" -d '{"email":"john@example.com","event_date":"2023-05-20T00:00:00Z"}'
```

## Development
### Adding New Features
1. If you're adding a new model, add it to `models/models.go`.
2. For new storage operations, update `storage/storage.go`.
3. Implement new handlers in `handlers/handlers.go`.
4. Add the new route in `main.go`.
5. Write tests for the new feature in `handlers/handlers_test.go`.

### Running Tests
To run all tests:
```
go test ./...
```

To run tests with coverage:
```
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Notes

- This implementation uses in-memory storage. For a production environment, consider replacing it with a database.
- The system currently generates random seat numbers. In a real-world scenario, you might want to implement a more sophisticated seat allocation system.
- There's no authentication or authorization implemented. In a production system, you would want to add these security features.
