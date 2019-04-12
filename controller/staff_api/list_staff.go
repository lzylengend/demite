package staff_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListStaffRequest struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	Key    string `json:"key"`
}

type ListStaffResponse struct {
	Data   []*staff              `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type staff struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	StaffNO    string `json:"staffno"`
	StaffDecs  string `json:"staffdesc"`
	StaffPhone string `json:"staffphone"`
}

type ListStaffApi struct {
}

func (ListStaffApi) GetRequest() interface{} {
	return &ListStaffRequest{}
}

func (ListStaffApi) GetResponse() interface{} {
	return &ListStaffResponse{}
}

func (ListStaffApi) GetApi() string {
	return "ListStaff"
}

func (ListStaffApi) GetDesc() string {
	return "列出员工"
}

func ListStaff(c *gin.Context) {
	req := &ListStaffRequest{}
	rsp := &ListStaffResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	objList, err := model.StaffDao.List(req.Key, req.Limit, req.Offset)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	count, err := model.StaffDao.Count(req.Key)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	data := make([]*staff, 0)
	for _, v := range objList {
		data = append(data, &staff{
			Id:         v.StaffId,
			Name:       v.StaffName,
			StaffNO:    v.StaffNO,
			StaffDecs:  v.StaffDecs,
			StaffPhone: v.StaffPhone,
		})
	}

	rsp.Data = data
	rsp.Count = count
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
