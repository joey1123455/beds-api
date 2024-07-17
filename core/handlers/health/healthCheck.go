package health

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joey1123455/beds-api/core/server"
)

type healthHandler struct {
	srv *server.Server
}

func NewHealthHandler(srv *server.Server) *healthHandler {
	return &healthHandler{srv: srv}
}
func (c *healthHandler) HealthCheck(ctx *gin.Context) {

	healthOutput := struct {
		Status string `json:"status"`
		Env    string `json:"env"`
		Host   string `json:"host"`
	}{
		Status: "ok",
		Host:   getOSHostName(),
		Env:    c.srv.Config.Environment,
	}
	ctx.JSON(http.StatusOK, healthOutput)
}

func getOSHostName() string {
	osHostName, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return osHostName
}
