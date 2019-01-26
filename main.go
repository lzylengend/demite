package main

import (
	"demite/conf"
	"demite/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()

	g.Use(gin.Recovery())
	g.Use(gin.Logger())

	c, err := conf.Init("./server.conf")
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Init(g)

	g.Run(c.Port)
}
