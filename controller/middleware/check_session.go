package middleware

import (
	//	"demite/controller"
	"demite/my_error"

	"github.com/gin-gonic/gin"
)

type commonRespose struct {
	Status *my_error.ErrorCommon `json:"status"`
}

func CheckSession(c *gin.Context) {
	// rsp := &commonRespose{}
	// _, err := controller.GetUserId(c)

	// if err != nil {
	// 	rsp.Status = my_error.NoLoginError()
	// 	c.JSON(200, rsp)
	// 	c.Abort()
	// 	return
	// }

	c.Next()
}

func CheckWxSession(c *gin.Context) {
	//rsp := &commonRespose{}
	//_, err := controller.GetWxUserId(c)
	//
	//if err != nil {
	//	rsp.Status = my_error.NoLoginError()
	//	c.JSON(200, rsp)
	//	c.Abort()
	//	return
	//}

	c.Next()
}
