package server

import (
	"bytes"
	"html/template"
	"log"
	"time"

	"github.com/dollarkillerx/survive/pkg/request"
	"github.com/dollarkillerx/survive/pkg/response"
	"github.com/gin-gonic/gin"
)

func (s *Server) registerRouter() {
	s.app.POST("/heartbeat", s.heartbeat)
	s.app.GET("/", s.dashboard)
}

func (s *Server) dashboard(ctx *gin.Context) {
	bufStr := bytes.NewBufferString("")
	parse, err := template.New("main").Parse(temp)
	if err != nil {
		log.Fatalln(err)
	}

	pNode := s.rNode
	err = parse.Execute(bufStr, map[string]interface{}{
		"Node": pNode,
	})
	ctx.Header("content-type", "text/html charset=utf-8")
	ctx.Writer.Write(bufStr.Bytes())
}

func (s *Server) heartbeat(ctx *gin.Context) {
	tiM := time.Now().UnixNano() / 1e6

	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(401, response.HeartbeatResponse{
			Code: 401,
		})
		return
	}

	var req request.HeartbeatRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, response.HeartbeatResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	node, ex := s.rNode[token]
	if !ex {
		ctx.JSON(401, response.HeartbeatResponse{
			Code: 401,
		})
		return
	}

	node.CPU = req.CPU
	node.CPUUsedPercent = req.CPUUsedPercent
	node.MemUsedPercent = req.MemUsedPercent
	node.LastPingTime = time.Now().Format("2006-01-02 15:04:05")
	node.Delay = tiM - req.Time
	node.IP = ctx.ClientIP()
	timeout := time.Now().Add(time.Second * 6)
	node.Timeout = &timeout

	s.rNode[token] = node

	ctx.JSON(200, response.HeartbeatResponse{Code: 0})
}
