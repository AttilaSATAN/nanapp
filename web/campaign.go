package web

import (
	"context"
	"encoding/json"

	"github.com/attilasatan/nanapp/structure"
	"github.com/kataras/iris/mvc"
	"github.com/mongodb/mongo-go-driver/mongo"
)

//Campaign holds MVC components of Campaign resource
type Campaign struct {
	Collection *mongo.Collection
}

//GetCampaign is controller for endpoint: `/api/campaign`
func (s Campaign) Get() mvc.Result {

	c, err := s.Collection.Find(context.Background(), nil)
	if err != nil {
		return mvc.Response{Err: err, Code: 500}
	}

	defer c.Close(context.Background())

	var res []*structure.Campaign = make([]*structure.Campaign, 0)

	for c.Next(context.Background()) {
		camp := &structure.Campaign{}
		err := c.Decode(camp)
		if err != nil {
			return mvc.Response{Err: err, Code: 500}
		}
		res = append(res, camp)
	}
	bRes, err := json.Marshal(res)
	if err != nil {
		return mvc.Response{Err: err, Code: 500}
	}
	return mvc.Response{
		Code:        200,
		ContentType: "application/json",
		Content:     bRes,
	}
}
