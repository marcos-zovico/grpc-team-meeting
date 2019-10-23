package main

import (
	"context"
	"fmt"
	"log"

	"github.com/msouza/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cound not connect: %v", err)
	}

	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)

	doSun(c)
	doMultiply(c)
}

func doSun(c calculatorpb.CalculatorServiceClient) {
	fmt.Printf("Starting sun to do a Unary RPC...")
	req := &calculatorpb.CalcRequest{
		Operation: &calculatorpb.Operation{
			NumberOne: 5.3,
			NumberTwo: 7,
		},
	}

	res, err := c.Sun(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doMultiply(c calculatorpb.CalculatorServiceClient) {
	fmt.Printf("Starting multiply to do a Unary RPC...")
	req := &calculatorpb.CalcRequest{
		Operation: &calculatorpb.Operation{
			NumberOne: 5.3,
			NumberTwo: 7,
		},
	}

	res, err := c.Multiply(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
