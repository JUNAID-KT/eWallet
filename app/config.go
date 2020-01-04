package app

import (
	"github.com/BurntSushi/toml"
	"github.com/JUNAID-KT/eWallet/models"
)

var Config models.Configuration

// Config constants
const (
	defaultServiceHost = "0.0.0.0:8080"
	defaultLogLevel    = 4 //default log level set to INFO
	defaultElasticHost = "http://localhost:9200"
)

// Init : reads config file config.toml and create AppConfig
func Init() {

	if _, err := toml.Decode("config.toml", &Config); err != nil {
		Config.Server = defaultServiceHost
		Config.LogLevel = defaultLogLevel
		Config.ElasticAddress = defaultElasticHost
	}
	InitializeLogger()
}
