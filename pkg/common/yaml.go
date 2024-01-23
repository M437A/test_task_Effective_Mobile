package common

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func init_yaml_file() {
	wd, err := os.Getwd()
	if err != nil {
		panic("Can't get current working directory")
	}

	configPath := filepath.Join(wd, "resourse")
	viper.AddConfigPath(configPath)
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		panic("Can't connect to yaml, error: " + err.Error())
	}
}
