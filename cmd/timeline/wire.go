//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gomicroim/go-timeline/internal/biz"
	"github.com/gomicroim/go-timeline/internal/conf"
	"github.com/gomicroim/go-timeline/internal/data"
	"github.com/gomicroim/go-timeline/internal/server"
	"github.com/gomicroim/go-timeline/internal/service"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
