package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Configuration {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Gagal Membaca file config, %s", err.Error())
	}

	var conf Configuration
	conf.DB_USERNAME = viper.GetString("DB_USERNAME")
	conf.DB_PASSWORD = viper.GetString("DB_PASSWORD")
	conf.DB_PORT = viper.GetString("DB_PORT")
	conf.DB_HOST = viper.GetString("DB_HOST")
	conf.DB_NAME = viper.GetString("DB_NAME")

	return conf
}
