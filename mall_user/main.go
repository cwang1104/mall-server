package main

import (
	"mall_user/handler"
	pb "mall_user/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"

	grpcc "github.com/asim/go-micro/plugins/client/grpc/v4"
	consul "github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/server/grpc/v4"
)

var (
	service = "mall_user"
	version = "latest"
)

func main() {
	consulReg := consul.NewRegistry()

	// Create service
	srv := micro.NewService(
		micro.Server(grpc.NewServer()),
		micro.Client(grpcc.NewClient()),
		micro.Name(service),
		micro.Registry(consulReg),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterMalluserHandler(srv.Server(), new(handler.Malluser))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
