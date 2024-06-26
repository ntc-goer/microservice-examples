// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/ntc-goer/microservice-examples/consumerservice/config"
	"github.com/ntc-goer/microservice-examples/consumerservice/service"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/consul"
)

// Injectors from wire.go:

//go:generate wire
func InitializeDependency(dcType string) (*CoreDependency, error) {
	configConfig, err := config.Load()
	if err != nil {
		return nil, err
	}
	healthService, err := service.NewHealthService()
	if err != nil {
		return nil, err
	}
	userService, err := service.NewUserService(configConfig)
	if err != nil {
		return nil, err
	}
	coreService := service.NewCoreService(healthService, userService)
	registry, err := consul.NewRegistry()
	if err != nil {
		return nil, err
	}
	coreDependency := NewCoreDependency(configConfig, coreService, registry)
	return coreDependency, nil
}

// wire.go:

type CoreDependency struct {
	Config           *config.Config
	Service          *service.CoreService
	ServiceDiscovery common.DiscoveryI
}

func NewCoreDependency(cfg *config.Config, coreSrv *service.CoreService, srvDis common.DiscoveryI) *CoreDependency {
	return &CoreDependency{
		Config:           cfg,
		Service:          coreSrv,
		ServiceDiscovery: srvDis,
	}
}
