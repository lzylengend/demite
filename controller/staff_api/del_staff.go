package staff_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type DelStaffRequest struct {
	Id int64 `json:"id"`
}

type DelStaffResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type DelStaffApi struct {
}

func (DelStaffApi) GetRequest() interface{} {
	return &DelStaffRequest{}
}

func (DelStaffApi) GetResponse() interface{} {
	return &DelStaffResponse{}
}

func (DelStaffApi) GetApi() string {
	return "DelStaff"
}

func (DelStaffApi) GetDesc() string {
	return "删除员工"
}

func DelStaff(c *gin.Context) {
	req := &DelStaffRequest{}
	rsp := &DelStaffResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	obj, err := model.StaffDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	obj.DataStatus = time.Now().Unix()

	err = model.StaffDao.Set(obj)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
