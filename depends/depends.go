package depends

import (
	"remote_commands_for_home_IoT/config"
	"remote_commands_for_home_IoT/mqtt"
)

// IAppDeps - dependency injection container
type IAppDeps interface {
	MQTT() mqtt.IMQTT
	Config() *config.Config
}

type di struct {
	mqtt mqtt.IMQTT
	cfg  *config.Config
}

func NewDI(
	mqtt mqtt.IMQTT,
	cfg *config.Config,
) IAppDeps {
	return &di{
		cfg:  cfg,
		mqtt: mqtt,
	}
}

func (di *di) MQTT() mqtt.IMQTT {
	return di.mqtt
}

func (di *di) Config() *config.Config {
	return di.cfg
}
