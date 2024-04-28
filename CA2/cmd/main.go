package main

import (
	"bufio"
	"fmt"
	"log"

	"os"
	"strconv"
	"strings"
	// "math/rand"
	// "time"
)

func ticketBuyer(id int, cliChannel chan string, orderChannel chan string) {
	logFileName := fmt.Sprintf("./log/logTicketBuyer-%d.txt", id)
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, fmt.Sprintf("%d >> ", id), log.LstdFlags)


	// // automatic load generator use if you want
	// go func() {
	// 	for {
	// 		sleepTime := rand.Intn(5) + 1
	// 		time.Sleep(time.Duration(sleepTime) * time.Second)
	// 		orderChannel <- fmt.Sprintf("%d buy 1 %d", id, sleepTime)
	// 		// write to logger
	// 		logger.Println("automatic ticket requested")
	// 		fmt.Println("Buyer ", id, " automatic ticket requested")
	// 	}
	// }()

	// // automatic load generator use if you want
	// go func() {
	// 	for {
	// 		sleepTime := rand.Intn(5) + 1
	// 		time.Sleep(time.Duration(sleepTime) * time.Second)
	// 		orderChannel <- fmt.Sprintf("%d events info", id)
	// 		// write to logger
	// 		logger.Println("automatic event info requested")
	// 		fmt.Println("Buyer ", id, " automatic event info requested")
	// 	}
	// }()

	for order := range cliChannel {
		orderChannel <- order
		// write to logger
		logger.Println("ticket requested from input")
		fmt.Println("Buyer ", id, " ticket requested from input")
	}

}

func loadBalancer(orderChannel chan string) {

	// logger
	logFileName := "./log/loadBalancer.txt"
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, "loadBalancer >> ", log.LstdFlags)

	eventInfoChannel := make(chan EventInfoRequest, 10)
	eventTicketChannel := make(chan TicketRequest, 10)

	ts := TicketServices{}

	go eventInfoHandler(&ts, eventInfoChannel)
	go eventTicketHandler(&ts, eventTicketChannel)

	previousBuyerID := -1
	resend := false

	for order := range orderChannel {
		// fmt.Println("Ticket sold: ", order)
		// logger.Println("Ticket sold: ", order)
		buyerID, _ := strconv.Atoi(strings.Split(order, " ")[0])

		if buyerID == previousBuyerID && !resend {
			// resend it to chanel
			logger.Println("resend ", buyerID, " to channel")
			resend = true
			orderChannel <- order
			continue
		}

		resend = false
		previousBuyerID = buyerID

		// send it to event list shower or event ticket buyer
		command := strings.Split(order, " ")[1]

		if command == "buy" {
			ticketBuyerID, _ := strconv.Atoi(strings.Split(order, " ")[0])
			eventID, _ := strconv.Atoi(strings.Split(order, " ")[2])
			numberOfTickets, _ := strconv.Atoi(strings.Split(order, " ")[3])
			logger.Println("ticket requested by ", ticketBuyerID)
			eventTicketChannel <- TicketRequest{NumberOfTickets: numberOfTickets, EventId: eventID, UserId: ticketBuyerID}
		} else if command == "events" {
			logger.Println("event info requested by ", buyerID)
			eventInfoChannel <- EventInfoRequest{UserId: buyerID}

		} else {
			printManual()
		}
	}
}

func eventInfoHandler(ts *TicketServices, inputChannel chan EventInfoRequest) {
	file, err := os.OpenFile("./log/eventInfoHandler.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, "eventInfoHandler >> ", log.LstdFlags)

	for infoReq := range inputChannel {
		ts.showEvents(infoReq)
		logger.Printf("event info requested by %d", infoReq.UserId)
	}
}

func eventTicketHandler(ts *TicketServices, inputChannel chan TicketRequest) {
	file, err := os.OpenFile("./log/eventTicketHandler.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, "eventTicketHandler >> ", log.LstdFlags)

	for ticketReq := range inputChannel {
		ts.buyTicket(ticketReq)
		logger.Printf("ticket requested by %d", ticketReq.UserId)
	}
}

func printManual() {
	command1 := "<buyerID> buy <eventID> <numberOfTickets>"
	command2 := "<buyerID> events info"
	fmt.Printf("Help : \n %s \n %s \n", command1, command2)
}

func main() {

	ticketBuyersNumber := 4

	cliChannels := make([]chan string, ticketBuyersNumber)
	for i := 0; i < ticketBuyersNumber; i++ {
		cliChannels[i] = make(chan string)
	}
	orderChannel := make(chan string, 10)

	for i := 0; i < ticketBuyersNumber; i++ {
		go ticketBuyer(i, cliChannels[i], orderChannel)
	}
	go loadBalancer(orderChannel)

	reader := bufio.NewReader(os.Stdin)

	// logger
	file, err := os.OpenFile("./log/cli.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, "cli >> ", log.LstdFlags)

	for {
		fmt.Print(" >>  ")
		text, _ := reader.ReadString('\n')
		// buyerID := strings.Split(text, " ")
		buyerID, err := strconv.Atoi(strings.Split(text, " ")[0])
		if err != nil {
			printManual()
			continue
		}
		cliChannels[buyerID] <- text
		logger.Println(text)
	}
}
