package main

import (
	"demite/conf"
	"demite/express_api"
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

	express_api.QurryExpress("")

	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("demite00"))
	if err != nil {
		fmt.Println(err)
		return
	}
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24 * 30})

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

	err = g.Run(c.Port)
	if err != nil {
		fmt.Println(err)
		return
	}
}
