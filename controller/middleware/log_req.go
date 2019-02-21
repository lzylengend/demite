package middleware

import (
	"bytes"
	"demite/my_error"
	"demite/mylog"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LogReq(c *gin.Context) {
	rsp := &commonRespose{}
	data, err := c.GetRawData()
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		c.Abort()
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	if strings.Contains(c.Request.URL.String(), "file") {
		if len(data) >= 1000 {
			data = []byte("文件数据")
		}
	}

	c.Next()

	rspData := blw.body.String()
	if strings.Contains(c.Request.URL.String(), "file") {
		if len(blw.body.String()) >= 1000 {
			rspData = "文件数据"
		}
	}

	mylog.LogInfo("path:" + c.Request.URL.String() + "| req:" + string(data) + "| rsp:" + rspData)
}
