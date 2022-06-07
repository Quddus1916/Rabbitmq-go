package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	Username     string `mapstructure:"USERNAME"`
	Email        string `mapstructure:"EMAIL"`
	SmtpHost     string `mapstructure:"SMTP_HOST"`
	SmtpPort     string `mapstructure:"SMTP_PORT"`
	SmtpPassword string `mapstructure:"SMTP_PASSWORD"`
}

func initconfig() (config Config) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	return config

}

func Getconfig() Config {
	config := initconfig()
	return config

}
