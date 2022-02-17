package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"task/conf"
	"task/core"
	"task/service"
)

func main() {
	conf.Init()
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	microServer := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address("127.0.0.1:8083"),
		micro.Registry(etcdReg),
	)
	microServer.Init()

	service.RegisterTaskServiceHandler(microServer.Server(),new(core.TaskService))
	microServer.Run()
}
