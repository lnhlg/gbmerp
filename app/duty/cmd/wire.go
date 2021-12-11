//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"gbmerp/app/duty/internal/biz"
	"gbmerp/app/duty/internal/conf"
	"gbmerp/app/duty/internal/data"
	"gbmerp/app/duty/internal/server"
	"gbmerp/app/duty/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(
	*conf.Server,
	*conf.Data,
	*conf.Registry,
	log.Logger,
	*tracesdk.TracerProvider,
) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
