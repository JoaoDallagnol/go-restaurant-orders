package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		Schema   string
		SSLMode  string
	}
	Auth struct {
		Secret            string
		ExpirationMinutes int `mapstructure:"expiration_minutes"`
	}
}

var AppConfig *Config

func LoadConfig() {
	v := viper.New()

	// Use config.yaml from internal/config relative to project root
	v.AddConfigPath("internal/config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.SetEnvPrefix("app")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("❌ Failed to read config file: %v", err)
	}

	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		log.Fatalf("❌ Failed to parse config into struct: %v", err)
	}

	AppConfig = cfg
	log.Println("✅ Configuration loaded successfully")
}
