package helloService

import (
	"context"
	pb "github.com/The0ne10/myTinderProto/hello_service/proto"
)

type Server struct {
	pb.UnimplementedHelloServiceServer
}

func New() *Server {
	return &Server{}
}

func (srv *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello, " + req.Name + "!",
	}, nil
}
