// Code generated by sysl DO NOT EDIT.
package simple

import (
	"github.com/anz-bank/sysl-go/codegen/tests/deps"
	"github.com/anz-bank/sysl-go/codegen/tests/downstream"
	"github.com/anz-bank/sysl-go/config"
	"github.com/anz-bank/sysl-go/core"
	"github.com/anz-bank/sysl-go/handlerinitialiser"
)

// DownstreamClients for Simple
type DownstreamClients struct {
	depsClient       *deps.Client
	downstreamClient *downstream.Client
}

// BuildRestHandlerInitialiser ...
func BuildRestHandlerInitialiser(serviceInterface ServiceInterface, callback core.RestGenCallback, downstream *DownstreamClients) handlerinitialiser.HandlerInitialiser {
	serviceHandler := NewServiceHandler(callback, &serviceInterface, downstream.depsClient, downstream.downstreamClient)
	serviceRouter := NewServiceRouter(callback, serviceHandler)
	return serviceRouter
}

// BuildDownstreamClients ...
func BuildDownstreamClients(cfg *config.DefaultConfig) (*DownstreamClients, error) {
	var err error = nil
	depsHTTPClient, depsErr := core.BuildDownstreamHTTPClient("deps", &cfg.GenCode.Downstream.(*DownstreamConfig).Deps)
	downstreamHTTPClient, downstreamErr := core.BuildDownstreamHTTPClient("downstream", &cfg.GenCode.Downstream.(*DownstreamConfig).Downstream)
	if depsErr != nil {
		return nil, depsErr
	}

	if downstreamErr != nil {
		return nil, downstreamErr
	}

	depsClient := deps.NewClient(depsHTTPClient, cfg.GenCode.Downstream.(*DownstreamConfig).Deps.ServiceURL)
	downstreamClient := downstream.NewClient(downstreamHTTPClient, cfg.GenCode.Downstream.(*DownstreamConfig).Downstream.ServiceURL)
	return &DownstreamClients{depsClient: depsClient,
		downstreamClient: downstreamClient,
	}, err
}

// NewDefaultConfig ...
func NewDefaultConfig() config.DefaultConfig {
	return config.DefaultConfig{
		Library: config.LibraryConfig{},
		GenCode: config.GenCodeConfig{Downstream: &DownstreamConfig{}},
	}

}