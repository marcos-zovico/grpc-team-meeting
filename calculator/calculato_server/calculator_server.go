package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/msouza/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sun(ctx context.Context, req *calculatorpb.CalcRequest) (*calculatorpb.CalcResponse, error) {
	fmt.Printf("Sun function was invoked with %v\n", req)
	result := req.GetOperation().GetNumberOne() + req.GetOperation().GetNumberTwo()

	response := &calculatorpb.CalcResponse{
		Result: result,
	}
	return response, nil
}

func (*server) Multiply(ctx context.Context, req *calculatorpb.CalcRequest) (*calculatorpb.CalcResponse, error) {
	fmt.Printf("Multiply function was invoked with %v\n", req)

	result := req.GetOperation().GetNumberOne() * req.GetOperation().GetNumberTwo()

	response := &calculatorpb.CalcResponse{
		Result: result,
	}
	return response, nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Faliled to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
