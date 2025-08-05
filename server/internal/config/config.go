package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	Port        int    `mapstructure:"APP_PORT"`
	Environment string `mapstructure:"APP_ENV"`
	AppName     string `mapstructure:"APP_NAME"`
	DebugMode   bool   `mapstructure:"DEBUG"`

	PGUsername string `mapstructure:"PG_USERNAME"`
	PGPassword string `mapstructure:"PG_PASSWORD"`
	PGHost     string `mapstructure:"PG_HOST"`
	PGPort     string `mapstructure:"PG_PORT"`
	PGDatabase string `mapstructure:"PG_DATABASE"`

	CHUsername string `mapstructure:"CH_USERNAME"`
	CHPassword string `mapstructure:"CH_PASSWORD"`
	CHHost     string `mapstructure:"CH_HOST"`
	CHPort     string `mapstructure:"CH_PORT"`
	CHDatabase string `mapstructure:"CH_DATABASE"`

	RedisHost string `mapstructure:"REDIS_HOST"`
	RedisPort string `mapstructure:"REDIS_PORT"`

	StorageURL        string `mapstructure:"STORAGE_URL"`
	StorageAccessKey  string `mapstructure:"STORAGE_ACCESS_KEY"`
	StorageSecretKey  string `mapstructure:"STORAGE_SECRET_KEY"`
	StorageBucketName string `mapstructure:"STORAGE_BUCKET_NAME"`

	ResendKey string `mapstructure:"RESEND_API_KEY"`

	JWTAccessKey  string `mapstructure:"JWT_ACCESS_KEY"`
	JWTRefreshKey string `mapstructure:"JWT_REFRESH_KEY"`
}

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("error reading config file, %s", err)
		}

		if err := viper.Unmarshal(&config); err != nil {
			log.Fatalf("error unmarshalling config, %s", err)
		}
	})

	return config
}
