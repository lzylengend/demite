package wx_user_api

import (
	"demite/conf"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"io"
	"os"
)

type UploadFileRequest struct {
}

type UploadFileResponse struct {
	Id     string                `json:"id"`
	Status *my_error.ErrorCommon `json:"status"`
}

type UploadFileApi struct {
}

func (UploadFileApi) GetRequest() interface{} {
	return &UploadFileRequest{}
}

func (UploadFileApi) GetResponse() interface{} {
	return &UploadFileResponse{}
}

func (UploadFileApi) GetApi() string {
	return "UploadFile"
}

func (UploadFileApi) GetDesc() string {
	return "上传文件，formdata，文件key为 file"
}

func UploadFile(c *gin.Context) {
	rsp := &UploadFileResponse{}

	id := uuid.NewV4().String()
	file, _, err := c.Request.FormFile("iFile")
	if err != nil {
		rsp.Status = my_error.FileParseError(err.Error())
		c.JSON(200, rsp)
		return
	}

	f, err := os.Create(conf.GetFilePath() + "/" + id)
	if err != nil {
		rsp.Status = my_error.FileWriteError(err.Error())
		c.JSON(200, rsp)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		rsp.Status = my_error.FileWriteError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Id = id
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
