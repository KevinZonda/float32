package rag

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

type IAgent interface {
	Search(query string) ([]string, error)
}

func RunAgent(a IAgent) {
	godotenv.Load(".env")
	gin.New()
	g := gin.New()
	g.Use(gin.Logger(), gin.Recovery())
	cs := cors.DefaultConfig()
	cs.AllowAllOrigins = true
	g.Use(cors.New(cs))
	g.POST("/query", func(c *gin.Context) {
		var req AgentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		rst, err := a.Search(req.Query)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, AgentResponse{Rst: rst})
	})
	listenAddr := os.Getenv("AGENT_LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = ":8080"
	}
	g.Run(listenAddr)
}

type AgentRequest struct {
	Query string `json:"query"`
}

type AgentResponse struct {
	Rst []string `json:"rst"`
}
