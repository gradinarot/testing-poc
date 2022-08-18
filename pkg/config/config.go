package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Filestorage *Filestorage
	Tasks       *Tasks
	Log         *LogConfig
}

type Filestorage struct {
	Endpoint string `envconfig:"CEBE_FILESTORAGE_ENDPOINT" required:"true"`
}

type Tasks struct {
	Endpoint   string `envconfig:"CEBE_TASKS_ENDPOINT" required:"true"`
	LocalToken string `envconfig:"CEBE_TASKS_LOCALTOKEN" default:"testing_localtoken"`
}

type LogConfig struct {
	Level  string `envconfig:"CEBE_LOG_LEVEL" default:"info"`
	Format string `envconfig:"CEBE_LOG_FORMAT" default:"json"`
}

func Load() *Config {
	c := Config{}

	err := envconfig.Process("	CEBE", &c)
	if err != nil {
		log.Fatalln(err)
	}
	return &c
}
