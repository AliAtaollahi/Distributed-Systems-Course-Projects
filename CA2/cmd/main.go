package main

import (
	"bufio"
	"fmt"
	"log"
<<<<<<< HEAD
	"math/rand"
	"os"
	"time"
=======

	"os"
	"strconv"
	"strings"

	// "math/rand"
	// "time"
>>>>>>> CA2
)

func ticketBuyer(id int, cliChannel chan string, orderChannel chan string) {
	logFileName := fmt.Sprintf("./log/logTicketBuyer-%d.txt", id)
<<<<<<< HEAD
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
=======
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
>>>>>>> CA2
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, fmt.Sprintf("%d >> ", id), log.LstdFlags)

<<<<<<< HEAD
	go func() {
		for order := range cliChannel {
			orderChannel <- fmt.Sprintf("Buyer %d - order : %s", id, order)
			// write to logger
			logger.Println("ticket requested from input")
			fmt.Println("Buyer ", id, " ticket requested from input")
		}
	}()

	for {
		sleepTime := rand.Intn(5) + 1
		time.Sleep(time.Duration(sleepTime) * time.Second)
		orderChannel <- fmt.Sprintf("Buyer %d - order : %s", id, "order")
		// write to logger
		logger.Println("ticket requested")
		fmt.Println("Buyer ", id, " ticket requested")
	}

}

func ticketSeller(orderChannel chan string) {
	for order := range orderChannel {
		fmt.Println("Ticket sold: ", order)
	}
=======
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

	for input := range cliChannel {
		_, err := strconv.Atoi(strings.Split(input, " ")[0])
		if err != nil { // output result
			fmt.Printf("Buyer %d output result: %s ", id, input)
		} else {
			orderChannel <- input
			logger.Println("request sent to load balancer" + input)
		}
	}
}

func loadBalancer(orderChannel chan string, eventInfoChannel chan EventInfoRequest, eventTicketChannel chan TicketRequest) {

	// logger
	logFileName := "./log/loadBalancer.txt"
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, "loadBalancer >> ", log.LstdFlags)

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

func eventInfoHandler(ts *TicketServices, inputChannel chan EventInfoRequest, outputChannels []chan string) {
	file, err := os.OpenFile("./log/eventInfoHandler.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, "eventInfoHandler >> ", log.LstdFlags)

	for infoReq := range inputChannel {
		message := ts.showEvents(infoReq)
		outputChannels[infoReq.UserId] <- message
		logger.Printf(message)
	}
}

func eventTicketHandler(ts *TicketServices, inputChannel chan TicketRequest, outputChannels []chan string) {
	file, err := os.OpenFile("./log/eventTicketHandler.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, "eventTicketHandler >> ", log.LstdFlags)

	for ticketReq := range inputChannel {
		message := ts.buyTicket(ticketReq)
		outputChannels[ticketReq.UserId] <- message
		logger.Printf(message)
	}
}

func printManual() {
	command1 := "<buyerID> buy <eventID> <numberOfTickets>"
	command2 := "<buyerID> events info"
	fmt.Printf("Help : \n %s \n %s \n", command1, command2)
>>>>>>> CA2
}

func main() {

	ticketBuyersNumber := 4

<<<<<<< HEAD
=======
	// all channels
>>>>>>> CA2
	cliChannels := make([]chan string, ticketBuyersNumber)
	for i := 0; i < ticketBuyersNumber; i++ {
		cliChannels[i] = make(chan string)
	}
<<<<<<< HEAD
	orderChannel := make(chan string, 10)
=======
	orderChannel := make(chan string)

	eventInfoChannel := make(chan EventInfoRequest)
	eventTicketChannel := make(chan TicketRequest)
>>>>>>> CA2

	for i := 0; i < ticketBuyersNumber; i++ {
		go ticketBuyer(i, cliChannels[i], orderChannel)
	}
<<<<<<< HEAD
	go ticketSeller(orderChannel)

	reader := bufio.NewReader(os.Stdin)

	// logger
	file, err := os.OpenFile("./log/cli.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
=======
	go loadBalancer(orderChannel, eventInfoChannel, eventTicketChannel)

	ts := TicketServices{}
	go eventInfoHandler(&ts, eventInfoChannel, cliChannels)
	go eventTicketHandler(&ts, eventTicketChannel, cliChannels)

	// logger
	file, err := os.OpenFile("./log/cli.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
>>>>>>> CA2
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, "cli >> ", log.LstdFlags)

<<<<<<< HEAD
	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println("You entered:", text)
		cliChannels[1] <- text
=======
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		buyerID, err := strconv.Atoi(strings.Split(text, " ")[0])
		if err != nil {
			printManual()
			continue
		}
		cliChannels[buyerID] <- text
>>>>>>> CA2
		logger.Println(text)
	}
}
