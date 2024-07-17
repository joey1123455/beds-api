package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joey1123455/beds-api/core/handlers/health"
	"github.com/joey1123455/beds-api/core/middleware"
	"github.com/joey1123455/beds-api/core/server"
)

func SetupRouter(srv *server.Server) {

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery()).Use(middleware.CORS())

	router.GET("/health", health.NewHealthHandler(srv).HealthCheck)

	RegisterUserRoute(srv, router.Group("/user"))
	RegisterAuthRoutes(srv, router.Group("/auth"))

	srv.Router = router

}
