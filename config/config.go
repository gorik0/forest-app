package config

import (
	"errors"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	DbUser     string `mapstructure:"DB_USER"`
	DBName     string `mapstructure:"DB_NAME"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	HttpAddr   string `mapstructure:"HTTP_ADDR"`
}

var cfgEnv = "DB_USER DB_NAME DB_PASSWORD DB_PORT DB_HOST HTTP_ADDR"
var cfg *Config

func InitConfig() error {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return errors.New("Fail to read in cfg!!! --->>> " + err.Error())
	}

	for _, env := range strings.Split(cfgEnv, " ") {
		err := viper.BindEnv(env)
		if err != nil {
			return errors.New("error while binding cfg env ---->>>" + err.Error())
		}

	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return errors.New("Fail to unmarshall   cfg!!! --->>> " + err.Error())

	}
	return nil

}
func GetConfig() *Config {
	return cfg
}
