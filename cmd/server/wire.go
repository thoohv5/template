//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"

	"github.com/thoohv5/template/internal/data"
	"github.com/thoohv5/template/internal/pkg/app"
	"github.com/thoohv5/template/internal/pkg/config"
	"github.com/thoohv5/template/internal/pkg/log"
	"github.com/thoohv5/template/internal/router"
	"github.com/thoohv5/template/internal/server"
	"github.com/thoohv5/template/internal/service"
	pap "github.com/thoohv5/template/pkg/app"
)

// initApp init application.
func initApp() (pap.IApp, func(), error) {
	panic(wire.Build(config.ProviderSet, log.ProviderSet, data.ProviderSet, service.ProviderSet, router.ProviderSet, server.ProviderSet, app.ProviderSet))
}
