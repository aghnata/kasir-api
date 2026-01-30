package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	Port   string
	DbConn string
}

// Load loads the configuration from environment variables
func Load() *Config {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	return &Config{
		Port:   viper.GetString("PORT"),
		DbConn: viper.GetString("DB_CONN"),
	}
}
