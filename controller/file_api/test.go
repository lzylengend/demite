package file_api

import (
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.String(200, `{"a":"b"}`)
	//c.JSON(200, `{"a":"b"}`)
	return
}
