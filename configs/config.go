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

//type EnvConfig struct {
//	PostgresAppDbPassword string `mapstructure:"POSTGRES_APP_DB_PASSWORD"`
//
//	FrontendOriginURI string `mapstructure:"FRONTEND_ORIGIN"`
//	JWTSecret         string `mapstructure:"JWT_SECRET"`
//	TokenExpriredIn   string `mapstructure:"TOKEN_EXPIRED_IN"`
//	TokenMaxAge       string `mapstructure:"TOKEN_MAX_AGE"`
//
//	GoogleClientID     string `mapstructure:"GOOGLE_OAUTH_CLIENT_ID"`
//	GoogleClientSecret string `mapstructure:"GOOGLE_OAUTH_CLIENT_SECRET"`
//	GoogleRedirectURI  string `mapstructure:"GOOGLE_OAUTH_REDIRECT_URI"`
//}

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
