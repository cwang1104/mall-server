package main

import (
	pbMiaosha "mall_seckill/proto/miaosha"
	rpcMiaoSha "mall_seckill/rpc/seckill"

	grpcc "github.com/asim/go-micro/plugins/client/grpc/v4"
	consul "github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/server/grpc/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "mall_seckill"
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
	pbMiaosha.RegisterMiaoShaHandler(srv.Server(), new(rpcMiaoSha.MiaoSha))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
