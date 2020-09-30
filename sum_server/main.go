package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"projects/grpcserverstreamingapi/sumpb/sumpb"
	"time"
)

const port = ":50051"

type server struct{
	sumpb.UnimplementedSumServiceServer
}

func (s *server) Sum(in *sumpb.SumRequest, stream sumpb.SumService_SumServer) error {
	log.Printf("Received %v and %v \n", in.GetFirstNumber(), in.GetSecondNumber())
	for i := 0; i < 3; i++ {
		sum := in.GetFirstNumber() + in.GetSecondNumber()
		stream.Send(&sumpb.SumResponse{Sum: sum})
		log.Println("Response sent")
		time.Sleep(time.Second)
	}
	return nil
}

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen to port: %v", err)
	}
	log.Println("Listening to the port")
	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}