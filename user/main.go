package main

import (
	"github.com/Qianjiachen55/micro-todoList/conf"
	"github.com/Qianjiachen55/micro-todoList/core"
	"github.com/Qianjiachen55/micro-todoList/services"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	conf.Init()
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
		)

	microServer := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(etcdReg),
		)
	microServer.Init()

	_=services.RegisterUserServiceHandler(microServer.Server(),new(core.UserService))
	_=microServer.Run()

}
