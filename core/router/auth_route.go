package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joey1123455/beds-api/core/handlers/auth"
	"github.com/joey1123455/beds-api/core/server"
	"github.com/joey1123455/beds-api/internal/security"
)

func RegisterAuthRoutes(srv *server.Server, router *gin.RouterGroup) {
	authGroup := router
	authGroup.POST("/login", auth.NewAuthHandler(srv, security.NewValidator(srv.Store, *srv.Config)).Login)
	authGroup.PATCH("/set-password", auth.NewAuthHandler(srv, security.NewValidator(srv.Store, *srv.Config)).PasswordSetting)
	authGroup.PATCH("/verify-email", auth.NewAuthHandler(srv, security.NewValidator(srv.Store, *srv.Config)).ActivateUser)
	authGroup.POST("/register", auth.NewAuthHandler(srv, security.NewValidator(srv.Store, *srv.Config)).RegisterUser)
	authGroup.POST("/create-profile", auth.NewAuthHandler(srv, security.NewValidator(srv.Store, *srv.Config)).CreateProfile)
	authGroup.GET("/request-password-reset", auth.NewAuthHandler(srv, security.NewValidator(srv.Store, *srv.Config)).RequestPasswordReset)
	authGroup.GET("/request-email-verification", auth.NewAuthHandler(srv, security.NewValidator(srv.Store, *srv.Config)).RequestUserActivation)
}
