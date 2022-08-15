package main

import (
	"context"
	"fmt"
	grpc "github.com/asim/go-micro/plugins/client/grpc/v4"
	consul "github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	pb "mall-client/proto/product"
	all_router "mall-client/router"
	"net/http"
)

var (
	service = "mall_product"
	version = "latest"
)

func testClient(ctx *gin.Context) {
	consulReg := consul.NewRegistry()
	// Create service

	srv := micro.NewService(
		micro.Client(grpc.NewClient()),
		micro.Registry(consulReg),
	)

	srv.Init()

	// Create client
	c := pb.NewMallproductService(service, srv.Client())

	// Call service
	rsp, err := c.Call(context.Background(), &pb.CallRequest{Name: "John"})
	if err != nil {
		log.Fatal(err)
	}

	log.Info(rsp)

	ctx.JSONP(http.StatusOK, rsp)

}

func main() {

	router := gin.Default()
	all_router.InitRouter(router)
	err := router.Run("127.0.0.1:8080")
	if err != nil {
		fmt.Println("run err", err)
		return
	}
}
