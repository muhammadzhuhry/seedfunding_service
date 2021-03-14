package infra

import "github.com/spf13/viper"

type Config struct {
	Server    string `mapstructure:"SERVER_ADDRESS"`
	ServerKey string `mapstructure:"SERVER_KEY"`
	ClientKey string `mapstructure:"CLIENT_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	//viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
