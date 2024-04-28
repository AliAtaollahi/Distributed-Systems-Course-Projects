package main

type TicketServices struct {
	cache Cache
}

func (ts *TicketServices) showEvents() {
	// lock it
	// for in cash
	// log it
	// print it on consol
}

func (ts *TicketServices) buyTicket(eventID int, numberOfTickets int) {
	// buy it
	// log it
	// print it on consol
	// make sure to lock when using cash part then unlock
}
