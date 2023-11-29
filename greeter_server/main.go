package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "helloworld"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The Server port")
)

// server is used to implement helloworld.GreeterServer
type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Recieved: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, errLis := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if errLis != nil {
		log.Fatalf("failed to listen: %v", errLis)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if errLis := s.Serve(lis); errLis != nil {
		log.Fatalf("failed to serve: %v", errLis)
	}
}
