package nanapp

import (
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type nanoConfig struct {
	MongoDBConnectionString string `env:"DB_CONNECTION_STRING" envDefault:"mongodb://127.0.0.1:27017"`
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

	conf := nanoConfig{}
	env.Parse(&conf)
	return &conf, nil
}
