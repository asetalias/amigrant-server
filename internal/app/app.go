package app

import (
	"io"

	v1 "github.com/asetalias/amigrant-server/internal/controller/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}

	server.setupRoutes()

	return server
}

func (s *Server) setupRoutes() {
	router := gin.Default()

	baseRouter := router.Group("/amigrant")

	// disable gin logs
	gin.DefaultWriter = io.Discard
	// disable debug logs
	gin.SetMode(gin.ReleaseMode)

	// register all v1 routes
	v1Router := v1.NewRouter()
	v1Router.SetupRoutes(baseRouter)

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
	}))

	s.router = router
}

func (s *Server) Start(port string) error {
	return s.router.Run(":" + port)
}