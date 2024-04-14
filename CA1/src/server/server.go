package main

import (
	"io"
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

func (s *Server) GetOrder(ctx context.Context, request *orderingsystem.OrdersRequest) (*orderingsystem.OrdersResponse, error) {
	log.Printf("Received message body from client for method `GetOrder`: %v", request)
	selectedOrders := []string{}
	for _, reqOrder := range request.OrdersIds {
		for _, order := range s.orders {
			if strings.Contains(order, reqOrder) {
				selectedOrders = append(selectedOrders, order)
			}
		}
	}
	currentTime := time.Now()
	return &orderingsystem.OrdersResponse{Orders: selectedOrders, Timestamp: strconv.FormatInt(currentTime.UnixNano(), 10)}, nil
}

func (s *Server) SearchOrders(request *orderingsystem.OrdersRequest, stream orderingsystem.OrderManagementService_SearchOrdersServer) error {
	log.Printf("Received message body from client for method `SearchOrders`: %v", request)
	for _, reqOrder := range request.OrdersIds {
		for _, order := range s.orders {
			if strings.Contains(order, reqOrder) {
				currentTime := time.Now()
				if err := stream.Send(&orderingsystem.OrderResponse{OrderId: order, Timestamp: strconv.FormatInt(currentTime.UnixNano(), 10)}); err != nil {
					log.Fatalf("Error happend when sending order for method `SearchOrders`: %v", err)
					return err
				}
			}
		}
	}

	return nil
}

func (s *Server) UpdateOrders(stream orderingsystem.OrderManagementService_UpdateOrdersServer) error {
	receivedOrders := []string{}
	count := 0
	for {
		orderIds, err := stream.Recv()
		count++
		log.Printf("Received message body from client for method `UpdateOrders`: number=%v, value= %v\n", count, orderIds)
		if err == io.EOF {
			currentTime := time.Now()
			log.Printf("Updated Orders: %v\n", s.orders)
			return stream.SendAndClose(&orderingsystem.OrdersResponse{Orders: receivedOrders, Timestamp: strconv.FormatInt(currentTime.UnixNano(), 10)})
		}
		if err != nil {
			log.Fatalf("Error happend when receiving orderId for method `UpdateOrders`: %v", err)
			return err
		}
		for _, newOrder := range orderIds.OrdersIds {
			add := true
			for _, order := range s.orders {
				if newOrder == order {
					add = false
					break
				}
			}
			if add {
				s.orders = append(s.orders, newOrder)
			}
			receivedOrders = append(receivedOrders, newOrder)
		}
	}
}

func (s *Server) ProcessOrders(stream orderingsystem.OrderManagementService_ProcessOrdersServer) error {
	count := 0
	for {
		request, err := stream.Recv()
		log.Printf("Received message body from client for method `ProcessOrders`: number=%v, value= %v\n", count, request)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error happend when receiving orderId for method `ProcessOrders`: %v", err)
			return err
		}

		for _, reqOrder := range request.OrdersIds {
			for _, order := range s.orders {
				if strings.Contains(order, reqOrder) {
					currentTime := time.Now()
					if err := stream.Send(&orderingsystem.OrderResponse{OrderId: order, Timestamp: strconv.FormatInt(currentTime.UnixNano(), 10)}); err != nil {
						log.Fatalf("Error happend when sending order for method `ProcessOrders`: %v", err)
						return err
					}
				}
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:9000")

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
	log.Println("Server successfully started...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
	}
}
