package nanapp

import (
	"os"
	"reflect"
	"testing"
)

func Test_initConfig(t *testing.T) {

	tests := []struct {
		name                     string
		want                     *nanoConfig
		envVars                  map[string]string
		checkDotEnvFileExistence bool
		customDotEnvFile         map[string]string
		wantErr                  bool
	}{
		{"default", &nanoConfig{"mongodb://localhost:27017"}, nil, false, nil, false},
		{"from file", &nanoConfig{"mongodb://localhost:27018"}, nil, false, map[string]string{"DB_CONNECTION_STRING": "mongodb://localhost:27018"}, false},
		{"from env var", &nanoConfig{"mongodb://localhost:27019"}, map[string]string{"DB_CONNECTION_STRING": "mongodb://localhost:27019"}, false, nil, false},
		{"file and env var", &nanoConfig{"mongodb://localhost:27021"}, map[string]string{"DB_CONNECTION_STRING": "mongodb://localhost:27021"}, false, map[string]string{"DB_CONNECTION_STRING": "mongodb://localhost:27020"}, false}}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if tt.checkDotEnvFileExistence && !isExistsDotEnv() {
				t.Errorf("initConfig() .env does not exists shouldHaveDotEnvFile = true")
				return
			}

			if tt.customDotEnvFile != nil {

				rename := false

				if isExistsDotEnv() {
					rename = true
					err := os.Rename(".env", "dotenv.bk")
					if err != nil {
						t.Error("initConfig() could not rename .env file for test, check your .env file, it could be renamed to dotenv.bk")
					}

					defer os.Rename("dotenv.bk", ".env")
				}

				file, err := os.OpenFile(".env", os.O_RDWR|os.O_CREATE, 0666)
				if err != nil {
					t.Error("initConfig() could not rename .env file for test, check your .env file, it could be renamed to dotenv.bk")
				}

				for k, v := range tt.customDotEnvFile {
					file.WriteString(k + "=" + v)
					if err != nil {
						t.Error("initConfig() could not populate .env file with env vars")
					}
				}

				file.Close()

				if err != nil {
					t.Error("initConfig() error while closing file")
				}

				if !rename {
					defer os.Remove(".env")
					if err != nil {
						t.Error("initConfig() error while removing custom .env file")
					}
				}
			}

			if tt.envVars != nil {
				for k, v := range tt.envVars {
					if err := os.Setenv(k, v); err != nil {
						t.Error("initConfig() there has been a problem in setting env map")
					}
				}

				defer os.Clearenv()
			}

			got, err := initConfig()

			if (err != nil) != tt.wantErr {
				t.Errorf("initConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
