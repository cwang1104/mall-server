package main

import (
	"mall_product/handler"
	pb "mall_product/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"

	grpcc "github.com/asim/go-micro/plugins/client/grpc/v4"
	consul "github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/server/grpc/v4"
)

var (
	service = "mall_product"
	version = "latest"
)

func main() {
	consulReg := consul.NewRegistry()
	// Create service
	srv := micro.NewService(
		micro.Server(grpc.NewServer()),
		micro.Client(grpcc.NewClient()),
		micro.Registry(consulReg),
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterMallproductHandler(srv.Server(), new(handler.Mallproduct))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
