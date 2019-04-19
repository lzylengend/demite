package wx_user_api

import (
	"demite/controller"
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type GetRepairApplyRequest struct {
	Id int64 `json:"id"`
}

type GetRepairApplyResponse struct {
	Data   *getRepairApply       `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type getRepairApply struct {
	GoodModel  string                `json:"goodmodel"`
	GoodUUID   string                `json:"gooduuid"`
	GoodName   string                `json:"goodname"`
	GoodPic    string                `json:"goodpic"`
	CreateTime int64                 `json:"createtime"`
	Phone      string                `json:"phone"`
	Name       string                `json:"name"`
	Hospital   string                `json:"hospital"`
	Office     string                `json:"office"`
	Faultdesc  string                `json:"faultdesc"`
	Faulttype  string                `json:"faulttype"`
	Fileid1    string                `json:"fileid1"`
	Fileid2    string                `json:"fileid2"`
	List       []*getRepairApplyData `json:"list"`
}

type getRepairApplyData struct {
	StaffName  string `json:"staffname"`
	StaffPhone string `json:"staffphone"`
	RepairTime int64  `json:"repairtime"`
	CreateTime int64  `json:"createtime"`
	Status     string `json:"currentstatus"`
	StaffNO    string `json:"staffno"`
	Reason     string `json:"reason"`
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
	return "获取报修延保"
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

	wxId, err := controller.GetWxUserId(c)
	if err != nil {
		rsp.Status = my_error.NoLoginError()
		c.JSON(200, rsp)
		return
	}

	r, err := model.RepairDao.GetByWIdAndxUserId(req.Id, wxId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	good, err := model.GoodsDao.GetByUUID(r.GoodUUID)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Data = &getRepairApply{
		GoodUUID:   good.GoodsUUID,
		GoodModel:  good.GoodsModel,
		GoodName:   good.GoodsName,
		CreateTime: r.CreateTime,
		Phone:      r.Phone,
		Name:       r.Name,
		Hospital:   r.Hospital,
		Office:     r.Office,
		Faultdesc:  r.FaultDesc,
		Faulttype:  r.FaultType,
		Fileid1:    r.FileId1,
		Fileid2:    r.FileId2,
		GoodPic:    good.GoodsPic,
	}

	objList, err := model.RepairScheduleDao.ListByRepairId(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	data := make([]*getRepairApplyData, 0)
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

		data = append(data, &getRepairApplyData{
			StaffName:  staff.StaffName,
			StaffPhone: staff.StaffPhone,
			StaffNO:    staff.StaffNO,
			RepairTime: v.RepairTime,
			CreateTime: v.CreateTime,
			Status:     string(v.Status),
			Reason:     v.Reason,
		})
	}

	rsp.Data.List = data
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
