package config

import "sync"

var cnfOnce = sync.Once{}

type Mode string

const DebugMode = Mode("debug")
const ReleaseMode = Mode("release")

type Apm struct {
	ServiceName string `mapstructure:"APM_SERVICE_NAME"`
	ServerURL   string `mapstructure:"APM_SERVER_URL"`
	SecretToken string `mapstructure:"APM_SECRET_TOKEN"`
	Environment string `mapstructure:"APM_ENVIRONMENT"`
}

type Config struct {
	Version     string `mapstructure:"VERSION"                  validate:"required"`
	Mode        Mode   `mapstructure:"MODE"                     validate:"required"`
	ServiceName string `mapstructure:"SERVICE_NAME"             validate:"required"`
	HttpPort    int    `mapstructure:"HTTP_PORT"                validate:"required"`

	//Apm                *Apm
	//HealthCheckRoute string `mapstructure:"HEALTH_CHECK_ROUTE"       validate:"required"`
}

var config *Config

func GetConfig() *Config {
	cnfOnce.Do(func() {
		loadConfig()
	})

	return config
}
