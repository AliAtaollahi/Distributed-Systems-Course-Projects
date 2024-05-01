package main

import "fmt"

type Ticket struct {
	Id      string
	EventId string
}

// function for ticket that returns its string representation
func (t Ticket) String() string {
	return t.Id + " " + t.EventId
}

type Event struct {
	Id               string
	Name             string
	Date             string
	TotalTickets     int
	AvailableTickets int
}

// function for event that returns its string representation
func (e Event) String() string {
	return fmt.Sprintf("\nId: %s, Name: %s, Date: %s, TotalTickets: %d, AvailableTickets: %d\n", e.Id, e.Name, e.Date, e.TotalTickets, e.AvailableTickets)
}

type TicketRequest struct {
	NumberOfTickets int
	EventId         int
	UserId          int
}

type EventInfoRequest struct {
	UserId int
}
