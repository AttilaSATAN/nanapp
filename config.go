package nanapp

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type nanoConfig struct {
	MongoDBConnectionString string `env:"DB_CONNECTION_STRING" envDefault:"mongodb://localhost:27017"`
}

func isExistsDotEnv() bool {
	_, err := os.Stat(".env")
	if err != nil {

		return !os.IsNotExist(err)
	}
	return true
}

func initConfig() (*nanoConfig, error) {
	if isExistsDotEnv() {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	fmt.Println("ENV", os.Getenv("DB_CONNECTION_STRING"))
	conf := nanoConfig{}
	env.Parse(&conf)
	return &conf, nil
}
