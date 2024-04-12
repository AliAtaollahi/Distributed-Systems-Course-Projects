package main

import (
	"dist-projects/ca1/src/orderingsystem"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect to 9000: %v", err)
	}

	defer conn.Close()

	c := orderingsystem.NewOrderManagementServiceClient(conn)
	req := orderingsystem.OrdersRequest{
		OrdersIds: []string{
            "apple",
			"cher",
        },
	}

	response, err := c.GetOrder(context.Background(), &req)

	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	log.Printf("Response from server: %v", response)

}