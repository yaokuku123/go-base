package main

import (
	"consul_grpc/pb"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
)

type UserService struct {
}

func (*UserService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (*pb.GetUserInfoRsp, error) {
	if req.UserId == 10 {
		return &pb.GetUserInfoRsp{
			UserId:   req.UserId,
			UserName: "yorick_10",
			Age:      25,
		}, nil
	}
	return &pb.GetUserInfoRsp{
		UserId:   req.UserId,
		UserName: "yorick_other",
		Age:      20,
	}, nil
}

func main() {
	// 服务注册
	// 1.初始化 consul 配置
	consulConfig := api.DefaultConfig()
	// 2.创建 consul 对象
	consulClient, _ := api.NewClient(consulConfig)
	// 3.注册的服务配置
	reg := api.AgentServiceRegistration{
		ID:      "userService id",
		Name:    "userService",
		Tags:    []string{"consul", "grpc"},
		Address: "127.0.0.1",
		Port:    8080,
		Check: &api.AgentServiceCheck{
			CheckID:  "consul grpc test",
			TCP:      "127.0.0.1:8080",
			Timeout:  "1s",
			Interval: "5s",
		},
	}
	// 4. 注册 grpc 服务到 consul 上
	consulClient.Agent().ServiceRegister(&reg)

	//////////////////////以下为 grpc 服务远程调用////////////////////////
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, new(UserService))
	listener, _ := net.Listen("tcp", "127.0.0.1:8080")
	defer listener.Close()
	fmt.Println("server start...")
	grpcServer.Serve(listener)
}
