package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`
	HTTP
	Storage
	Minio
}

type HTTP struct {
	Address  string        `yaml:"address" env-required:"true"`
	TimeOut  time.Duration `yaml:"timeout" env-required:"true"`
	TimeIdle time.Duration `yaml:"time_idle" env-required:"true"`
}

type Storage struct {
	Host     string `yaml:"DB_HOST" env-required:"true"`
	Port     string `yaml:"DB_PORT" env-required:"true"`
	User     string `yaml:"DB_USER" env-required:"true"`
	Database string `yaml:"DB_NAME" env-required:"true"`
	Password string `yaml:"DB_PASSWORD" env-required:"true"`
}

type Minio struct {
	Host     string `yaml:"HOST" env-required:"true"`
	Port     string `yaml:"PORT" env-required:"true"`
	User     string `yaml:"USER" env-required:"true"`
	Password string `yaml:"PASSWORD" env-required:"true"`
	Bucket   string `yaml:"BUCKET" env-required:"true"`
}

func MustLoad() Config {
	var cfg Config

	path := os.Getenv("CONFIG_PATH")

	if path == "" {
		panic("config file not found")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found")
	}

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
