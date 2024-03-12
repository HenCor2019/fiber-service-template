package v1

import (
	"fmt"
	"os"

	health_rts "github.com/HenCor2019/fiber-service-template/api/v1/health"
	"github.com/HenCor2019/fiber-service-template/middleware/notfound"
	"github.com/spf13/cast"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type API struct {
	HealtRts health_rts.HealthRtr
}

func New(
	healthRt health_rts.HealthRtr,
) *API {
	return &API{
		HealtRts: healthRt,
	}
}

func (api *API) Start(app *fiber.App) error {
	app.Use(recover.New())

	v1 := app.Group("api/v1")
	v1.Route("healthcheck", api.HealtRts.Routes)

	v1.Use(notfound.NotFoundHandler)
	PORT := os.Getenv("PORT")
	return app.Listen(fmt.Sprintf(":%s", cast.ToString(PORT)))
}
