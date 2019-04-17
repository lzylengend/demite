package remote_apply_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GetRemoteApplyRequest struct {
	Id int64 `json:"id"`
}

type GetRemoteApplyResponse struct {
	Data   []*remoteDetal        `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type remoteDetal struct {
	UserName   string `json:"username"`
	WxUserName string `json:"wxusername"`
	StaffName  string `json:"staffname"`
	StaffPhone string `json:"staffphone"`
	RemoteTime int64  `json:"remotetime"`
	DealTime   int64  `json:"dealtime"`
	CreateTime int64  `json:"createtime"`
	Status     string `json:"currentstatus"`
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
	return "列出请求"
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

	objList, err := model.RemoteScheduleDao.ListByRemoteId(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	data := make([]*remoteDetal, 0)
	for _, v := range objList {
		wxUserName := ""
		if v.WxUserId != 0 {
			wxuser, err := model.WxUserDao.GetById(v.WxUserId)
			if err != nil {
				rsp.Status = my_error.DbError(err.Error())
				c.JSON(200, rsp)
				return
			}
			wxUserName = wxuser.NickName
		}

		createname := ""
		if v.CreateId != 0 {
			user, err := model.UserDao.GetById(v.CreateId)
			if err != nil {
				rsp.Status = my_error.DbError(err.Error())
				c.JSON(200, rsp)
				return
			}

			createname = user.UserName
		}

		staff := &model.Staff{}
		if v.StaffId != 0 {
			staff, err = model.StaffDao.Get(v.StaffId)
			if err != nil {
				rsp.Status = my_error.DbError(err.Error())
				c.JSON(200, rsp)
				return
			}
		}

		data = append(data, &remoteDetal{
			UserName:   createname,
			WxUserName: wxUserName,
			StaffName:  staff.StaffName,
			StaffPhone: staff.StaffPhone,
			RemoteTime: v.RemoteTime,
			CreateTime: v.CreateTime,
			DealTime:   v.DealTime,
			Status:     string(v.Status),
		})
	}

	rsp.Data = data
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
