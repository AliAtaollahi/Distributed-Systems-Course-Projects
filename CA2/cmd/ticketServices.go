package main

import (
	"fmt"
)

// "fmt"
// "log"
// "os"

type TicketServices struct {
	EventCache   Cache
	lock         sync.Mutex
	ticketLogger *log.Logger
}

func (ts *TicketServices) InitializeCash() {
	ts.EventCache.Initialize()

	logFileName := "./log/ticketLogger.txt"
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// defer file.Close()
	ts.ticketLogger = log.New(file, "ticketLogger >> ", log.LstdFlags)
}

func (ts *TicketServices) showEvents(req EventInfoRequest) string {

	// lock it
	// sync.Map
	// for in cash
	// log it
	message := fmt.Sprintf("handler: event info requested by %d \n", req.UserId)
	return message
}

func (ts *TicketServices) buyTicket(tr TicketRequest) string {

	// print number of tickets
	fmt.Printf("number of tickets %d \n", tr.NumberOfTickets)

	ts.lock.Lock()
	message := ""
	// get event from cache
	event, ok := ts.EventCache.Get(fmt.Sprintf("%d", tr.EventId))
	if !ok {
		message = fmt.Sprintf("\n event %d not found \n", tr.EventId)
	} else {
		// check if tickets are available
		if event.(Event).AvailableTickets >= tr.NumberOfTickets {
			// craete a new event with updated available tickets
			// updatesEvent := event
			updatesEvent := Event{
				Id:               event.(Event).Id,
				Name:             event.(Event).Name,
				Date:             event.(Event).Date,
				TotalTickets:     event.(Event).TotalTickets,
				AvailableTickets: event.(Event).AvailableTickets - tr.NumberOfTickets,
			}
			ts.EventCache.Set(updatesEvent.Id, updatesEvent)
			ts.ticketLogger.Printf("%d ticket bout for event %d buy user %d \n", tr.NumberOfTickets, tr.EventId, tr.UserId)
			message = fmt.Sprintf("\n %d ticket bout for event %d buy user %d \n", tr.NumberOfTickets, tr.EventId, tr.UserId)
		} else {
			message = fmt.Sprintf("\n not enough tickets for event %d \n", tr.EventId)
		}
	}
	ts.lock.Unlock()
	return message
}
