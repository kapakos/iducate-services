package config

import (
	"github.com/BurntSushi/toml"
	"path/filepath"
	"log"
)

type AppConfig struct {
	Endpoints endpointInfo
}

type endpointInfo struct {
	CourseraUrl string  `toml:"courseraUrl"`
	UdacityUrl  string	`toml:"udacityUrl"`
}

func Get() AppConfig {
	var config AppConfig
	filePrefix, _ := filepath.Abs("./config/config.toml")
	if _, err := toml.DecodeFile(filePrefix, &config); err != nil {
		log.Fatal("Config file is missing: ", err)
	}
	return config;
}