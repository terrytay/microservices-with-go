package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/terrytay/microservices-with-go/product-api/data"
	"github.com/terrytay/microservices-with-go/product-api/utils"
)

type Health struct {
	l *log.Logger
}

func NewHealth(l *log.Logger) *Health {
	return &Health{l}
}

func (h Health) HealthCheck(w http.ResponseWriter, r *http.Request) {
	h.l.Println("[DEBUG] Health Status Check")
	healthResponse := data.Health{Message: "OK", Success: true, Timestamp: time.Now().UnixNano()}

	err := utils.ToJSON(healthResponse, w)
	if err != nil {
		h.l.Println("[ERROR] health check", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
