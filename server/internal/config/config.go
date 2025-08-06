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

	PGUsername string `mapstructure:"DB_USERNAME"`
	PGPassword string `mapstructure:"DB_PASSWORD"`
	PGHost     string `mapstructure:"DB_HOST"`
	PGPort     string `mapstructure:"DB_PORT"`
	PGDatabase string `mapstructure:"DB_DATABASE"`

	CHUsername string `mapstructure:"CLICKHOUSE_USER"`
	CHPassword string `mapstructure:"CLICKHOUSE_PASSWORD"`
	CHHost     string `mapstructure:"CLICKHOUSE_HOST"`
	CHPort     string `mapstructure:"CLICKHOUSE_PORT"`
	CHDatabase string `mapstructure:"CLICKHOUSE_DB"`

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
