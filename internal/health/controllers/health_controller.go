package controllers

import (
	"net/http"

	"github.com/HenCor2019/fiber-service-template/internal/health/services"
)

type HealthCheckController interface {
	HealthCheckHandler(w http.ResponseWriter, r *http.Request)
}

type Controller struct {
	healthCheckService services.HealthCheckService
}

func New(healthCheckService services.HealthCheckService) HealthCheckController {
	return &Controller{healthCheckService: healthCheckService}
}

func (c *Controller) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	healthStatus := c.healthCheckService.CheckHealth()

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(healthStatus)); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
