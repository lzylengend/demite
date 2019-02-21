package file_api

import (
	"demite/conf"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"io"
	"os"
)

type UploadFileRequest struct {
}

type UploadFileResponse struct {
	Id     string                `json:"id"`
	Status *my_error.ErrorCommon `json:"status"`
}

func UploadFile(c *gin.Context) {
	rsp := &UploadFileResponse{}
	id := uuid.NewV4().String()
	file, _, err := c.Request.FormFile("file")
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
