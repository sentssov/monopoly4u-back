package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}
	Database struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		SSLMode  string `mapstructure:"sslmode"`
	}
}

func InitConfig() (config *Config, err error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return config, err
}
