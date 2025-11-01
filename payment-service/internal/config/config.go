package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		Schema   string `mapstructure:"schema"`
		SSLMode  string `mapstructure:"sslmode"`
	} `mapstructure:"database"`

	RabbitMQ struct {
		Url          string `mapstructure:"url"`
		ExchangeName string `mapstructure:"exchange_name"`
		RoutingKey   string `mapstructure:"routing_key"`
	} `mapstructure:"rabbitmq"`

	OrderService struct {
		ConnectTimeout int    `mapstructure:"connect_timeout"`
		BaseURL        string `mapstructure:"base_url"`
		Endpoint       struct {
			GetOrderByID string `mapstructure:"get_order_by_id"`
		} `mapstructure:"endpoint"`
	} `mapstructure:"order_service"`
}

var AppConfig *Config

func LoadConfig() {
	v := viper.New()

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
