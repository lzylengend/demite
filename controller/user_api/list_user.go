package user_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListUserRequest struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	Key    string `json:"key"`
}

type ListUserResponse struct {
	Data   []*userData           `json:"data"`
	Count  int64                 `json:"count"`
	Status *my_error.ErrorCommon `json:"status"`
}

type userData struct {
	Name      string `json:"name"`
	GroupName string `json:"groupname"`
	Id        int64  `json:"id"`
}

func ListUser(c *gin.Context) {
	req := &ListUserRequest{}
	rsp := &ListUserResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	res, err := model.UserDao.ListByKey(req.Limit, req.Offset, req.Key)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	n, err := model.UserDao.CountByKey(req.Key)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	resList := make([]*userData, 0)
	for _, v := range res {
		g, err := model.UserGroupDao.Get(v.UserGroupId)
		if err != nil {
			rsp.Status = my_error.DbError(err.Error())
			c.JSON(200, rsp)
			return
		}

		resList = append(resList, &userData{
			Name:      v.UserName,
			Id:        v.UserId,
			GroupName: g.UserGroupName,
		})
	}

	rsp.Data = resList
	rsp.Status = my_error.NoError()
	rsp.Count = n
	c.JSON(200, rsp)
	return

}
