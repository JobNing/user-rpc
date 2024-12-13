package main

import (
	"github.com/JobNing/corehub/grpc"
	grpc2 "google.golang.org/grpc"
	"user-rpc/api"
	"user-rpc/migrate"

	_ "user-rpc/config"
)

func main() {
	err := migrate.AutoMigrate()
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
