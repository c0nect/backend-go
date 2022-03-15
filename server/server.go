package server

import (
	"MoviesBack/server/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   ":3000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	s.server.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowCredentials: true,
	}))

	r := routes.ConfigRoutes(s.server)
	log.Printf("Server is running on port %s", s.port)
	log.Fatal(r.Run(s.port))
}
