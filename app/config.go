package app

import (
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

var cfgOnce sync.Once
var cfgInstance *config

type config struct {
	Env                    string        `mapstructure:"ENV"`
	ServerPort             int           `mapstructure:"SERVER_PORT"`
	CorsHost               string        `mapstructure:"CORS_HOST"`
	SqlShow                bool          `mapstructure:"SQL_SHOW"`
	DbDriver               string        `mapstructure:"DATABASE_DRIVER"`
	DbConn                 string        `mapstructure:"DATABASE_CONN"`
	AuthTokenSymmetricKey  string        `mapstructure:"AUTH_TOKEN_SYMMETRIC_KEY"`
	AuthTokenDuration      time.Duration `mapstructure:"AUTH_TOKEN_DURATION"`
	AuthGoogleClientID     string        `mapstructure:"AUTH_GOOGLE_CLIENT_ID"`
	AuthGoogleClientSecret string        `mapstructure:"AUTH_GOOGLE_CLIENT_SECRET"`
}

func Config() *config {
	cfgOnce.Do(func() {
		if cfgInstance == nil {
			viper.SetConfigFile(".env")

			viper.AutomaticEnv()

			err := viper.ReadInConfig()
			if err != nil {
				log.Fatal("failed to load config", err)
			}

			err = viper.Unmarshal(&cfgInstance)
			if err != nil {
				log.Fatal("failed to load config", err)
			}
		}
	})

	return cfgInstance
}
