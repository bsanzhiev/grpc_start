package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server
	conn, errConn := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if errConn != nil {
		log.Fatalf("did not connect: %v", errConn)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, errR := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if errR != nil {
		log.Fatalf("could not greet: %v", errR)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	//For Say hello again
	r, errAgain := c.SayHelloAgain(ctx, &pb.HelloRequest{Name: *name})
	if errAgain != nil {
		log.Fatalf("could not greet: %v", errAgain)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
