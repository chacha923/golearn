package grpc

import (
	"net"
	pb "golearn/communication/grpc/pb"
	"context"
	"log"
	"google.golang.org/grpc"
)

const(
	port = ":50051"
)

type server struct{}


func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func RunGrpcServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}