package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	mservice "github.com/micro/go-micro/service"

	//"github.com/micro/go-micro/service"
	"github.com/micro/go-plugins/registry/consul"
	"go-micro-grpc/ServiceImpl"
	"go-micro-grpc/Services"
)

func main() {
	csReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	prdService := micro.NewService(
		micro.Name("prodService"),
		micro.Address(":8011"),
		micro.Registry(csReg),
	)

	prdService.Init()

	Service.RegisterProdServiceHandler(prdService.Server(),new(ServiceImpl.ProdService))

	prdService.Run()
}
