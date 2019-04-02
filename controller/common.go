package controller

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const SessionUserId = "uid"
const SessionWxUserId = "wxuid"

func GetUserId(c *gin.Context) (int64, error) {
	session := sessions.Default(c)
	userId := session.Get(SessionUserId)
	if userId == nil {
		return 0, errors.New("未登录")
	}

	return userId.(int64), nil
}

func GetWxUserId(c *gin.Context) (int64, error) {
	return 1, nil
	session := sessions.Default(c)
	userId := session.Get(SessionWxUserId)
	if userId == nil {
		return 0, errors.New("未登录")
	}

	return userId.(int64), nil
}
