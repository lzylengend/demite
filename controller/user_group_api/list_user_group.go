package user_group_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListUserGroupRequest struct {
}

type ListUserGroupResponse struct {
	Data   []*userGroup          `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type userGroup struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`

	AuthDelGoods     bool `json:"authdelgoods"`
	AuthShieldWxUser bool `json:"authshieldwxuser"`
	AuthUserManage   bool `json:"authusermanage"`
	AuthDelStaff     bool `json:"authdelstaff"`
}

type ListUserGroupApi struct {
}

func (ListUserGroupApi) GetRequest() interface{} {
	return &ListUserGroupRequest{}
}

func (ListUserGroupApi) GetResponse() interface{} {
	return &ListUserGroupResponse{
		Data: []*userGroup{
			&userGroup{},
		},
	}
}

func (ListUserGroupApi) GetApi() string {
	return "ListUserGroup"
}

func (ListUserGroupApi) GetDesc() string {
	return "列出组"
}

func ListUserGroup(c *gin.Context) {
	req := &ListUserGroupRequest{}
	rsp := &ListUserGroupResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	objList, err := model.UserGroupDao.List()
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	count, err := model.UserGroupDao.Count()
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	data := make([]*userGroup, 0)
	for _, v := range objList {
		data = append(data, &userGroup{
			Id:               v.UserGroupId,
			Name:             v.UserGroupName,
			AuthDelGoods:     v.AuthDelGoods,
			AuthShieldWxUser: v.AuthShieldWxUser,
			AuthUserManage:   v.AuthUserManage,
			AuthDelStaff:     v.AuthDelStaff,
		})
	}

	rsp.Data = data
	rsp.Count = count
	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
	return
}
