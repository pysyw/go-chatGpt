package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PORT    string
	API_KEY string
	ENV     string
	Logger  struct {
		STATIC_BASE_PATH string
		FILENAME         string
		DEFAULT_DIVISION string
		DIVISION_TIME    struct {
			RotationTime string
			MAX_AGE      string
		}
	}
}

var GLOBAL_CONFIG Config

// TODO: use
func init() {
	viper.SetConfigFile("./config.toml")
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&GLOBAL_CONFIG)
	if err != nil {
		panic(err)
	}
}
