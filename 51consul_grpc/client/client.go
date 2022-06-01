package main

import (
	"consul_grpc/pb"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

func main() {
	// 服务发现
	// 1.初始化 consul 配置
	consulConfig := api.DefaultConfig()
	// 2.创建 consul 对象
	consulClient, _ := api.NewClient(consulConfig)
	// 3.服务发现. 从 consul 上, 获取健康的服务
	//  参数：
	//	service: 服务名。 -- 注册服务时，指定该string
	//	tag：外名/别名。 如果有多个， 任选一个
	//	passingOnly：是否通过健康检查。 true
	//	q：查询参数。 通常传 nil
	//	返回值：
	//	ServiceEntry： 存储服务的切片。
	//	QueryMeta：额外查询返回值。 nil
	//	error： 错误信息
	services, _, _ := consulClient.Health().Service("userService", "grpc", true, nil)
	// 4.简单的负载均衡，获取服务地址
	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)
	//////////////////////以下为 grpc 服务远程调用///////////////////////////
	grpcConn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	grpcClient := pb.NewUserServiceClient(grpcConn)
	req := pb.GetUserInfoReq{UserId: 1}
	rsp, _ := grpcClient.GetUserInfo(context.Background(), &req)
	fmt.Println(rsp)
}
