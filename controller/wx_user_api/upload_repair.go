package wx_user_api

import (
	"demite/controller"
	//"demite/model"
	"demite/my_error"

	"github.com/gin-gonic/gin"
)

type UploadRepairRequest struct {
	GoodUUID string `json:"gooduuid"`
}

type UploadRepairResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

func UploadRepair(c *gin.Context) {
	req := &UploadRepairRequest{}
	rsp := &UploadRepairResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	_, err = controller.GetWxUserId(c)
	if err != nil {
		rsp.Status = my_error.NoLoginError()
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
