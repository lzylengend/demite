package wx_user_banken_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
	"time"
)

type UpdateWxUserRequest struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Hospital string `json:"hospital"`
	Position string `json:"position"`
}

type UpdateWxUserResponse struct {
	Data   *getWxData            `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type UpdateWxUserApi struct {
}

func (UpdateWxUserApi) GetRequest() interface{} {
	return &UpdateWxUserRequest{}
}

func (UpdateWxUserApi) GetResponse() interface{} {
	return &UpdateWxUserResponse{
		Data: &getWxData{},
	}
}

func (UpdateWxUserApi) GetApi() string {
	return "UpdateWxUser"
}

func (UpdateWxUserApi) GetDesc() string {
	return "修改用户"
}

func UpdateWxUser(c *gin.Context) {
	req := &UpdateWxUserRequest{}
	rsp := &UpdateWxUserResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	obj, err := model.WxUserDao.GetById(req.Id)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	obj.Name = req.Name
	obj.Position = req.Position
	obj.Hospital = req.Hospital
	obj.Phone = req.Phone
	obj.UpdateTime = time.Now().Unix()

	err = model.WxUserDao.Set(obj)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
