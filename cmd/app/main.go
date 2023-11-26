package main

import (
	"context"

	"github.com/mateusnasciment/APIGOLANG/cmd/app/controllers"
	"github.com/mateusnasciment/APIGOLANG/cmd/app/webserver"
	"github.com/mateusnasciment/APIGOLANG/internal/application/mappers"
	appservices "github.com/mateusnasciment/APIGOLANG/internal/application/services"
	cacherepositories "github.com/mateusnasciment/APIGOLANG/internal/infra/cache-repositories"
	"github.com/mateusnasciment/APIGOLANG/internal/infra/data"
	eventhandlers "github.com/mateusnasciment/APIGOLANG/internal/infra/event-handlers"
	"github.com/mateusnasciment/APIGOLANG/internal/infra/repositories"
	infraservices "github.com/mateusnasciment/APIGOLANG/internal/infra/services"
	"github.com/mateusnasciment/APIGOLANG/pkg/conf"
	"github.com/mateusnasciment/APIGOLANG/pkg/logger"
	"go.uber.org/fx"
)

func RegisterWebServer(lc fx.Lifecycle, ws webserver.WebServer) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go ws.StartServer()
			return nil
		},
		OnStop: func(_ context.Context) error {
			go ws.ShutdownServer()
			return nil
		},
	})
}

func main() {
	app := fx.New(
		conf.Module,
		logger.Module,
		data.Module,
		repositories.Module,
		cacherepositories.Module,
		mappers.Module,
		infraservices.Module,
		eventhandlers.Module,
		appservices.Module,
		controllers.Module,
		webserver.Module,
		fx.Invoke(RegisterWebServer),
	)
	app.Run()
}
