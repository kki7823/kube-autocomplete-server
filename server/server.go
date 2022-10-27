package server

import "github.com/gin-gonic/gin"

type server struct {
	router *gin.Engine
	config *ServerConfig
}

type ServerConfig struct {
	port string
}

func NewServer(o *Option) *server {
	cfg := &ServerConfig{
		port: o.Port,
	}

	s := &server{
		router: gin.Default(),
		config: cfg,
	}

	return s
}

func (s *server) Run() error {
	return s.router.Run(":" + s.config.port)
}
