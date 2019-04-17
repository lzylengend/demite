package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListRemoteApplyRequest struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type ListRemoteApplyResponse struct {
	Data   []*remote             `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type remote struct {
	Id          int64  `json:"id"`
	RemoteTime  int64  `json:"remotetime"`
	ApplyStatus string `json:"applystatus"`
	CreateTime  int64  `json:"createtime"`
}

type ListRemoteApplyApi struct {
}

func (ListRemoteApplyApi) GetRequest() interface{} {
	return &ListRemoteApplyRequest{}
}

func (ListRemoteApplyApi) GetResponse() interface{} {
	return &ListRemoteApplyResponse{}
}

func (ListRemoteApplyApi) GetApi() string {
	return "ListRemoteApply"
}

func (ListRemoteApplyApi) GetDesc() string {
	return "列出远程申请"
}

func ListRemoteApply(c *gin.Context) {
	req := &ListRemoteApplyRequest{}
	rsp := &ListRemoteApplyResponse{}
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

	objList, err := model.RemoteDao.ListByWxUserId(wxId, req.Limit, req.Offset)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	count, err := model.RemoteDao.CountByWxUserId(wxId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	data := make([]*remote, 0)
	for _, v := range objList {
		data = append(data, &remote{
			RemoteTime:  v.RemoteTime,
			CreateTime:  v.CreateTime,
			ApplyStatus: string(v.Status),
			Id:          v.RemoteId,
		})
	}

	rsp.Data = data
	rsp.Count = count
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
