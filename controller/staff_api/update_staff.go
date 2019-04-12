package staff_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type UpdateStaffRequest struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	StaffDecs  string `json:"staffdesc"`
	StaffPhone string `json:"staffphone"`
}

type UpdateStaffResponse struct {
	Status *my_error.ErrorCommon `json:"status"`
}

type UpdateStaffApi struct {
}

func (UpdateStaffApi) GetRequest() interface{} {
	return &UpdateStaffRequest{}
}

func (UpdateStaffApi) GetResponse() interface{} {
	return &UpdateStaffResponse{}
}

func (UpdateStaffApi) GetApi() string {
	return "UpdateStaff"
}

func (UpdateStaffApi) GetDesc() string {
	return "修改员工"
}

func UpdateStaff(c *gin.Context) {
	req := &UpdateStaffRequest{}
	rsp := &UpdateStaffResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if req.Name == "" {
		rsp.Status = my_error.NotNilError("员工名字")
		c.JSON(200, rsp)
		return
	}

	obj, err := model.StaffDao.Get(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	obj.StaffPhone = req.StaffPhone
	obj.StaffName = req.Name
	obj.StaffDecs = req.StaffDecs

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
