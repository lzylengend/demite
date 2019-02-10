package main

import (
	"demite/conf"
	"demite/model"
	_ "demite/mylog"
	"demite/router"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()

	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("demite00"))
	if err != nil {
		fmt.Println(err)
		return
	}

	g.Use(sessions.Sessions("s", store))
	g.Use(gin.Recovery())
	g.Use(gin.Logger())

	c, err := conf.Init("./server.conf")
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Init(g)

	err = model.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	g.Run(c.Port)
}
