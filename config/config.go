package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//Config type for DevX
type DevXAppConfig struct {
	HealthCheckTime int
	DBDriver        string
	DBConfig        map[string]string
	// ReleaseMode     string
}

//ReadConfig load configuration information from file
func ReadConfig(configFile string) (*DevXAppConfig, error) {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return &DevXAppConfig{
			DBDriver: "sqlite3",
			DBConfig: map[string]string{},
			// ReleaseMode: gin.DebugMode,
		}, err
	}
	config := &DevXAppConfig{}
	json.Unmarshal(file, config)

	if config.DBDriver == "" {
		log.Println("Use sqlite3 as default")
		config.DBDriver = "sqlite3"
	}

	if config.DBConfig == nil {
		config.DBConfig = map[string]string{}
	}

	// if config.ReleaseMode == "" {
	// config.ReleaseMode = gin.DebugMode
	// }

	return config, err
}
