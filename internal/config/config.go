package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	GRPCServer GRPCServer `yaml:"grpc_server"`
}

type GRPCServer struct {
	Network string `yaml:"network" env-default:"tcp"`
	Address string `yaml:"address" env-default:":50051"`
}

func InitConfig() (*AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	configPath := os.Getenv("CONFIG_PATH")

	var cfg AppConfig
	err = cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
