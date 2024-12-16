package models

import (
	"time"
)

type Event struct {
	ID          int
	Name        string    `binding: "required`
	Description string    `binding: "required`
	Location    string    `binding: "required`
	Date        time.Time `binding: "required`
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() {
	// later added to the database
	events = append(events, e)
}

func GetAllEvents() []Event {
	// later added to the database
	return events
}
