package v1

import (
	"github.com/asetalias/amigrant-server/internal/dto"
	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
}

func NewRouteHandler() *RouteHandler {
	return &RouteHandler{}
}

func (rh *RouteHandler) Ping(ctx *gin.Context) (*string, *dto.ErrorResponse) {
	response := "pong"

	return &response, nil
}
