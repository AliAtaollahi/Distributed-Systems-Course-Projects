package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func ticketBuyer(id int, cliChannel chan string, orderChannel chan string) {
	logFileName := fmt.Sprintf("./log/logTicketBuyer-%d.txt", id)
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, fmt.Sprintf("%d >> ", id), log.LstdFlags)

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
	go ticketSeller(orderChannel)

	reader := bufio.NewReader(os.Stdin)

	// logger
	file, err := os.OpenFile("./log/cli.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, "cli >> ", log.LstdFlags)

	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println("You entered:", text)
		cliChannels[1] <- text
		logger.Println(text)
	}
}
