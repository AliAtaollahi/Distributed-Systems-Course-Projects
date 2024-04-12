package main

import (
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"dist-projects/ca1/src/orderingsystem"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)


type Server struct {
	orders []string
	orderingsystem.UnimplementedOrderManagementServiceServer
}

func (s *Server) GetOrder(ctx context.Context, request *orderingsystem.OrderRequest) (*orderingsystem.OrderResponse, error){
	log.Printf("Received message body from client: %v", request)
	selectedOrders := []string{}
	for _, reqOrder := range request.OrdersIds {
		for _, order := range s.orders {
			if strings.Contains(order, reqOrder) {
				selectedOrders = append(selectedOrders, order)
			}
		}
	}
	currentTime := time.Now()
	return &orderingsystem.OrderResponse{Orders: selectedOrders, Timestamp: strconv.FormatInt(currentTime.UnixNano(), 10)}, nil
}


func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)

	}

	s := Server{
		orders: []string{
            "banana",
            "apple",
            "orange",
            "grape",
            "red apple",
            "kiwi",
            "mango",
            "pear",
            "cherry",
            "green apple",
        },
	}


	grpcServer := grpc.NewServer()

	orderingsystem.RegisterOrderManagementServiceServer(grpcServer, &s)



	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
	}
}