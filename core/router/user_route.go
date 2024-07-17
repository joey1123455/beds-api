package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joey1123455/beds-api/core/handlers/users"
	"github.com/joey1123455/beds-api/core/server"
)

func RegisterUserRoute(srv *server.Server, router *gin.RouterGroup) {

	userGroup := router
	userGroup.GET("/:user_id", users.NewUserHandler(srv).GetUser)
	// userGroup.PATCH("/update", users.NewUserHandler(srv).UpdateUser)
	// userGroup.GET("/:id", users.NewUserHandler(srv).GetUser)
}
