package server

import (
	"sync"

	"github.com/dollarkillerx/survive/internel/conf"
	"github.com/dollarkillerx/survive/pkg/models"
	"github.com/gin-gonic/gin"
)

type Server struct {
	app *gin.Engine

	mu    sync.Mutex
	rNode map[string]*models.Node
}

func NewServer() *Server {
	ser := &Server{
		app:   gin.Default(),
		rNode: map[string]*models.Node{},
	}

	for _, v := range conf.CONFIG.Nodes {
		ser.rNode[v.Token] = &models.Node{
			Name:  v.Name,
			Token: v.Token,
		}
	}

	ser.registerRouter()
	return ser
}

func (s *Server) Run() error {
	go s.detect()
	return s.app.Run(conf.CONFIG.ListenAddr)
}
