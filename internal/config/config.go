package config

import (
	"context"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
	"github.com/pkg/errors"
)

const DefaultConfigPath = "configs/config.yaml"

type Config struct {
	Server struct {
		HTTP struct {
			Host string `config:"host"`
			Port string `config:"host"`
		} `config:"http"`
		Grpc struct {
			Host string `config:"host"`
			Port string `config:"host"`
		} `config:"grpc"`
	} `config:"server"`
	Logger struct {
		Path  string `config:"path"`
		Level string `config:"level"`
	} `config:"logger"`
	DB struct {
		Postgres struct {
			Dialect string `config:"dialect"`
			Dsn     string `config:"dsn"`
		} `config:"postgres"`
		Memory struct {
			MaxSize uint `config:"maxsize"`
		} `config:"memory"`
	} `config:"db"`
	Repository struct {
		Type string `config:"type"`
	} `config:"repository"`
	Redis struct {
		Address  string `config:"address"`
		Password string `config:"password"`
	} `config:"redis"`
	Kafka struct {
		Address       string `config:"address"`
		BatchTimeout  string `config:"batchtimeout"`
		DialerTimeout string `config:"dialertimeout"`
		Group         string `config:"group"`
	} `config:"kafka"`
}

func New(filePath string) (*Config, error) {
	loader := confita.NewLoader(
		file.NewBackend(filePath),
	)

	cfg := &Config{}
	err := loader.Load(context.Background(), cfg)
	if err != nil {
		return nil, errors.Wrap(err, "cannot load config")
	}

	return cfg, nil
}
