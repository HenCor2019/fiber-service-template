package main

import (
	"context"
	"net/http"

	health_cntlr "github.com/HenCor2019/fiber-service-template/internal/health/controllers"
	health_svc "github.com/HenCor2019/fiber-service-template/internal/health/services"

	"github.com/HenCor2019/fiber-service-template/api/config"
	v1 "github.com/HenCor2019/fiber-service-template/api/v1"
	health_rts "github.com/HenCor2019/fiber-service-template/api/v1/health"

	"go.uber.org/fx"
)

func main() {
	appModule := fx.Options(

		fx.Provide(
			health_rts.New,
			health_cntlr.New,
			health_svc.New,
		),

		fx.Provide(
			v1.New,
			http.NewServeMux,
			config.New,
		),

		fx.Invoke(setLifeCycle),
	)
	container := fx.New(appModule)
	container.Run()
}

func setLifeCycle(
	lc fx.Lifecycle,
	a *v1.API,
	router *config.RouteBundle,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go a.Start(router) // nolint

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
