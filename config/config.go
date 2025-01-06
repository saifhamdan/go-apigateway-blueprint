// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppEnv string `mapstructure:"APP_ENV"`
	// HTTP config
	HttpHost         string `mapstructure:"HTTP_HOST"`
	HttpPort         string `mapstructure:"HTTP_PORT"`
	HttpCookieSecret string `mapstructure:"HTTP_COOKIE_SECRET"`
	// Postgres config
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresName     string `mapstructure:"POSTGRES_NAME"`
	// Redis config
	RedisHost string `mapstructure:"REDIS_HOST"`
	RedisPort string `mapstructure:"REDIS_PORT"`
	// NATS config
	NatsHost string `mapstructure:"NATS_HOST"`
	NatsPort string `mapstructure:"NATS_PORT"`
	// Logging config
	LogFile  string `mapstructure:"LOG_FILE"`
	LogLevel string `mapstructure:"LOG_LEVEL"`
}

func NewConfig(path string) (*Config, error) {
	config := &Config{}

	viper.AddConfigPath(path)
	viper.SetConfigName(".env") // Set the correct config file name without the extension
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(config)

	return config, err
}
