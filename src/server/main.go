package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-test/greeter"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("受信: %s", req.Name)
	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	fmt.Println("サーバー起動: port 50051")
	s.Serve(lis)
}
