package routes

import (
	health_ctnlr "github.com/HenCor2019/fiber-service-template/internal/health/controllers"
	"github.com/gofiber/fiber/v2"
)

type HealthRtr interface {
	Routes(router fiber.Router)
}

type Rtr struct {
	health_ctnlr health_ctnlr.HealthCheckController
}

func New(health_ctnlr health_ctnlr.HealthCheckController) HealthRtr {
	return &Rtr{health_ctnlr}
}

func (module *Rtr) Routes(router fiber.Router) {
	router.Get("/", module.health_ctnlr.HealthCheckHandler)

}
