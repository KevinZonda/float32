package main

import "github.com/KevinZonda/float32/exec/svr/handler"

func main() {
	initAll()
	g.POST("/query", handler.Search)
	g.GET("/history", handler.QueryHistory)
	startGin()
}
