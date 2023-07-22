package util

import (
	"errors"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	// Viper uses mapstructure to decode the environment variables into the Config struct
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path ...string) (config Config, err error) {
	_, b, _, _ := runtime.Caller(0)
	dir := filepath.Dir(b) + "/.."

	if len(path) > 0 {
		dir = path[0]
	}

	viper.AddConfigPath(dir)
	// set name of .env file
	viper.SetConfigName("app")
	// could be json/xml etc, in this case is .env
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return config, errors.New("Config file not found")
		}

		return
	}

	err = viper.Unmarshal(&config)
	return
}
