package controller

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const SessionUserId = "uid"

func GetUserId(c *gin.Context) (int64, error) {
	session := sessions.Default(c)
	userId := session.Get(SessionUserId)
	if userId == nil {
		return 0, errors.New("未登录")
	}

	return userId.(int64), nil
}
