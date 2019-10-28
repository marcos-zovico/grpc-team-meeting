package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/msouza/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("CAlculator client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cound not connect: %v", err)
	}

	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)

	// doUnary(c)
	// doServerStreaming(c)
	doClientStreaming(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
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

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a PrimeDecomposition Server Streaming RPC...")

	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 12390392840,
	}

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeDecomposition RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}

}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Printf("Starting to do a Client Streaming RPC...\n")

	stream, err := c.ComputeAverage(context.Background())

	if err != nil {
		log.Fatalf("error while opening stream ")
	}

	numbers := []int32{3, 5, 9, 54, 23}

	// we iterate over our slice and send each message individualy
	for _, number := range numbers {
		fmt.Printf("Sending number: %v\n", number)
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while reciving response: %v", err)
	}

	fmt.Printf("The avarege is: %v\n", res.GetAverage())
}
