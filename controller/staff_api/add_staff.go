package staff_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type AddStaffRequest struct {
	Name       string `json:"name"`
	StaffNO    string `json:"staffno"`
	StaffDecs  string `json:"staffdesc"`
	StaffPhone string `json:"staffphone"`
}

type AddStaffResponse struct {
	Id     int64                 `json:"id"`
	Status *my_error.ErrorCommon `json:"status"`
}

type AddStaffApi struct {
}

func (AddStaffApi) GetRequest() interface{} {
	return &AddStaffRequest{}
}

func (AddStaffApi) GetResponse() interface{} {
	return &AddStaffResponse{}
}

func (AddStaffApi) GetApi() string {
	return "AddStaff"
}

func (AddStaffApi) GetDesc() string {
	return "新增员工"
}

func AddStaff(c *gin.Context) {
	req := &AddStaffRequest{}
	rsp := &AddStaffResponse{}
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

	b, _, err := model.StaffDao.GetAndExistByNO(req.StaffNO)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if b {
		rsp.Status = my_error.IdExistError("员工号")
		c.JSON(200, rsp)
		return
	}

	s, err := model.StaffDao.Add(req.Name, req.StaffNO, req.StaffDecs, req.StaffPhone)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Id = s.StaffId
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
