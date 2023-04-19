package config

import (
	"log"

	"github.com/spf13/viper"
)

var Global config

type config struct {
	Env           string `mapstructure:"GO_ENV"`
	Version       int    `mapstructure:"VERSION"`
	Port          int    `mapstructure:"PORT"`
	CookieName    string `mapstructure:"COOKIE_NAME"`
	FirestoreCred string `mapstructure:"FIRESTORE_CREDENTIALS"`
	CorsUrl       string `mapstructure:"CORS_URL"`
	PublicKey     string `mapstructure:"PUBLIC_KEY"`
}

var environmentToFileMaps = map[string]string{
	"prod":    "production",
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

	if env == "test" {
		viper.AddConfigPath("../config")
	} else {
		viper.AddConfigPath("src/config")
	}

	viper.SetConfigName("default")
	viper.ReadInConfig()

	viper.SetConfigName(configFile)
	viper.MergeInConfig()

	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	return
}

func Setup(env string) {
	if env == "" {
		env = "local"
	}
	var err error
	Global, err = loadConfig(env)
	if err != nil {
		log.Fatal("Invalid application configuration: ", err)
	}
}
