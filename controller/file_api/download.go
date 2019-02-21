package file_api

import (
	"demite/conf"
	"demite/my_error"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type DownloadFileRequest struct {
	Id string `json:"id"`
}

type DownloadFileResponse struct {
	Data   string                `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

func DownloadFile(c *gin.Context) {
	req := &DownloadFileRequest{}
	rsp := &DownloadFileResponse{}

	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	data, err := ioutil.ReadFile(conf.GetFilePath() + "/" + req.Id)
	if err != nil {
		rsp.Status = my_error.FileReadError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Data = base64.StdEncoding.EncodeToString(data)
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
