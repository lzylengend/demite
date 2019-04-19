package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GetRemoteApplyRequest struct {
	Id int64 `json:"id"`
}

type GetRemoteApplyResponse struct {
	Data   *getRemoteApply       `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type getRemoteApply struct {
	CreateTime    int64                 `json:"createtime"`
	Phone         string                `json:"phone"`
	Name          string                `json:"name"`
	Hospital      string                `json:"hospital"`
	Office        string                `json:"office"`
	Faultdesc     string                `json:"faultdesc"`
	FaultDescSelf string                `json:"faultdescself"`
	Fileid1       string                `json:"fileid1"`
	Fileid2       string                `json:"fileid2"`
	RemoteTime    int64                 `json:"remotetime"`
	List          []*getRemoteApplyData `json:"list"`
}

type getRemoteApplyData struct {
	StaffName  string `json:"staffname"`
	StaffPhone string `json:"staffphone"`
	RemoteTime int64  `json:"remotetime"`
	DealTime   int64  `json:"dealtime"`
	CreateTime int64  `json:"createtime"`
	Status     string `json:"currentstatus"`
	StaffNO    string `json:"staffno"`
	Reason     string `json:"reason"`
}

type GetRemoteApplyApi struct {
}

func (GetRemoteApplyApi) GetRequest() interface{} {
	return &GetRemoteApplyRequest{}
}

func (GetRemoteApplyApi) GetResponse() interface{} {
	return &GetRemoteApplyResponse{}
}

func (GetRemoteApplyApi) GetApi() string {
	return "GetRemoteApply"
}

func (GetRemoteApplyApi) GetDesc() string {
	return "获取远程维修信息"
}

func GetRemoteApply(c *gin.Context) {
	req := &GetRemoteApplyRequest{}
	rsp := &GetRemoteApplyResponse{}
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

	r, err := model.RemoteDao.GetByWIdAndxUserId(req.Id, wxId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Data = &getRemoteApply{
		RemoteTime:    r.RemoteTime,
		CreateTime:    r.CreateTime,
		Phone:         r.Phone,
		Name:          r.Name,
		Hospital:      r.Hospital,
		Office:        r.Office,
		Faultdesc:     r.FaultDesc,
		FaultDescSelf: r.FaultDescSelf,
		Fileid1:       r.FileId1,
		Fileid2:       r.FileId2,
	}

	objList, err := model.RemoteScheduleDao.ListByRemoteId(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	data := make([]*getRemoteApplyData, 0)
	for _, v := range objList {
		staff := &model.Staff{}
		if v.StaffId != 0 {
			staff, err = model.StaffDao.Get(v.StaffId)
			if err != nil {
				rsp.Status = my_error.DbError(err.Error())
				c.JSON(200, rsp)
				return
			}
		}

		data = append(data, &getRemoteApplyData{
			StaffName:  staff.StaffName,
			StaffPhone: staff.StaffPhone,
			StaffNO:    staff.StaffNO,
			RemoteTime: v.RemoteTime,
			DealTime:   v.DealTime,
			CreateTime: v.CreateTime,
			Status:     string(v.Status),
			Reason:     v.Reason,
		})
	}

	rsp.Data.List = data
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
