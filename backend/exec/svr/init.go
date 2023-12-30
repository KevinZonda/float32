package main

import (
	"fmt"
	"github.com/KevinZonda/float32/exec/svr/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

var cli *openai.Client

func initAll() {
	godotenv.Load(".env")
	fmt.Println("Work at:", os.Getenv("PWD"))
	initChatGPT()
	initGin()
	initDb()
}

func initDb() {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		panic("DB_URL not set")
	}
	db.InitDb(dbUrl)
}

func initChatGPT() {
	token := os.Getenv("OPENAI")
	ep := os.Getenv("OPENAI_ENDPOINT")
	cfg := openai.DefaultConfig(token)
	if ep != "" {
		cfg.BaseURL = ep
	}
	cli = openai.NewClientWithConfig(cfg)
}

var g *gin.Engine

func initGin() {
	g = gin.Default()
	config := cors.DefaultConfig()
	if strings.TrimSpace(os.Getenv("DEBUG")) == "1" {
		gin.SetMode(gin.DebugMode)
		config.AllowAllOrigins = true
		g.Use(cors.New(config))
	} else {
		gin.SetMode(gin.ReleaseMode)
		config.AllowOrigins = []string{"https://float32.app"}
	}
	g.Use(cors.New(config))
}

func startGin() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = ":8080"
	}
	g.Run(listenAddr)
}
