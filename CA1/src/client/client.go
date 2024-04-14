package main

import (
	"dist-projects/ca1/src/orderingsystem"
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateOrderRequest(orders []string) *orderingsystem.OrdersRequest {
	return &orderingsystem.OrdersRequest{
		OrdersIds: orders,
	}
}

func GetGetOrderRequest() *orderingsystem.OrdersRequest {
	return CreateOrderRequest([]string{
		"apple",
		"cher",
	},
	)
}

func GetOrderCall(client orderingsystem.OrderManagementServiceClient) {
	log.Println("Try calling method `GetOrder`...")

	req := GetGetOrderRequest()
	response, err := client.GetOrder(context.Background(), req)

	if err != nil {
		log.Fatalf("Error calling `GetOrder`: %v", err)
	}

	log.Printf("Response from server for method `GetOrder`: %v\n", response)
}

func SearchOrdersRequest() *orderingsystem.OrdersRequest {
	return GetGetOrderRequest()
}

func SearchOrdersCall(client orderingsystem.OrderManagementServiceClient) {
	log.Println("Try calling method `SearchOrders`...")

	req := SearchOrdersRequest()
	stream, err := client.SearchOrders(context.Background(), req)
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
}

func GetUpdateOrdersRequest() []*orderingsystem.OrdersRequest {
	return []*orderingsystem.OrdersRequest{
		CreateOrderRequest([]string{
			"yellow apple",
			"green apple",
		},
		),
		CreateOrderRequest([]string{
			"strawberry",
		},
		),
		CreateOrderRequest([]string{
			"lemon",
			"pineapple",
		},
		),
	}
}

func UpdateOrdersCall(client orderingsystem.OrderManagementServiceClient) {
	log.Println("Try calling method `UpdateOrders`...")

	stream, err := client.UpdateOrders(context.Background())
	if err != nil {
		log.Fatalf("Error calling `UpdateOrders`: %v", err)
	}

	updateReqs := GetUpdateOrdersRequest()

	for _, updateReq := range updateReqs {
		if err := stream.Send(updateReq); err != nil {
			log.Fatalf("Error happend when sending orderReq for method `UpdateOrders`: %v", err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error happend when closing stream for method `UpdateOrders`: %v", err)
	}
	log.Printf("Response from server for method `UpdateOrders`: %v\n", reply)
}

func GetProcessOrdersRequest() []*orderingsystem.OrdersRequest {
	return []*orderingsystem.OrdersRequest{
		GetGetOrderRequest(),
		CreateOrderRequest([]string{
			"orange",
			"nana",
		},
		),
	}
}

func ProcessOrdersCall(client orderingsystem.OrderManagementServiceClient) {
	log.Println("Try calling method `ProcessOrders`...")

	processReqs := GetProcessOrdersRequest()
	stream, err := client.ProcessOrders(context.Background())
	if err != nil {
		log.Fatalf("Error calling `ProcessOrders`: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		count := 1
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Error getting server response for method `ProcessOrders`: %v", err)
			}
			log.Printf("Response from server for method `ProcessOrders`: number=%v, value= %v\n", count, res)
			count++
		}
	}()
	for _, processReq := range processReqs {
		if err := stream.Send(processReq); err != nil {
			log.Fatalf("Error happend when sending orderReq for method `ProcessOrders`: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect to 9000: %v", err)
	}

	log.Print("Successfully connected to server!")

	defer conn.Close()

	client := orderingsystem.NewOrderManagementServiceClient(conn)

	GetOrderCall(client)
	SearchOrdersCall(client)
	UpdateOrdersCall(client)
	ProcessOrdersCall(client)
}
