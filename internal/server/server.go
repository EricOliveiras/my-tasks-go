package server

import (
	"github/ericoliveiras/basic-crud-go/internal/config"
	"github/ericoliveiras/basic-crud-go/internal/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Gin    *gin.Engine
	DB     *gorm.DB
	Config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Gin:    gin.Default(),
		DB:     db.Init(cfg),
		Config: cfg,
	}
}

func (server *Server) Start(addr string) error {
	return server.Gin.Run(":" + addr)
}
