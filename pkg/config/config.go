package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	FilePath        string `json:"file_path"`
	EnabledDbRecord int    `json:"enabled_db_record"`
	Database        Database
}

type Database struct {
	Host     string
	User     string
	Password string
	DbName   string `json:"db_name"`
}

func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return config, err
}
