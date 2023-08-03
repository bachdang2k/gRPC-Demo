package main

import (
	"GRPC_Course/calculator/calculatorpb"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	calculator.UnimplementedCalculatorServiceServer
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {

	fmt.Printf("port %d", *port)
	fmt.Println()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	calculator.RegisterCalculatorServiceServer(s, &server{})

	fmt.Println("calculator is running ... ")

	//s.Serve(lis): run server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("err while create listen %v", err)
	}

}
