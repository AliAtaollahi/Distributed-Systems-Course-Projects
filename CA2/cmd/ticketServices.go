package main

import (
	"fmt"
)

// "fmt"
// "log"
// "os"

type TicketServices struct {
	cache Cache
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

	// logFileName := fmt.Sprintf("./log/ticketHandler.txt")
	// file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// logger := log.New(file, fmt.Sprintf("event >> "), log.LstdFlags)

	// buy it
	// log it
	// logger.Println("ticket sb bout .....")
	// make sure to lock when using cash part then unlock

	message := fmt.Sprintf("handler: ticket requested by %d \n", tr.UserId)
	return message
}
