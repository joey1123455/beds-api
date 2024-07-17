package users

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joey1123455/beds-api/core/server"
	"github.com/joey1123455/beds-api/internal/common"
	"github.com/joey1123455/beds-api/internal/domain"
	"github.com/joey1123455/beds-api/internal/logger"
)

type UserHandler struct {
	server *server.Server
}

func NewUserHandler(server *server.Server) *UserHandler {
	return &UserHandler{
		server: server,
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("user_id")

	parseUUID, err := uuid.Parse(id)
	if err != nil {
		server.SendBadRequest(c, common.ErrInvalidUUID)
		return
	}
	user, err := h.server.Store.GetUser(c, parseUUID)
	if err != nil {
		server.SendNotFound(c, common.ErrNotExist)
		return
	}

	profile, err := h.server.Store.GetUserProfileByUserID(c, parseUUID)
	if err != nil {
		logger.ErrorLogger(err)
		server.SendParsingError(c, common.ErrFailToProcess)
		return
	}

	userProfile := domain.UserResponses(profile, user)

	server.SendSuccess(c, "User retrieved successfully", userProfile)
}
