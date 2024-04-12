package main

import (
	"dist-projects/ca1/src/orderingsystem"
	"io"
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
		log.Fatalf("Error calling `GetOrder`: %v", err)
	}

	log.Printf("Response from server for method `GetOrder`: %v\n", response)

	stream, err := c.SearchOrders(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error calling `SearchOrders`: %v", err)
	}
	count := 1
	for {

		order, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error getting server response for method `SearchOrders`: %v", err)
		}
		log.Printf("Response from server for method `SearchOrders`: number=%v, value= %v\n", count, order)
		count++
	}

	sendStream, err := c.UpdateOrders(context.Background())
	if err != nil {
		log.Fatalf("Error calling `UpdateOrders`: %v", err)
	}

	updateReqs := []*orderingsystem.OrdersRequest{
		&orderingsystem.OrdersRequest{
			OrdersIds: []string{
				"yellow apple",
				"green apple",
			},
		},
		&orderingsystem.OrdersRequest{
			OrdersIds: []string{
				"strawberry",
			},
		},
		&orderingsystem.OrdersRequest{
			OrdersIds: []string{
				"lemon",
				"pineapple",
			},
		},
	}

	for _, updateReq := range updateReqs {
		if err := sendStream.Send(updateReq); err != nil {
			log.Fatalf("Error happend when sending orderReq for method `UpdateOrders`: %v", err)
		}
	}
	reply, err := sendStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error happend when closing stream for method `UpdateOrders`: %v", err)
	}
	log.Printf("Response from server for method `UpdateOrders`: %v\n", reply)

}
