package main

import (
	"flag"
	"github.com/JobNing/corehub/config"
	"github.com/JobNing/corehub/grpc"
	_ "github.com/JobNing/user-rpc/config"
	"github.com/JobNing/user-rpc/internal/api"
	"github.com/JobNing/user-rpc/migrate"
	grpc2 "google.golang.org/grpc"
)

var configFile = flag.String("f", "config/config.yaml", "the config file")

func main() {
	flag.Parse()
	err := config.InitViper(*configFile)
	if err != nil {
		panic(err)
	}

	err = migrate.AutoMigrate()
	if err != nil {
		panic(err)
	}

	err = grpc.RegisterGRPC(8081, func(s *grpc2.Server) {
		api.Register(s)
	})
	if err != nil {
		panic(err)
	}

}

//package main
//
//import (
//	"context"
//	"flag"
//	"fmt"
//	"log"
//	"net"
//
//	"google.golang.org/grpc"
//	pb "google.golang.org/grpc/examples/helloworld/helloworld"
//)
//
//var (
//	port = flag.Int("port", 50051, "The server port")
//)
//
//// server is used to implement helloworld.GreeterServer.
//type server struct {
//	pb.UnimplementedGreeterServer
//}
//
//// SayHello implements helloworld.GreeterServer
//func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
//	log.Printf("Received: %v", in.GetName())
//	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
//}
//
//func main() {
//	flag.Parse()
//
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	s := grpc.NewServer()
//
//	pb.RegisterGreeterServer(s, &server{})
//
//	log.Printf("server listening at %v", lis.Addr())
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}
