package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	TokenUsers map[string]string `envconfig:"TOKEN_USERS"` // example: TOKEN_USERS="red:1,green:2,blue:3"
	MQTT
}

type MQTT struct {
	ClientID string `envconfig:"MQTT_CLIENT_ID" default:"test_example"`
	Endpoint string `envconfig:"MQTT_ENDPOINT" default:"tcp://0.0.0.0:1883"`
	User     string `envconfig:"MQTT_USER" default:"test"`
	Password string `envconfig:"MQTT_PASSWORD" default:"test"`
}

func New() *Config {
	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("failed to load envconfig, err: %s", err)
	}
	return &cfg
}
