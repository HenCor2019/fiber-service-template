package v1

import (
	"fmt"
	"net/http"
	"os"

	"github.com/HenCor2019/fiber-service-template/api/config"
	health_rts "github.com/HenCor2019/fiber-service-template/api/v1/health"
)

const (
	PREFIX  = "/api"
	VERSION = "v1"
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

func (api *API) Start(v1 *config.RouteBundle) error {
	PORT := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", PORT)

	fmt.Println("Server started on port ", PORT)
	return http.ListenAndServe(addr, v1)
}
