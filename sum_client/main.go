package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"projects/grpcserverstreamingapi/sumpb/sumpb"
)

const port = "localhost:50051"

func main(){
	conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to dial port: %v", err)
	}
	c := sumpb.NewSumServiceClient(conn)
	req := sumpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 3,
	}
	stream, err := c.Sum(context.Background(), &req)
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}
	for{
		streamresponse, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Printf("The Sum Is: %v", streamresponse.GetSum())
	}
}