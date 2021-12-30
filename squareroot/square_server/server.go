package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/ltbatista/squareroot/squareroot/squarepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) SquareRoot(ctx context.Context, req *squarepb.SquareRootRequest) (*squarepb.SquareRootResponse, error) {
	fmt.Printf("Squareroot function was invoked with %v\n", req)
	number := req.GetNumber()

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}
	return &squarepb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	fmt.Println("Hello, i'm the square root server...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	squarepb.RegisterSquareRootServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
