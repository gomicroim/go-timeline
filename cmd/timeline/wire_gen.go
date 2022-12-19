// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gomicroim/go-timeline/internal/biz"
	"github.com/gomicroim/go-timeline/internal/conf"
	"github.com/gomicroim/go-timeline/internal/data"
	"github.com/gomicroim/go-timeline/internal/server"
	"github.com/gomicroim/go-timeline/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	sequenceCache := data.NewSequence(dataData)
	messageUseCase := biz.NewMessageUseCase(dataData, confData, sequenceCache, logger)
	timelineService := service.NewTimelineService(messageUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, timelineService, logger)
	httpServer := server.NewHTTPServer(confServer, timelineService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
