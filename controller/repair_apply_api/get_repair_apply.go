package repair_apply_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GetRepairApplyRequest struct {
	Id int64 `json:"id"`
}

type GetRepairApplyResponse struct {
	Data   *repairApplyDetal     `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type repairDetal struct {
	UserName   string `json:"username"`
	WxUserName string `json:"wxusername"`
	StaffName  string `json:"staffname"`
	StaffPhone string `json:"staffphone"`
	RepairTime int64  `json:"repairtime"`
	CreateTime int64  `json:"createtime"`
	Status     string `json:"currentstatus"`
	Reason     string `json:"reason"`
}

type repairApplyDetal struct {
	GoodName  string         `json:"goodname"`
	GoodModel string         `json:"goodmodel"`
	Hospital  string         `json:"hospital"`
	Office    string         `json:"office"`
	Phone     string         `json:"phone"`
	Name      string         `json:"name"`
	FaultDesc string         `json:"faultdesc"`
	FaultType string         `json:"faulttype"`
	FileId1   string         `json:"fileid1"`
	FileId2   string         `json:"fileid2"`
	Status    string         `json:"status"`
	Data      []*repairDetal `json:"data"`
}

type GetRepairApplyApi struct {
}

func (GetRepairApplyApi) GetRequest() interface{} {
	return &GetRepairApplyRequest{}
}

func (GetRepairApplyApi) GetResponse() interface{} {
	return &GetRepairApplyResponse{}
}

func (GetRepairApplyApi) GetApi() string {
	return "GetRepairApply"
}

func (GetRepairApplyApi) GetDesc() string {
	return "列出请求"
}

func GetRepairApply(c *gin.Context) {
	req := &GetRepairApplyRequest{}
	rsp := &GetRepairApplyResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	obj, err := model.RepairDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	good, err := model.GoodsDao.GetByUUID(obj.GoodUUID)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	detail := &repairApplyDetal{
		GoodName:  good.GoodsName,
		GoodModel: obj.GoodModel,
		Hospital:  obj.Hospital,
		Office:    obj.Office,
		Phone:     obj.Phone,
		Name:      obj.Name,
		FaultDesc: obj.FaultDesc,
		FaultType: obj.FaultType,
		FileId1:   obj.FileId1,
		FileId2:   obj.FileId2,
		Status:    string(obj.Status),
	}

	objList, err := model.RepairScheduleDao.ListByRepairId(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	data := make([]*repairDetal, 0)
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

		data = append(data, &repairDetal{
			UserName:   createname,
			WxUserName: wxUserName,
			StaffName:  staff.StaffName,
			StaffPhone: staff.StaffPhone,
			RepairTime: v.RepairTime,
			CreateTime: v.CreateTime,
			Status:     string(v.Status),
			Reason:     v.Reason,
		})
	}

	rsp.Data = detail
	rsp.Data.Data = data
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
