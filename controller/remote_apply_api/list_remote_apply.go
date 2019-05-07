package remote_apply_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListRemoteRequest struct {
	Limit       int64  `json:"limit"`
	Offset      int64  `json:"offset"`
	Name        string `json:"name"`
	ApplyStatus string `json:"applystatus"`
}

type ListRemoteResponse struct {
	Data   []*remote             `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type remote struct {
	Id            int64  `json:"id"`
	Hospital      string `json:"hospital"`
	Office        string `json:"office"`
	Phone         string `json:"phone"`
	Name          string `json:"name"`
	FaultDesc     string `json:"faultdesc"`
	FaultDescSelf string `json:"faultdescself"`
	Status        string `json:"applystatus"`
	FileId1       string `json:"fileid1"`
	FileId2       string `json:"fileid2"`
	RemoteTime    int64  `json:"remotetime"`
}

type ListRemoteApi struct {
}

func (ListRemoteApi) GetRequest() interface{} {
	return &ListRemoteRequest{}
}

func (ListRemoteApi) GetResponse() interface{} {
	return &ListRemoteResponse{}
}

func (ListRemoteApi) GetApi() string {
	return "ListRemote"
}

func (ListRemoteApi) GetDesc() string {
	return "列出请求"
}

func ListRemote(c *gin.Context) {
	req := &ListRemoteRequest{}
	rsp := &ListRemoteResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	objList, err := model.RemoteDao.ListByStatus(req.Name, req.Limit, req.Offset, req.ApplyStatus)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	count, err := model.RemoteDao.CountByStatus(req.Name, req.ApplyStatus)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	data := make([]*remote, 0)
	for _, v := range objList {
		data = append(data, &remote{
			Id:            v.RemoteId,
			Hospital:      v.Hospital,
			Office:        v.Office,
			Phone:         v.Phone,
			Name:          v.Name,
			FaultDesc:     v.FaultDesc,
			FaultDescSelf: v.FaultDescSelf,
			FileId1:       v.FileId1,
			FileId2:       v.FileId2,
			Status:        string(v.Status),
			RemoteTime:    v.RemoteTime,
		})
	}

	rsp.Data = data
	rsp.Count = count
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
