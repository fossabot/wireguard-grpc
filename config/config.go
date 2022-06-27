package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	LogLevel          string `yaml:"loglevel" env:"LOG_LEVEL" env-description:"application log level"`
	Port              int    `yaml:"port" env:"PORT" env-default:"8081" env-description:"GRPC server port"`
	WgExecutable      string `yaml:"wg" env:"WG_EXE" env-default:"wg" env-description:"wireguard executable"`
	WgQuickExecutable string `yaml:"wgquick" env:"WGQUICK_EXE" env-default:"wg-quick" env-description:"wg-quick wrapper"`
}

// ReadConfig reads configuration
func ReadConfig() *Config {
	cfg := Config{}
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatalf("Unable to read configuration %v", err)
	}
	return &cfg
}
