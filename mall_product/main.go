package main

import (
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	pbProduct "mall_product/proto/product"
	pbSeckill "mall_product/proto/seckill"
	rpcProduct "mall_product/rpc/product"
	rpcSeckill "mall_product/rpc/seckill"

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
	//srv.Init()

	pbProduct.RegisterProductsHandler(srv.Server(), new(rpcProduct.Product))
	pbSeckill.RegisterSecKillsHandler(srv.Server(), new(rpcSeckill.Seckill))
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
