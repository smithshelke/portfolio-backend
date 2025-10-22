package httphandler

import (
	"fmt"
	"net/http"

	"shelke.dev/api/internal/ports"
)

type HealthCheckHandler struct {
	healthCheckService ports.HealthCheckService
}

func NewHealthCheckHandler(healthCheckService ports.HealthCheckService) *HealthCheckHandler {
	return &HealthCheckHandler{
		healthCheckService: healthCheckService,
	}
}

func (h *HealthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	status := h.healthCheckService.CheckHealth()
	fmt.Fprintf(w, "Health: %s", status)
}
