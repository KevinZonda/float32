package main

import (
	"fmt"
	"github.com/KevinZonda/float32/exec/svr/db"
	"github.com/KevinZonda/float32/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"io"
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

func main() {
	initAll()
	g.POST("/query", queryQuestion)
	g.GET("/history", ginHistory)
	startGin()
}

func writeBySubStr(w io.Writer, sb, buf *strings.Builder, delta string, subStr rune) (needContinue bool) {
	rs := []rune(delta)
	if idx := utils.IndexOfRunes(rs, subStr); idx > 0 {
		toPrint := buf.String() + string(rs[:idx+1])
		w.Write([]byte(toPrint))
		buf.Reset()
		sb.WriteString(string(rs[idx+1:]))
		needContinue = true
	}
	return
}

func writeBySubStrs(w io.Writer, sb, buf *strings.Builder, delta string, subStrs ...rune) (needContinue bool) {
	for _, subString := range subStrs {
		if writeBySubStr(w, sb, buf, delta, subString) {
			needContinue = true
			return
		}
	}
	return
}
