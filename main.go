package main

import (
	"flag"
	"github.com/JobNing/corehub/config"
	"github.com/JobNing/corehub/grpc"
	"github.com/JobNing/user-rpc/api"
	_ "github.com/JobNing/user-rpc/config"
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
