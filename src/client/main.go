package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "grpc-test/greeter"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	name := "World"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	res1, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SayHello: %s", res1.Message)
}
