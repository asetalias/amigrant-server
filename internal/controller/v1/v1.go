package v1

import (
	v1 "github.com/asetalias/amigrant-server/internal/handlers/v1"
	"github.com/asetalias/amigrant-server/util"
	"github.com/gin-gonic/gin"
)

type Router struct {
	handlers *v1.RouteHandler
}

func NewRouter() *Router {
	return &Router{
		handlers: v1.NewRouteHandler(),
	}
}

func (r *Router) SetupRoutes(route *gin.RouterGroup) {
	v1 := route.Group("/v1")

	v1.GET("/ping", util.HandlerWrapper(r.handlers.Ping))
}
