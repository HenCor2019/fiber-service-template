package controllers

import (
	"github.com/HenCor2019/fiber-service-template/internal/health/services"
	"github.com/gofiber/fiber/v2"
)

type HealthCheckController interface {
	HealthCheckHandler(ctx *fiber.Ctx) error
}

type Controller struct {
	healthCheckService services.HealthCheckService
}

func New(healthCheckService services.HealthCheckService) HealthCheckController {
	return &Controller{healthCheckService: healthCheckService}
}

func (c *Controller) HealthCheckHandler(ctx *fiber.Ctx) error {
	healthStatus := c.healthCheckService.CheckHealth()
	return ctx.SendString(healthStatus)
}
