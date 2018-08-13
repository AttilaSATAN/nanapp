package nanapp

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/attilasatan/nanapp/structure"

	"github.com/kataras/iris/httptest"
)

func TestAPIIntegration(t *testing.T) {
	err := Init("data.json")
	if err != nil {
		log.Fatal(err)
	}

	e := httptest.New(t, App)

	e.GET("/").Expect().Status(httptest.StatusOK).
		ContentType("text/html", "utf-8").Body().Equal("<h1>Hello Nanocorp!</h1>")

	cur, err := Collection.Find(nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	var res []*structure.Campaign = make([]*structure.Campaign, 0)
	for cur.Next(nil) {
		camp := &structure.Campaign{}
		err := cur.Decode(camp)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, camp)
	}

	campaignJSON, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	e.GET("/api/campaign").Expect().Status(httptest.StatusOK).
		Body().Equal(string(campaignJSON))

}
