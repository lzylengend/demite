package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type RemoteApplyRequest struct {
	Phone         string `json:"phone"`
	Name          string `json:"name"`
	Hospital      string `json:"hospital"`
	Office        string `json:"office"`
	Faultdesc     string `json:"faultdesc"`
	FaultDescSelf string `json:"faultdescself"`
	Fileid1       string `json:"fileid1"`
	Fileid2       string `json:"fileid2"`
	RemoteTime    int64  `json:"remotetime"`
}

type RemoteApplyResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type RemoteApplyApi struct {
}

func (RemoteApplyApi) GetRequest() interface{} {
	return &RemoteApplyRequest{}
}

func (RemoteApplyApi) GetResponse() interface{} {
	return &RemoteApplyResponse{}
}

func (RemoteApplyApi) GetApi() string {
	return "RemoteApply"
}

func (RemoteApplyApi) GetDesc() string {
	return "远程申请"
}

func RemoteApply(c *gin.Context) {
	req := &RemoteApplyRequest{}
	rsp := &RemoteApplyResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	wxId, err := controller.GetWxUserId(c)
	if err != nil {
		rsp.Status = my_error.NoLoginError()
		c.JSON(200, rsp)
		return
	}

	_, err = model.RemoteDao.Apply(req.Phone, req.Name, req.Hospital, req.Office, req.RemoteTime,
		req.Faultdesc, req.FaultDescSelf, req.Fileid1, req.Fileid2, wxId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
