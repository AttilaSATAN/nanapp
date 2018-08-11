package nanapp

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/attilasatan/nanapp/structure"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func initDB(connectionString string) (collection *mongo.Collection, err error) {
	client, err := mongo.NewClient(connectionString)
	if err != nil {
		return
	}
	err = client.Connect(context.TODO())
	if err != nil {
		return
	}

	collection = client.Database("nanoapp").Collection("campaigns") //TODO: this should be added to config and config tests
	return

}

func populateDB(dataFile string) (err error) {

	var data []structure.Campaign

	err = Collection.Drop(nil)
	if err != nil {
		return
	}
	raw, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return
	}
	if err = json.Unmarshal(raw, &data); err != nil {
		return
	}
	s := make([]interface{}, len(data))
	for i, v := range data {
		s[i] = v
	}
	res, err := Collection.InsertMany(nil, s)
	if err != nil {
		return
	}
	log.Println(res)
	return nil
}
