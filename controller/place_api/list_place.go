package place_api

import (
	"demite/model"
	"demite/my_error"
	"github.com/gin-gonic/gin"
)

type ListPlaceRequest struct {
	UpPlaceId int64 `json:"upplaceid"`
}

type ListPlaceResponse struct {
	Data   []*placeData          `json:"data"`
	Status *my_error.ErrorCommon `json:"status"`
}

type placeData struct {
	PlaceId   int64  `json:"placeid"`
	PlaceName string `json:"placename"`
}

type ListPlaceApi struct {
}

func (ListPlaceApi) GetRequest() interface{} {
	return &ListPlaceRequest{}
}

func (ListPlaceApi) GetResponse() interface{} {
	return &ListPlaceResponse{
		Data: []*placeData{
			&placeData{},
		},
	}
}

func (ListPlaceApi) GetApi() string {
	return "ListPlace"
}

func (ListPlaceApi) GetDesc() string {
	return "列出地域"
}

func ListPlace(c *gin.Context) {
	req := &ListPlaceRequest{}
	rsp := &ListPlaceResponse{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Status = my_error.JsonError(err.Error())
		c.JSON(200, rsp)
		return
	}

	res, err := model.PlaceDao.ListByUpId(req.UpPlaceId)
	if err != nil {
		rsp.Status = my_error.DbError(err.Error())
		c.JSON(200, rsp)
		return
	}

	rsp.Data = make([]*placeData, 0)
	for _, v := range res {
		rsp.Data = append(rsp.Data, &placeData{
			PlaceId:   v.PlaceId,
			PlaceName: v.PlaceName,
		})
	}

	rsp.Status = my_error.NoError()
	c.JSON(200, rsp)
}
