package main

import (
	"time"
)

type Ticket struct {
	Id      string
	EventId string
}

type Event struct {
	Id               string
	Name             string
	Date             time.Time
	TotalTickets    int
	AvailableTickets int
}

type TicketRequest struct {
	NumberOfTickets int
	EventId int
	UserId  int
}

type EventInfoRequest struct {
	UserId int
}
