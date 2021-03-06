package class_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type AddClassRequest struct {
	Name      string `json:"name"`
	UpClassId int64  `json:"upclassid"`
}

type AddClassResponse struct {
	ClassId int64                 `json:"classid"`
	Status  *my_error.ErrorCommon `json:"status"`
}

type AddClassApi struct {
}

func (AddClassApi) GetRequest() interface{} {
	return &AddClassRequest{}
}

func (AddClassApi) GetResponse() interface{} {
	return &AddClassResponse{}
}

func (AddClassApi) GetApi() string {
	return "AddClass"
}

func (AddClassApi) GetDesc() string {
	return "新增分类"
}

func AddClass(c *gin.Context) {
	req := &AddClassRequest{}
	rsp := &AddClassResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	if req.Name == "" {
		rsp.Status = my_error.NotNilError("分组名")
		c.JSON(200, rsp)
		return
	}

	path := ""
	if req.UpClassId != 0 {
		upClass, err := model.ClassDao.GetClassById(req.UpClassId)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}
		path = upClass.Path
	}

	class, err := model.ClassDao.AddClass(req.Name, req.UpClassId, path)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.ClassId = class.ClassId
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
