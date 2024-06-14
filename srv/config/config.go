package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	GRPC *grpcConfig `yaml:"grpc"`
	DNS  *DNSConfig  `yaml:"dns"`
}

type DNSConfig struct {
	PathResolve string `yaml:"pathResolve"`
}

type grpcConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func NewConfig() *Config {
	return parseConfig()
}

func parseConfig() *Config {
	file, err := os.ReadFile("srv/config/config.yaml")
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}
	var cfg Config

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("failed to parse config file: %v", err)
	}

	if cfg.GRPC.Port == 0 {
		port, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
		if err != nil {
			log.Fatalf("failed to parse config file: %v", err)
		}
		if port == 0 {
			port = 44044
		}

		cfg.GRPC.Port = port
	}

	return &cfg
}