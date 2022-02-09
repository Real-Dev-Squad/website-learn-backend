package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var Global config

type config struct {
	Env           string `mapstructure:"GO_ENV"`
	Port          string `mapstructure:"PORT"`
	FirestoreCred string `mapstructure:"FIRESTORE_CREDENTIALS"`
	CorsUrl       string `mapstructure:"CORS_URL"`
}

const LOCAL = "local"

var environmentToFileMap = map[string]string{
	"prod":    "production",
	"dev":     "development",
	"test":    "testing",
	"default": "default",
	"local":   "local",
}

func getConfigFile() string {
	var configFile = environmentToFileMap[os.Getenv("env")]
	if configFile == "" {
		configFile = LOCAL
	}
	return configFile
}

func loadConfig() (config config, err error) {
	var configFile = getConfigFile()
	fmt.Println("Loading configs from", configFile)

	viper.AddConfigPath("src/config")
	viper.SetConfigName(configFile)
	viper.SetConfigType("toml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func init() {
	var err error
	Global, err = loadConfig()
	if err != nil {
		log.Fatal("invalid application configuration: ", err)
	}
}
