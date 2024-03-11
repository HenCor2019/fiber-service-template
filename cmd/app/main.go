package main

import (
	"context"

	health_cntlr "github.com/HenCor2019/fiber-service-template/internal/health/controllers"
	health_svc "github.com/HenCor2019/fiber-service-template/internal/health/services"

	v1 "github.com/HenCor2019/fiber-service-template/api/v1"
	health_rts "github.com/HenCor2019/fiber-service-template/api/v1/health"

	"github.com/gofiber/fiber/v2"
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
			fiber.New,
		),

		fx.Invoke(setLifeCycle),
	)
	container := fx.New(appModule)
	container.Run()
}

func setLifeCycle(
	lc fx.Lifecycle,
	a *v1.API,
	app *fiber.App,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go a.Start(app) // nolint

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
