package nanapp

import (
	"github.com/mongodb/mongo-go-driver/mongo"
)

var (
	//Collection is Mongodb collection reference
	Collection *mongo.Collection
	//Config holds the main configuration options
	Config *nanoConfig
	//DataFile is the abs file path of data.json
	DataFile string
)

//Init is the main initiation point for the server.
func Init(dataFile string) (err error) {

	DataFile = dataFile

	Config, err := initConfig()
	if err != nil {
		return
	}

	Collection, err = initDB(Config.MongoDBConnectionString)
	if err != nil {
		return
	}
	err = populateDB(DataFile)
	if err != nil {
		return
	}
	return
}
