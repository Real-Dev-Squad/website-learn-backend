package config

import (
	"log"

	"github.com/spf13/viper"
)

var Global config

type config struct {
	Env           string `mapstructure:"GO_ENV"`
	Port          int    `mapstructure:"PORT"`
	Cookie        string `mapstructure:"COOKIE_NAME"`
	FirestoreCred string `mapstructure:"FIRESTORE_CREDENTIALS"`
	CorsUrl       string `mapstructure:"CORS_URL"`
}

var environmentToFileMaps = map[string]string{
	"prod":    "production",
	"dev":     "development",
	"test":    "testing",
	"default": "default",
	"local":   "local",
}

func getConfigFile(env string) string {
	var configFile = environmentToFileMaps[env]
	return configFile
}

func loadConfig(env string) (config config, err error) {
	var configFile = getConfigFile(env)

	viper.AddConfigPath("src/config")

	viper.SetConfigName("default")
	viper.ReadInConfig()

	viper.SetConfigName(configFile)
	viper.MergeInConfig()

	err = viper.Unmarshal(&config)
	return
}

func Setup(env string) {
	var err error
	Global, err = loadConfig(env)
	if err != nil {
		log.Fatal("Invalid application configuration: ", err)
	}
}
