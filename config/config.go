/*

Manages the config variables for this web application. Loads config
from config.json and environment variables if they exist

*/

package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port        int    `json:"port"`
	Debug       bool   `json:"debug"`
	DatabaseURL string `json:"database_url"`
}

// This loads config-dev or config-prod depending on what environment it detects
// The environment is specified using a ENVIRONMENT_TYPE=PROD environment variable set
// on the Heroku/Docker environment
// Returns true if success, false if not
func LoadConfig() (bool, *Config) {
	config := new(Config)
	configEnvironment := os.Getenv("ENVIRONMENT_TYPE")
	// Default to the development config since it is safer
	configFileName := "config-dev.json"

	if configEnvironment == "PROD" {
		// Development Environment
		configFileName = "config-prod.json"
	}

	// Open our jsonFile
	jsonFile, err := os.Open(configFileName)
	// if we os.Open returns an error then handle it and return an empty config
	if err != nil {
		fmt.Println(err)
		return false, config
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	jsonParser := json.NewDecoder(jsonFile)

	parseError := jsonParser.Decode(config)
	if parseError != nil {
		// Return config as is, program will end anyway as we return false here
		fmt.Println(err)
		return false, config
	}
	return true, config
}
