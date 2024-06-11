package main

import (
	"context"
	orderpb "github.com/ntc-goer/microservice-examples/orderservice/proto"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	dp, err := InitializeDependency("consul")
	if err != nil {
		log.Fatalf("fail to init dependency %v", err)
	}
	ctx := context.Background()

	// Setup grpc server
	lis, err := net.Listen("tcp", ":"+dp.Config.GRPCPort)
	if err != nil {
		log.Fatalf("error listening port %v", err)
	}
	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, dp.ServiceImpl)
	// Register to discovery service
	instanceId := serviceregistration.GenerateInstanceId(dp.Config.ServiceName)
	if err := dp.ServiceDiscovery.RegisterService(instanceId, dp.Config.ServiceName, serviceregistration.GetCurrentIP(), dp.Config.GRPCPort, "http://host.docker.internal:8080/order/health"); err != nil {
		log.Fatalf("RegisterService fail: %v", err)
	}
	defer dp.ServiceDiscovery.Deregister(ctx, instanceId)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
