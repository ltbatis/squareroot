package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ltbatista/squareroot/squareroot/squarepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hello I'm client...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := squarepb.NewSquareRootServiceClient(cc)
	doErrorUnary(c)
}

func doErrorUnary(c squarepb.SquareRootServiceClient) {
	fmt.Println("Starting to do a SquareRoot unary...")

	// correct call
	doErrorCall(c, 10)

	// error call
	doErrorCall(c, -2)
}

func doErrorCall(c squarepb.SquareRootServiceClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &squarepb.SquareRootRequest{
		Number: n,
	})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			fmt.Printf("Error message from server: %v\n", respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("Big error calling SquareRoot: %v", err)
			return
		}
	}
	fmt.Printf("Result of square root of %v: %.4f\n\n", n, res.GetNumberRoot())
	time.Sleep(1000 * time.Millisecond)
}
