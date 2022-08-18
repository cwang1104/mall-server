package main

import (
	"fmt"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	_ "mall_user/conf"
	_ "mall_user/models"
	pbUser "mall_user/proto/user"
	"mall_user/service/rpcUser"

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

	err := pbUser.RegisterUserHandler(srv.Server(), new(rpcUser.User))
	if err != nil {
		fmt.Println("pbUser register new handler err", err)
		return
	}
	// Register handler
	//pb.RegisterMalluserHandler(srv.Server(), new(handler.Malluser))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
