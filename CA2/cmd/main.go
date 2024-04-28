package main

import (
	"bufio"
	"fmt"
	"log"

	// "math/rand"
	"os"
	"strconv"
	"strings"
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

	go func() {
		for order := range cliChannel {
			orderChannel <- fmt.Sprintf("Buyer %d - order : %s", id, order)
			// write to logger
			logger.Println("ticket requested from input")
			fmt.Println("Buyer ", id, " ticket requested from input")
		}
	}()

	for {
	}

	// for {
	// 	sleepTime := rand.Intn(5) + 1
	// 	time.Sleep(time.Duration(sleepTime) * time.Second)
	// 	orderChannel <- fmt.Sprintf("Buyer %d - order : %s", id, "order")
	// 	// write to logger
	// 	logger.Println("ticket requested")
	// 	fmt.Println("Buyer ", id, " ticket requested")
	// }

}

func loadBalancer(orderChannel chan string) {
	// write from beging and create from begining if exists too
	logFileName := fmt.Sprintf("./log/loadBalancer.txt")
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, fmt.Sprintf("loadBalancer >> "), log.LstdFlags)

	for order := range orderChannel {
		fmt.Println("Ticket sold: ", order)
		logger.Println("Ticket sold: ", order)
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
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println("You entered:", text)
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
